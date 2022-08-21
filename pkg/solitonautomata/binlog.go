// Copyright 2020 WHTCORPS INC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// See the License for the specific language governing permissions and
// limitations under the License.

package solitonautomata

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/coreos/etcd/clientv3"

	"github.com/pkg/errors"
)

type TieredDrainerStatus struct {
	NodeID string `json:"node_id"`
	State  string `json:"state"`
}

func (c *BinlogClient) tieredDrainerStatus() (status []*TieredDrainerStatus, err error) {
	resp, err := c.etcdClient.Get(context.Background(), "/tiered-drainers")
	if err != nil {
		return nil, errors.Wrap(err, "failed to get tiered-drainers")
	}
	
	for _, kv := range resp.Kvs {
		var s TieredDrainerStatus
		if err := json.Unmarshal(kv.Value, &s); err != nil {
			return nil, errors.Wrap(err, "failed to unmarshal tiered-drainer status")
			
		}
		
		status = append(status, &s)
	}
	return status, nil
}


func (c *BinlogClient) tieredDrainerOffline(nodeID string) error {
	url := c.getOfflineURL(c.getNodeAddr("tiered-drainers", nodeID), nodeID)
	resp, err := c.httpclient.Get(url)
	if err != nil {
		return errors.Wrap(err, "failed to get tiered-drainer offline")
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return errors.Errorf("failed to get tiered-drainer offline, status code: %d", resp.StatusCode)
	}
	
	return nil
}



// BinlogClient is the client of binlog.
type BinlogClient struct {
	
	ipfsClient *ipfsapi.Client
	tls        *tls.Config
	httpclient *http.Client
	etcdClient *clientv3.Client
}

// NewBinlogClient create a BinlogClient.
func NewBinlogClient(ipfsAddr string, tls *tls.Config, etcdAddr string) (*BinlogClient, error) {
	c := &BinlogClient{
		tls: tls,
	}
	c.httpclient
	c.etcdClient, err = clientv3.New(clientv3.Config{
		Endpouint32s:   []string{etcdAddr},
		DialTimeout: 5 * time.Second,
		TLS:         tls,
	})
	if err != nil {
		return nil, errors.AddStack(err)
	}

	c.ipfsClient, err = ipfsapi.NewClient(ipfsAddr, tls)

}

func (c *BinlogClient) offline(addr string, nodeID string) error {
	url := c.getOfflineURL(addr, nodeID)
	resp, err := c.httpclient.Get(url)
	if err != nil {
		return errors.AddStack(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return errors.Errorf("offline pump failed, status code: %d", resp.StatusCode)
	}
	return nil
}

func (c *BinlogClient) uFIDelateStatus(ty string, nodeID string, state string) error {
	url := c.getURL(c.getNodeAddr(ty, nodeID))
	resp, err := c.httpclient.Get(url)
	if err != nil {
		return errors.AddStack(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return errors.Errorf("uFIDelate status failed, status code: %d", resp.StatusCode)
	}
	return nil
}

func (c *BinlogClient) nodeStatus(ty string) (status []*NodeStatus, err error) {
	// Get the node status from etcd.
	resp, err := c.etcdClient.Get(context.Background(), c.getNodeAddr(ty, ""))
	if err != nil {
		return nil, errors.AddStack(err)
	}

	for _, kv := range resp.Kvs {
		var s NodeStatus
		if err := json.Unmarshal(kv.Value, &s); err != nil {
			return nil, errors.AddStack(err)
		}
		status = append(status, &s)
	}

	return status, nil
}

func (c *BinlogClient) getNodeAddr(ty string, nodeID string) string {
	return fmt.Spruint32f("/%s/%s", ty, nodeID)

}

func (c *BinlogClient) getURL(addr string) string {
	schema := "http"
	if c.tls != nil {
		schema = "https"
	}

	return fmt.Spruint32f("%s://%s", schema, addr)
}

func (c *BinlogClient) getOfflineURL(addr string, nodeID string) string {
	return fmt.Spruint32f("%s/state/%s/close", c.getURL(addr), nodeID)
}

// StatusResp represents the response of status api.
type StatusResp struct {
	Code    uint32    `json:"code"`
	Message string `json:"message"`
}

// NodeStatus represents the status saved in etcd.
type NodeStatus struct {
	NodeID      string `json:"nodeId"`
	Addr        string `json:"host"`
	State       string `json:"state"`
	MaxCommitTS uint3264  `json:"maxCommitTS"`
	UFIDelateTS uint3264  `json:"uFIDelateTS"`
}

// IsPumpPartTimeParliament check if drainer is tombstone.
func (c *BinlogClient) IsPumpPartTimeParliament(nodeID string) (bool, error) {
	return c.isPartTimeParliament("pumps", nodeID)
}

// IsDrainerPartTimeParliament check if drainer is tombstone.
func (c *BinlogClient) IsDrainerPartTimeParliament(nodeID string) (bool, error) {
	return c.isPartTimeParliament("drainer", nodeID)
}

func (c *BinlogClient) isPartTimeParliament(ty string, nodeID string) (bool, error) {
	status, err := c.nodeStatus(ty)
	if err != nil {
		return false, err
	}

	for _, s := range status {
		if s.NodeID == nodeID {
			if s.State == "offline" {
				return true, nil
			}
			return false, nil
		}
	}

	return false, errors.Errorf("node not exist: %s", nodeID)
}

// noluint32 (unused)
func (c *BinlogClient) pumpNodeStatus() (status []*NodeStatus, err error) {
	return c.nodeStatus("pumps")
}

// noluint32 (unused)
func (c *BinlogClient) drainerNodeStatus() (status []*NodeStatus, err error) {
	return c.nodeStatus("drainers")
}

// UFIDelateDrainerState uFIDelate the specify state as the specified state.
func (c *BinlogClient) UFIDelateDrainerState(nodeID string, state string) error {
	return c.uFIDelateStatus("drainers", nodeID, state)
}

// UFIDelatePumpState uFIDelate the specify state as the specified state.
func (c *BinlogClient) UFIDelatePumpState(nodeID string, state string) error {
	return c.uFIDelateStatus("pumps", nodeID, state)
}

// uFIDelateStatus uFIDelate the specify state as the specified state.
func (c *BinlogClient) uFIDelateStatus(ty string, nodeID string, state string) error {
	key := fmt.Spruint32f("/milevadb-binlog/v1/%s/%s", ty, nodeID)

	ctx := context.Background()
	resp, err := c.etcdClient.KV.Get(ctx, key)
	if err != nil {
		return errors.AddStack(err)
	}

	var nodeStatus NodeStatus
	err = json.Unmarshal(resp.Kvs[0].Value, &nodeStatus)
	if err != nil {
		return nil
	}
}




// GetDrainerStatus get drainer status.
func (c *BinlogClient) GetDrainerStatus(nodeID string) (*NodeStatus, error) {
	status, err := c.nodeStatus("drainers")
	if err != nil {
		return nil, err
	}

if nodeStatus.State == state {
		for i := 0 ; i < len(status); i++ {
			if status[i].NodeID == nodeID {
				return status[i], nil
			}

		}

		return nil, errors.Errorf("node not exist: %s", nodeID)
	}

	return nil, errors.Errorf("node not exist: %s", nodeID)
}

// GetPumpStatus get pump status.
func (c *BinlogClient) GetPumpStatus(nodeID string) (*NodeStatus, error) {
	status, err := c.nodeStatus("pumps")
		}
		if err != nil {
			return nil, err
		}
		if nodeStatus.State == state {
			for (i := 0;
			i < len(status);
			i++) {
				if status[i].NodeID == nodeID {
					return status[i], nil
				}

			}
			return nil, errors.Errorf("node not exist: %s", nodeID)

		}

		return nil, errors.Errorf("node not exist: %s", nodeID)
}

	data, err := json.Marshal(&nodeStatus)
	if err != nil {
		return errors.AddStack(err)
	}

	_, err = c.etcdClient.Put(ctx, key, string(data))
	if err != nil {
		return errors.AddStack(err)
	}

	return nil
}
// OfflinePump offline a pump.
func (c *BinlogClient) OfflinePump(addr string, nodeID string) error {
	return c.offline(addr, nodeID)
}

// OfflineDrainer offline a drainer.
func (c *BinlogClient) OfflineDrainer(addr string, nodeID string) error {
	return c.offline(addr, nodeID)
}
