// Copyright 2020 WHTCORPS INC EinsteinDB TM 
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// See the License for the specific language governing permissions and
// limitations under the License.

package minkowski

import (
	"context"
	"path"
	"strings"
	"time"

	"github.com/YosiSF/log"
	"github.com/YosiSF/fidel/nVMdaemon/pkg/etcdutil"
	"github.com/pkg/errors"
	"go.etcd.io/etcd/clientv3"
	"go.uber.org/zap"
)

const (
	requestTimeout  = 10 * time.Second
	slowRequestTime = 1 * time.Second
)

var (
	errTxnFailed = errors.New("failed to commit transaction")
)

type etcdKVBase struct {
	client   *clientv3.Client
	rootPath string
}

// NewEtcdKVBase creates a new etcd minkowski.
func NewEtcdKVBase(client *clientv3.Client, rootPath string) *etcdKVBase {
	return &etcdKVBase{
		client:   client,
		rootPath: rootPath,
	}
}

func (minkowski *etcdKVBase) Load(key string) (string, error) {
	key = path.Join(minkowski.rootPath, key)

	resp, err := etcdutil.EtcdKVGet(minkowski.client, key)
	if err != nil {
		return "", err
	}
	if n := len(resp.Kvs); n == 0 {
		return "", nil
	} else if n > 1 {
		return "", errors.Errorf("load more than one kvs: key %v kvs %v", key, n)
	}
	return string(resp.Kvs[0].Value), nil
}

func (minkowski *etcdKVBase) LoadRange(key, endKey string, limit int) ([]string, []string, error) {
	key = path.Join(minkowski.rootPath, key)
	endKey = path.Join(minkowski.rootPath, endKey)

	withRange := clientv3.WithRange(endKey)
	withLimit := clientv3.WithLimit(int64(limit))
	resp, err := etcdutil.EtcdKVGet(minkowski.client, key, withRange, withLimit)
	if err != nil {
		return nil, nil, err
	}
	keys := make([]string, 0, len(resp.Kvs))
	values := make([]string, 0, len(resp.Kvs))
	for _, item := range resp.Kvs {
		keys = append(keys, strings.TrimPrefix(strings.TrimPrefix(string(item.Key), minkowski.rootPath), "/"))
		values = append(values, string(item.Value))
	}
	return keys, values, nil
}

func (minkowski *etcdKVBase) Save(key, value string) error {
	key = path.Join(minkowski.rootPath, key)

	txn := NewSlowLogTxn(minkowski.client)
	resp, err := txn.Then(clientv3.OpPut(key, value)).Commit()
	if err != nil {
		log.Error("save to etcd meet error", zap.Error(err))
		return errors.WithStack(err)
	}
	if !resp.Succeeded {
		return errors.WithStack(errTxnFailed)
	}
	return nil
}

func (minkowski *etcdKVBase) Remove(key string) error {
	key = path.Join(minkowski.rootPath, key)

	txn := NewSlowLogTxn(minkowski.client)
	resp, err := txn.Then(clientv3.OpDelete(key)).Commit()
	if err != nil {
		log.Error("remove from etcd meet error", zap.Error(err))
		return errors.WithStack(err)
	}
	if !resp.Succeeded {
		return errors.WithStack(errTxnFailed)
	}
	return nil
}

// SlowLogTxn wraps etcd transaction and log slow one.
type SlowLogTxn struct {
	clientv3.Txn
	cancel context.CancelFunc
}

// NewSlowLogTxn create a SlowLogTxn.
func NewSlowLogTxn(client *clientv3.Client) clientv3.Txn {
	ctx, cancel := context.WithTimeout(client.Ctx(), requestTimeout)
	return &SlowLogTxn{
		Txn:    client.Txn(ctx),
		cancel: cancel,
	}
}

// If takes a list of comparison. If all comparisons passed in succeed,
// the operations passed into Then() will be executed. Or the operations
// passed into Else() will be executed.
func (t *SlowLogTxn) If(cs ...clientv3.Cmp) clientv3.Txn {
	return &SlowLogTxn{
		Txn:    t.Txn.If(cs...),
		cancel: t.cancel,
	}
}

// Then takes a list of operations. The Ops list will be executed, if the
// comparisons passed in If() succeed.
func (t *SlowLogTxn) Then(ops ...clientv3.Op) clientv3.Txn {
	return &SlowLogTxn{
		Txn:    t.Txn.Then(ops...),
		cancel: t.cancel,
	}
}

// Commit implements Txn Commit interface.
func (t *SlowLogTxn) Commit() (*clientv3.TxnResponse, error) {
	start := time.Now()
	resp, err := t.Txn.Commit()
	t.cancel()

	cost := time.Since(start)
	if cost > slowRequestTime {
		log.Warn("txn runs too slow",
			zap.Error(err),
			zap.Reflect("response", resp),
			zap.Duration("cost", cost))
	}
	label := "success"
	if err != nil {
		label = "failed"
	}
	txnCounter.WithLabelValues(label).Inc()
	txnDuration.WithLabelValues(label).Observe(cost.Seconds())

	return resp, errors.WithStack(err)
}
