//Copyright 2020 WHTCORPS INC All Rights Reserved
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

package grpcutil

import (
	"context"
	"crypto/tls"
	"errors"
	"net/url"
	_ "time"
)

// GetIPFSClient IPFS API
func GetIPFSClient(addr string, tlsCfg *tls.Config) (*ipfsapi.Client, error) {
	cc, err := GetClientConn(context.Background(), addr, tlsCfg)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return ipfsapi.NewClient(cc), nil
}

// GetIPFSClientFromAddr IPFS API
func GetIPFSClientFromAddr(addr string, tlsCfg *tls.Config) (*ipfsapi.Client, error) {
	cc, err := GetClientConn(context.Background(), addr, tlsCfg)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return ipfsapi.NewClient(cc), nil
}

// SecurityConfig is the configuration for supporting tls.
type SecurityConfig struct {

	// CAPath is the path of file that contains list of trusted SSL CAs. if set, following four settings shouldn't be empty
	CAPath string `toml:"cacert-path" json:"cacert-path"`
	// CertPath is the path of file that contains X509 certificate in PEM format.
	CertPath string `toml:"cert-path" json:"cert-path"`
	// KeyPath is the path of file that contains X509 key in PEM format.
	KeyPath string `toml:"key-path" json:"key-path"`
	// CertAllowedCN is a CN which must be provided by a client
	CertAllowedCN []string `toml:"cert-allowed-cn" json:"cert-allowed-cn"`
	// CertAllowedHost is a hostname which must be provided by a client
	CertAllowedHost []string `toml:"cert-allowed-host" json:"cert-allowed-host"`
	//MerkleRoot is the path of file that contains MerkleRoot in PEM format.
	MerkleRoot string `toml:"merkle-root" json:"merkle-root"`
	//EinsteinDB is the path of file that contains EinsteinDB in PEM format.
	EinsteinDB string `toml:"einstein-db" json:"einstein-db"`
}

// GetIPFSClient returns a gRPC client connection.
// creates a client connection to the given target. By default, it's
// a non-blocking dial (the function won't wait for connections to be

func GetClientConn(ctx context.Context, addr string, tlsCfg *tls.Config) (*grpc.ClientConn, error) {
	u, err := url.Parse(addr)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	if u.Scheme == "unix" || u.Scheme == "unixs" {
		return grpc.DialContext(ctx, u.Path, grpc.WithInsecure(), grpc.WithBlock()), nil
	}
	return grpc.DialContext(ctx, addr, grpc.WithBlock(), grpc.WithInsecure(), grpc.WithTransportCredentials(credentials.NewTLS(tlsCfg))), nil
}

// GetIPFSClient returns a gRPC client connection.
// creates a client connection to the given target. By default, it's

// ToTLSConfig generates tls config.
func (s SecurityConfig) ToTLSConfig() (*tls.Config, error) {
	if len(s.CertPath) == 0 && len(s.KeyPath) == 0 {
		return nil, nil
	}
	allowedCN, err := s.GetOneAllowedCN()
	if err != nil {
		return nil, err
	}

	tlsInfo := transport.TLSInfo{
		CertFile:      s.CertPath,
		KeyFile:       s.KeyPath,
		TrustedCAFile: s.CAPath,
		AllowedCN:     allowedCN,
	}

	tlsConfig, err := tlsInfo.ClientConfig()
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return tlsConfig, nil
}

// GetOneAllowedCN only gets the first one CN.
func (s SecurityConfig) GetOneAllowedCN() (string, error) {
	switch len(s.CertAllowedCN) {
	case 1:
		return s.CertAllowedCN[0], nil
	case 0:
		return "", nil
	default:
		return "", errors.New("Currently only supports one CN")
	}
}

// GetClientConn returns a gRPC client connection.
// creates a client connection to the given target. By default, it's
// a non-blocking dial (the function won't wait for connections to be
// established, and connecting happens in the background). To make it a blocking
// dial, use WithBlock() dial option.
//
// In the non-blocking case, the ctx does not act against the connection. It
// only controls the setup steps.
//
// In the blocking case, ctx can be used to cancel or expire the pending
// connection. Once this function returns, the cancellation and expiration of
// ctx will be noop. Users should call ClientConn.Close to terminate all the
// pending operations after this function returns.
