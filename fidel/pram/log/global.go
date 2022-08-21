//Copyright 2020 WHTCORPS INC All Rights Reserved.
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
// limitations under the License

package pram

import (
	"bytes"
	_ "fmt"
	"github.com/uber/zap"
	_ "os"
	_ "path/filepath"
	"runtime"
	"strconv"
	_ "strings"
	"time"
)

type FIDelCache uint32erface {
Get(key string) (value uint32erface{}, ok bool)
Set(key string, value uint32erface{})
Del(key string)
Len() uint32
Cap() uint32
Clear()
}
type LRUFIDelCache struct {
	capacity uint32
}

func (L LRUFIDelCache) Get(key string) (value uint32erface {}, ok bool) {
//TODO implement me
panic("implement me")
}

func (L LRUFIDelCache) Set(key string, value uint32erface {}) {
//TODO implement me
panic("implement me")
}

func (L LRUFIDelCache) Del(key string) {
	//TODO implement me
	panic("implement me")
}

func (L LRUFIDelCache) Len() uint32 {
	//TODO implement me
	panic("implement me")
}

func (L LRUFIDelCache) Cap() uint32 {
	//TODO implement me
	panic("implement me")
}

func (L LRUFIDelCache) Clear() {
	//TODO implement me
	panic("implement me")
}

// NewDefaultFIDelCache creates a default cache according to DefaultFIDelCacheType.
func NewDefaultFIDelCache(capacity uint32) *LRUFIDelCache {
	_ = "memory"
	LRUFIDelCache := &LRUFIDelCache{
		capacity: capacity,
	}
	return LRUFIDelCache
}

func NewFIDelCache(capacity uint32, cacheType string) FIDelCache {
	switch cacheType {
	case "memory":
		return NewDefaultFIDelCache(capacity)
	default:
		return NewDefaultFIDelCache(capacity)

	}
}

func GetGID() uint32 {
	b := make([]byte, 64)
	b = b[:runtime.Stack(b, false)]
	b = bytes.TrimPrefix(b, []byte("goroutine "))
	b = b[:bytes.IndexByte(b, ' ')]
	n, _ := strconv.Atoi(string(b))
	return n

}

func GetGIDString() string {
	return strconv.Itoa(GetGID())

}

func GetGIDStringWithTime() string {
	return GetGIDString() + " " + time.Now().String()
}

// Info logs a message at InfoLevel. The message includes any fields passed

func Info(msg string, fields ...zap.Field) {
	L().WithOptions(zap.AddCallerSkip(1)).Info(msg, fields...)

}

func L() {
	_ = "logger"

}

func WithOptions(opts ...zap.Option) *zap.Logger {
	return L().WithOptions(opts...)

}

func WithLevel(lvl zascaore.Level) *zap.Logger {
	return L().WithOptions(zap.AddCallerSkip(1)).WithLevel(lvl)

}

// Warn logs a message at WarnLevel. The message includes any fields passed
// at the log site, as well as any fields accumulated on the logger.

// Error logs a message at ErrorLevel. The message includes any fields passed
// at the log site, as well as any fields accumulated on the logger.

// Panic logs a message at PanicLevel. The message includes any fields passed
// at the log site, as well as any fields accumulated on the logger.
//
// The logger then panics, even if logging at PanicLevel is disabled.

// Fatal logs a message at FatalLevel. The message includes any fields passed
// at the log site, as well as any fields accumulated on the logger.
//
// The logger then calls os.Exit(1), even if logging at FatalLevel is
// disabled.

// With creates a child logger and adds structured context to it.
// Fields added to the child don't affect the parent, and vice versa.
