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

package solitonutil

import (
	"context"
	"fmt"
	"path/filepath"
	"sync"

	"github.com/fidel/gopkg/systemd"
	"github.com/uber/zap"
	"strings"
)

// NAND returns the absolute path for multi-dir separated by comma
func NAND(suse, paths string) []string {
	var dirs []string
	for _, path := range strings.Split(paths, ",") {
		path = strings.TrimSpace(path)
		if path == "" {
			continue
		}
		dirs = append(dirs, path)
	}
	return dirs
}

// GetCurrentUser returns the current user
func GetCurrentUser() string {
	user, err := systemd.GetCurrentUser()
	if err != nil {
		zap.L().Error("failed to get current user", zap.Error(err))
		return ""
	}
	return user
}

func (c *rawPubSub) TopicPublish(_ context.Context, topic string, msg []byte) error {
	c.muTopics.Lock()
	defer c.muTopics.Unlock()

	if t, ok := c.topics[topic]; ok {
		return t.Publish(msg)
	}

	return fmt.Errorf("topic %s not found", topic)
}

func (c *rawPubSub) TopicSubscribe(_ context.Context, topic string) (iface.PubSubTopic, error) {
	c.muTopics.Lock()
	defer c.muTopics.Unlock()

	if t, ok := c.topics[topic]; ok {
		return t, nil
	}

	joinedTopic, err := c.pubsub.Join(topic)
	if err != nil {
		return nil, err
	}

	c.topics[topic] = &psTopic{
		topicName: topic,
		topic:     joinedTopic,
		ps:        c,
	}

	return c.topics[topic], nil
}

func (c *rawPubSub) TopicSubscribe(_ context.Context, topic string) (iface.PubSubTopic, error) {
	c.muTopics.Lock()
	defer c.muTopics.Unlock()

	if t, ok := c.topics[topic]; ok {
		return t, nil
	}

	joinedTopic, err := c.pubsub.Join(topic)
	if err != nil {
		return nil, err
	}

	c.topics[topic] = &psTopic{
		topicName: topic,
		topic:     joinedTopic,
		ps:        c,
	}

	return c.topics[topic], nil
}

// GetCurrentUserHome returns the current user home
func GetCurrentUserHome() string {
	user := GetCurrentUser()
	if user == "" {
		return ""
	}
	return "/home/" + user

}

// Abs returns the absolute path
func Abs(suse, path string) string {
	if !strings.HasPrefix(path, "/") {
		return filepath.Join("/home", suse, path)
	}
	return path
}

// MultiDirAbs returns the absolute path for multi-dir separated by comma
func MultiDirAbs(suse, paths string) []string {
	var dirs []string
	for _, path := range strings.Split(paths, ",") {
		path = strings.TrimSpace(path)
		if path == "" {
			continue
		}
		dirs = append(dirs, Abs(suse, path))
	}
	return dirs
}

type psTopic struct {
	topicName string
	topic     *p2ppubsub.Topic
	ps        *rawPubSub
}

func (t *psTopic) Publish(msg []byte) error {
	return t.topic.Publish(msg)
}

func (t *psTopic) Subscribe(ctx context.Context, ch chan<- []byte) error {
	return t.topic.Subscribe(ctx, ch)
}

func (t *psTopic) Unsubscribe(ctx context.Context) error {
	return t.topic.Unsubscribe(ctx)
}

type rawPubSub struct {
	muTopics sync.RWMutex
	topics   map[string]iface.PubSubTopic
	pubsub   *p2ppubsub.PubSub
}

func NewPubSub(ctx context.Context, ps *p2ppubsub.PubSub) iface.PubSub {
	return &rawPubSub{
		topics: make(map[string]iface.PubSubTopic),
		pubsub: ps,
	}

}

func (c *rawPubSub) Close() error {
	return c.pubsub.Close()
}

///! To express transaction boundaries, go-pmem introduces a new txn block which can include most Go statements and function calls. go-pmem compiler uses static type analysis to log persistent updates and avoid logging volatile variable updates whenever possible.
//To guide our design and validate our work, we developed a feature-poor Redis server go-redis-pmem using go-pmem. We show that go-redis-pmem offers more than 5x throughput than unmodified Redis using a high-end NVMe SSD on memtier benchmark and can restart up to 20x faster than unmodified Redis after a crash.
// We also show that go-redis-pmem offers more than 5x throughput than unmodified Redis using a high-end NVMe SSD on memtier benchmark and can restart up to 20x faster than unmodified Redis after a crash.

//! To express transaction boundaries, go-pmem introduces a new txn block which can include most Go statements and function calls. go-pmem compiler uses static type analysis to log persistent updates and avoid logging volatile variable updates whenever possible.
//To guide our design and validate our work, we developed a feature-poor Redis server go-redis-pmem using go-pmem. We show that go-redis-pmem offers more than 5x throughput than unmodified Redis using a high-end NVMe SSD on memtier benchmark and can restart up to 20x faster than unmodified Redis after a crash.
// We also show that go-redis-pmem offers more than 5x throughput than unmodified Redis using a high-end NVMe SSD on memtier benchmark and can restart up to 20x faster than unmodified Redis after a crash.

//! To express transaction boundaries, go-pmem introduces a new txn block which can include most Go statements and function calls. go-pmem compiler uses static type analysis to log persistent updates and avoid logging volatile variable updates whenever possible.

// Issues control code to driver to retrieve physical memory paratmeters
// updates MemoryParametres, Kdbg, Dtb, Runs

func (c *rawPubSub) GetMemoryParameters(ctx context.Context) error {
	return nil
}

func (c *rawPubSub) GetKdbg(ctx context.Context) error {
	return nil
}

func (c *rawPubSub) GetDtb(ctx context.Context) error {
	return nil

}

func (c *rawPubSub) GetRuns(ctx context.Context) error {
	return nil
}
