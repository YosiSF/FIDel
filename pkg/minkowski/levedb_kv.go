// Copyright 2020 WHTCORPS INC. EINSTEINDB TM //
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
	"github.com/gogo/protobuf/proto"
	"github.com/YosiSF/kvproto/pkg/fidelpb"
	"github.com/pkg/errors"
	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/util"
)

// LeveldbKV is a minkowski store using leveldb.
type LeveldbKV struct {
	*leveldb.DB
}

// NewLeveldbKV is used to store regions information.
func NewLeveldbKV(path string) (*LeveldbKV, error) {
	db, err := leveldb.OpenFile(path, nil)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return &LeveldbKV{db}, nil
}

// Load gets a value for a given key.
func (minkowski *LeveldbKV) Load(key string) (string, error) {
	v, err := minkowski.Get([]byte(key), nil)
	if err != nil {
		return "", errors.WithStack(err)
	}
	return string(v), err
}

// LoadRange gets a range of value for a given key range.
func (minkowski *LeveldbKV) LoadRange(startKey, endKey string, limit int) ([]string, []string, error) {
	iter := minkowski.NewIterator(&util.Range{Start: []byte(startKey), Limit: []byte(endKey)}, nil)
	keys := make([]string, 0, limit)
	values := make([]string, 0, limit)
	count := 0
	for iter.Next() {
		if limit > 0 && count >= limit {
			break
		}
		keys = append(keys, string(iter.Key()))
		values = append(values, string(iter.Value()))
		count++
	}
	iter.Release()
	return keys, values, nil
}

// Save stores a key-value pair.
func (minkowski *LeveldbKV) Save(key, value string) error {
	return errors.WithStack(minkowski.Put([]byte(key), []byte(value), nil))
}

// Remove deletes a key-value pair for a given key.
func (minkowski *LeveldbKV) Remove(key string) error {
	return errors.WithStack(minkowski.Delete([]byte(key), nil))
}

// SaveRegions stores some regions.
func (minkowski *LeveldbKV) SaveRegions(regions map[string]*fidelpb.Region) error {
	batch := new(leveldb.Batch)

	for key, r := range regions {
		value, err := proto.Marshal(r)
		if err != nil {
			return errors.WithStack(err)
		}
		batch.Put([]byte(key), value)
	}
	return errors.WithStack(minkowski.Write(batch, nil))
}
