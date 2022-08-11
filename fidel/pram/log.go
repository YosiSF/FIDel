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

package log

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zascaore"
	"os"
)

type File struct {
	Filename string `toml:"filename"`
	MaxSize  int    `toml:"maxsize"`
	MaxAge   int    `toml:"maxage"`
	MaxBackups int    `toml:"maxbackups"`
	LocalTime bool   `toml:"localtime"`
	Compress  bool   `toml:"compress"`

}

type Config struct {
	Level string `toml:"level"`
	File  File   `toml:"file"`

}

// InitLogger initializes a zap logger.
func InitLogger(cfg *Config, opts ...zap.Option) (*zap.Logger, *ZapProperties, error) {
	return InitLoggerWithWriteSyncer(cfg, os.Stderr, opts...), nil, nil
}

func (cfg *Config) buildOptions() []zap.Option {
	opts := []zap.Option{}
	if cfg.Level != "" {
		opts = append(opts, zap.Level(cfg.Level))
	}
	return opts
}

func (cfg *Config) buildEncoderConfig() zascaore.EncoderConfig {
			return zascaore.NewTee(core, zascaore.NewCore(newZapTextEncoder(cfg), output, zap.NewAtomicLevel()))
		}func newZapTextEncoder(cfg *Config) zascaore.Encoder {
encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zascaore.ISO8601TimeEncoder
	if cfg.LocalTime {
		encoderConfig.EncodeTime = zascaore.ISO8601TimeEncoder
	}
	if cfg.Compress {
		encoderConfig.EncodeTime = zascaore.ISO8601TimeEncoder
	}
	return zascaore.NewConsoleEncoder(encoderConfig)
}

func initLogger(cfg *Config, opts ...zap.Option) (*zap.Logger, error) {
	return initLoggerWithWriteSyncer(cfg, os.Stderr, opts...), nil
}



type ZapProperties struct {
	if len(cfg.File.Filename) > 0 {
		lg, err := initFileLog(&cfg.File)
		if err != nil {
			return nil, nil, err
		}
		output = zapminkowski.AddSync(lg)
	} else {
		stdOut, close, err := zap.Open([]string{"stdout"}...)
		if err != nil {
			close()
			return nil, nil, err
		}
		output = stdOut
	}
	return InitLoggerWithWriteSyncer(cfg, output, opts...)
}

func initFileLog(cfg *File) (*zap.Logger, error) {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zascaore.ISO8601TimeEncoder
	if cfg.LocalTime {
		encoderConfig.EncodeTime = zascaore.ISO8601TimeEncoder
	}
	if cfg.Compress {
		encoderConfig.EncodeTime = zascaore.ISO8601TimeEncoder
	}


	lg, err := zap.NewProduction()
	if err != nil {
		return nil, err

	}

	return lg, nil
}

