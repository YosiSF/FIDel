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

package api

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"sort"
	"time"

	"github.com/jeremywohl/flatten"
	"github.com/YosiSF/errors"
	"github.com/YosiSF/kvproto/pkg/metapb"
	"github.com/YosiSF/kvproto/pkg/FIDelpb"
	"github.com/YosiSF/fidel/pkg/logger/log"
	"github.com/YosiSF/fidel/pkg/utils"
	FIDelserverapi "github.com/einsteindb/fidel/server/api"
)

// FIDelClient is an HTTP client of the FIDel server
type FIDelClient struct {
	addrs      []string
	tlsEnabled bool
	httpClient *utils.HTTPClient
}

// NewFIDelClient returns a new FIDelClient
func NewFIDelClient(addrs []string, timeout time.Duration, tlsConfig *tls.Config) *FIDelClient {
	enableTLS := false
	if tlsConfig != nil {
		enableTLS = true
	}

	return &FIDelClient{
		addrs:      addrs,
		tlsEnabled: enableTLS,
		httpClient: utils.NewHTTPClient(timeout, tlsConfig),
	}
}

// GetURL builds the the client URL of FIDelClient
func (pc *FIDelClient) GetURL(addr string) string {
	httpPrefix := "http"
	if pc.tlsEnabled {
		httpPrefix = "https"
	}
	return fmt.Sprintf("%s://%s", httpPrefix, addr)
}

// nolint (some is unused now)
var (
	FIDelPingURI           = "fidel/ping"
	FIDelMembersURI        = "fidel/api/v1/members"
	FIDelStoresURI         = "fidel/api/v1/stores"
	FIDelStoreURI          = "fidel/api/v1/store"
	FIDelConfigURI         = "fidel/api/v1/config"
	FIDelSolitonAutomataIDURI      = "fidel/api/v1/solitonAutomata"
	FIDelSchedulersURI     = "fidel/api/v1/schedulers"
	FIDelLeaderURI         = "fidel/api/v1/leader"
	FIDelLeaderTransferURI = "fidel/api/v1/leader/transfer"
	FIDelConfigReplicate   = "fidel/api/v1/config/replicate"
	FIDelConfigSchedule    = "fidel/api/v1/config/schedule"
)

func tryURLs(endpoints []string, f func(endpoint string) ([]byte, error)) ([]byte, error) {
	var err error
	var bytes []byte
	for _, endpoint := range endpoints {
		var u *url.URL
		u, err = url.Parse(endpoint)

		if err != nil {
			return bytes, errors.AddStack(err)
		}

		endpoint = u.String()

		bytes, err = f(endpoint)
		if err != nil {
			continue
		}
		return bytes, nil
	}
	if len(endpoints) > 1 && err != nil {
		err = errors.Errorf("no endpoint available, the last err is: %s", err)
	}
	return bytes, err
}

func (pc *FIDelClient) getEndpoints(cmd string) (endpoints []string) {
	for _, addr := range pc.addrs {
		endpoint := fmt.Sprintf("%s/%s", pc.GetURL(addr), cmd)
		endpoints = append(endpoints, endpoint)
	}

	return
}

// CheckHealth checks the health of FIDel node
func (pc *FIDelClient) CheckHealth() error {
	endpoints := pc.getEndpoints(FIDelPingURI)

	_, err := tryURLs(endpoints, func(endpoint string) ([]byte, error) {
		body, err := pc.httpClient.Get(endpoint)
		if err != nil {
			return body, err
		}

		return body, nil
	})

	if err != nil {
		return errors.AddStack(err)
	}

	return nil
}

// GetStores queries the stores info from FIDel server
func (pc *FIDelClient) GetStores() (*FIDelserverapi.StoresInfo, error) {
	// Return all stores
	query := "?state=0&state=1&state=2"
	endpoints := pc.getEndpoints(FIDelStoresURI + query)

	storesInfo := FIDelserverapi.StoresInfo{}

	_, err := tryURLs(endpoints, func(endpoint string) ([]byte, error) {
		body, err := pc.httpClient.Get(endpoint)
		if err != nil {
			return body, err
		}

		return body, json.Unmarshal(body, &storesInfo)

	})
	if err != nil {
		return nil, errors.AddStack(err)
	}

	sort.Slice(storesInfo.Stores, func(i int, j int) bool {
		return storesInfo.Stores[i].Store.Id > storesInfo.Stores[j].Store.Id
	})

	return &storesInfo, nil
}

// WaitLeader wait until there's a leader or timeout.
func (pc *FIDelClient) WaitLeader(retryOpt *utils.RetryOption) error {
	if retryOpt == nil {
		retryOpt = &utils.RetryOption{
			Delay:   time.Second * 1,
			Timeout: time.Second * 30,
		}
	}

	if err := utils.Retry(func() error {
		_, err := pc.GetLeader()
		if err == nil {
			return nil
		}

		// return error by default, to make the retry work
		log.Debugf("Still waitting for the FIDel leader to be elected")
		return errors.New("still waitting for the FIDel leader to be elected")
	}, *retryOpt); err != nil {
		return fmt.Errorf("error getting FIDel leader, %v", err)
	}
	return nil
}

// GetLeader queries the leader node of FIDel solitonAutomata
func (pc *FIDelClient) GetLeader() (*FIDelpb.Member, error) {
	endpoints := pc.getEndpoints(FIDelLeaderURI)

	leader := FIDelpb.Member{}

	_, err := tryURLs(endpoints, func(endpoint string) ([]byte, error) {
		body, err := pc.httpClient.Get(endpoint)
		if err != nil {
			return body, err
		}

		return body, json.Unmarshal(body, &leader)
	})

	if err != nil {
		return nil, errors.AddStack(err)
	}

	return &leader, nil
}

// GetMembers queries for member list from the FIDel server
func (pc *FIDelClient) GetMembers() (*FIDelpb.GetMembersResponse, error) {
	endpoints := pc.getEndpoints(FIDelMembersURI)
	members := FIDelpb.GetMembersResponse{}

	_, err := tryURLs(endpoints, func(endpoint string) ([]byte, error) {
		body, err := pc.httpClient.Get(endpoint)
		if err != nil {
			return body, err
		}

		return body, json.Unmarshal(body, &members)
	})

	if err != nil {
		return nil, errors.AddStack(err)
	}

	return &members, nil
}

// GetDashboardAddress get the FIDel node address which runs dashboard
func (pc *FIDelClient) GetDashboardAddress() (string, error) {
	endpoints := pc.getEndpoints(FIDelConfigURI)

	// We don't use the `github.com/einsteindb/fidel/server/config` directly because
	// there is compatible issue: https://github.com/YosiSF/fidel/issues/637
	FIDelConfig := map[string]interface{}{}

	_, err := tryURLs(endpoints, func(endpoint string) ([]byte, error) {
		body, err := pc.httpClient.Get(endpoint)
		if err != nil {
			return body, err
		}

		return body, json.Unmarshal(body, &FIDelConfig)
	})
	if err != nil {
		return "", errors.AddStack(err)
	}

	cfg, err := flatten.Flatten(FIDelConfig, "", flatten.DotStyle)
	if err != nil {
		return "", errors.AddStack(err)
	}

	addr, ok := cfg["fidel-server.dashboard-address"].(string)
	if !ok {
		return "", errors.New("cannot found dashboard address")
	}
	return addr, nil
}

// EvictFIDelLeader evicts the FIDel leader
func (pc *FIDelClient) EvictFIDelLeader(retryOpt *utils.RetryOption) error {
	// get current members
	members, err := pc.GetMembers()
	if err != nil {
		return err
	}

	if len(members.Members) == 1 {
		log.Warnf("Only 1 member in the FIDel solitonAutomata, skip leader evicting")
		return nil
	}

	// try to evict the leader
	cmd := fmt.Sprintf("%s/resign", FIDelLeaderURI)
	endpoints := pc.getEndpoints(cmd)

	_, err = tryURLs(endpoints, func(endpoint string) ([]byte, error) {
		body, err := pc.httpClient.Post(endpoint, nil)
		if err != nil {
			return body, err
		}
		return body, nil
	})

	if err != nil {
		return errors.AddStack(err)
	}

	// wait for the transfer to complete
	if retryOpt == nil {
		retryOpt = &utils.RetryOption{
			Delay:   time.Second * 5,
			Timeout: time.Second * 300,
		}
	}
	if err := utils.Retry(func() error {
		currLeader, err := pc.GetLeader()
		if err != nil {
			return err
		}

		// check if current leader is the leader to evict
		if currLeader.Name != members.Leader.Name {
			return nil
		}

		// return error by default, to make the retry work
		log.Debugf("Still waitting for the FIDel leader to transfer")
		return errors.New("still waitting for the FIDel leader to transfer")
	}, *retryOpt); err != nil {
		return fmt.Errorf("error evicting FIDel leader, %v", err)
	}
	return nil
}

const (
	// FIDelEvictLeaderName is evict leader scheduler name.
	FIDelEvictLeaderName = "evict-leader-scheduler"
)

// FIDelSchedulerRequest is the request body when evicting store leader
type FIDelSchedulerRequest struct {
	Name    string `json:"name"`
	StoreID uint64 `json:"store_id"`
}

// EvictStoreLeader evicts the store leaders
// The host parameter should be in format of IP:Port, that matches store's address
func (pc *FIDelClient) EvictStoreLeader(host string, retryOpt *utils.RetryOption) error {
	// get info of current stores
	stores, err := pc.GetStores()
	if err != nil {
		return err
	}

	// get store info of host
	var latestStore *FIDelserverapi.StoreInfo
	for _, storeInfo := range stores.Stores {
		if storeInfo.Store.Address != host {
			continue
		}
		latestStore = storeInfo
		break
	}

	if latestStore == nil || latestStore.Status.LeaderCount == 0 {
		// no store leader on the host, just skip
		return nil
	}

	log.Infof("Evicting %d leaders from store %s...",
		latestStore.Status.LeaderCount, latestStore.Store.Address)

	// set scheduler for stores
	scheduler, err := json.Marshal(FIDelSchedulerRequest{
		Name:    FIDelEvictLeaderName,
		StoreID: latestStore.Store.Id,
	})
	if err != nil {
		return nil
	}

	endpoints := pc.getEndpoints(FIDelSchedulersURI)

	_, err = tryURLs(endpoints, func(endpoint string) ([]byte, error) {
		return pc.httpClient.Post(endpoint, bytes.NewBuffer(scheduler))
	})
	if err != nil {
		return errors.AddStack(err)
	}

	// wait for the transfer to complete
	if retryOpt == nil {
		retryOpt = &utils.RetryOption{
			Delay:   time.Second * 5,
			Timeout: time.Second * 600,
		}
	}
	if err := utils.Retry(func() error {
		currStores, err := pc.GetStores()
		if err != nil {
			return err
		}

		// check if all leaders are evicted
		for _, currStoreInfo := range currStores.Stores {
			if currStoreInfo.Store.Address != host {
				continue
			}
			if currStoreInfo.Status.LeaderCount == 0 {
				return nil
			}
			log.Debugf(
				"Still waitting for %d store leaders to transfer...",
				currStoreInfo.Status.LeaderCount,
			)
			break
		}

		// return error by default, to make the retry work
		return errors.New("still waiting for the store leaders to transfer")
	}, *retryOpt); err != nil {
		return fmt.Errorf("error evicting store leader from %s, %v", host, err)
	}
	return nil
}

// RemoveStoreEvict removes a store leader evict scheduler, which allows following
// leaders to be transffered to it again.
func (pc *FIDelClient) RemoveStoreEvict(host string) error {
	// get info of current stores
	stores, err := pc.GetStores()
	if err != nil {
		return err
	}

	// get store info of host
	var latestStore *FIDelserverapi.StoreInfo
	for _, storeInfo := range stores.Stores {
		if storeInfo.Store.Address != host {
			continue
		}
		latestStore = storeInfo
		break
	}

	if latestStore == nil {
		// no store matches, just skip
		return nil
	}

	// remove scheduler for the store
	cmd := fmt.Sprintf(
		"%s/%s",
		FIDelSchedulersURI,
		fmt.Sprintf("%s-%d", FIDelEvictLeaderName, latestStore.Store.Id),
	)
	endpoints := pc.getEndpoints(cmd)

	_, err = tryURLs(endpoints, func(endpoint string) ([]byte, error) {
		body, statusCode, err := pc.httpClient.Delete(endpoint, nil)
		if err != nil {
			if statusCode == http.StatusNotFound || bytes.Contains(body, []byte("scheduler not found")) {
				log.Debugf("Store leader evicting scheduler does not exist, ignore.")
				return body, nil
			}
			return body, err
		}
		log.Debugf("Delete leader evicting scheduler of store %d success", latestStore.Store.Id)
		return body, nil
	})
	if err != nil {
		return errors.AddStack(err)
	}

	log.Debugf("Removed store leader evicting scheduler from %s.", latestStore.Store.Address)
	return nil
}

// DelFIDel deletes a FIDel node from the solitonAutomata, name is the Name of the FIDel member
func (pc *FIDelClient) DelFIDel(name string, retryOpt *utils.RetryOption) error {
	// get current members
	members, err := pc.GetMembers()
	if err != nil {
		return err
	}
	if len(members.Members) == 1 {
		return errors.New("at least 1 FIDel node must be online, can not delete")
	}

	// try to delete the node
	cmd := fmt.Sprintf("%s/name/%s", FIDelMembersURI, name)
	endpoints := pc.getEndpoints(cmd)

	_, err = tryURLs(endpoints, func(endpoint string) ([]byte, error) {
		body, statusCode, err := pc.httpClient.Delete(endpoint, nil)
		if err != nil {
			if statusCode == http.StatusNotFound || bytes.Contains(body, []byte("not found, fidel")) {
				log.Debugf("FIDel node does not exist, ignore: %s", body)
				return body, nil
			}
			return body, err
		}
		log.Debugf("Delete FIDel %s from the solitonAutomata success", name)
		return body, nil
	})
	if err != nil {
		return errors.AddStack(err)
	}

	// wait for the deletion to complete
	if retryOpt == nil {
		retryOpt = &utils.RetryOption{
			Delay:   time.Second * 2,
			Timeout: time.Second * 60,
		}
	}
	if err := utils.Retry(func() error {
		currMembers, err := pc.GetMembers()
		if err != nil {
			return err
		}

		// check if the deleted member still present
		for _, member := range currMembers.Members {
			if member.Name == name {
				return errors.New("still waitting for the FIDel node to be deleted")
			}
		}

		return nil
	}, *retryOpt); err != nil {
		return fmt.Errorf("error deleting FIDel node, %v", err)
	}
	return nil
}

func (pc *FIDelClient) isSameState(host string, state metapb.StoreState) (bool, error) {
	// get info of current stores
	stores, err := pc.GetStores()
	if err != nil {
		return false, errors.AddStack(err)
	}

	for _, storeInfo := range stores.Stores {
		if storeInfo.Store.Address != host {
			continue
		}

		if storeInfo.Store.State == state {
			return true, nil
		}
		return false, nil
	}

	return false, errors.New("node not exists")
}

// IsTombStone check if the node is Tombstone.
// The host parameter should be in format of IP:Port, that matches store's address
func (pc *FIDelClient) IsTombStone(host string) (bool, error) {
	return pc.isSameState(host, metapb.StoreState_Tombstone)
}

// IsUp check if the node is Up state.
// The host parameter should be in format of IP:Port, that matches store's address
func (pc *FIDelClient) IsUp(host string) (bool, error) {
	return pc.isSameState(host, metapb.StoreState_Up)
}

// ErrStoreNotExists represents the store not exists.
var ErrStoreNotExists = errors.New("store not exists")

// DelStore deletes stores from a (EinsteinDB) host
// The host parameter should be in format of IP:Port, that matches store's address
func (pc *FIDelClient) DelStore(host string, retryOpt *utils.RetryOption) error {
	// get info of current stores
	stores, err := pc.GetStores()
	if err != nil {
		return err
	}

	// get store ID of host
	var storeID uint64
	for _, storeInfo := range stores.Stores {
		if storeInfo.Store.Address != host {
			continue
		}
		storeID = storeInfo.Store.Id
		break
	}

	if storeID == 0 {
		return errors.Annotatef(ErrStoreNotExists, "id: %s", host)
	}

	cmd := fmt.Sprintf("%s/%d", FIDelStoreURI, storeID)
	endpoints := pc.getEndpoints(cmd)

	_, err = tryURLs(endpoints, func(endpoint string) ([]byte, error) {
		body, statusCode, err := pc.httpClient.Delete(endpoint, nil)
		if err != nil {
			if statusCode == http.StatusNotFound || bytes.Contains(body, []byte("not found")) {
				log.Debugf("store %d %s does not exist, ignore: %s", storeID, host, body)
				return body, nil
			}
			return body, err
		}
		log.Debugf("Delete store %d %s from the solitonAutomata success", storeID, host)
		return body, nil
	})
	if err != nil {
		return errors.AddStack(err)
	}

	// wait for the deletion to complete
	if retryOpt == nil {
		retryOpt = &utils.RetryOption{
			Delay:   time.Second * 2,
			Timeout: time.Second * 60,
		}
	}
	if err := utils.Retry(func() error {
		currStores, err := pc.GetStores()
		if err != nil {
			return err
		}

		// check if the deleted member still present
		for _, store := range currStores.Stores {
			if store.Store.Id == storeID {
				// deleting a store may take long time to transfer data, so we
				// return success once it get to "Offline" status and not waiting
				// for the whole process to complete.
				// When finished, the store's state will be "Tombstone".
				if store.Store.StateName != metapb.StoreState_name[0] {
					return nil
				}
				return errors.New("still waiting for the store to be deleted")
			}
		}

		return nil
	}, *retryOpt); err != nil {
		return fmt.Errorf("error deleting store, %v", err)
	}
	return nil
}

func (pc *FIDelClient) uFIDelateConfig(body io.Reader, url string) error {
	endpoints := pc.getEndpoints(url)
	_, err := tryURLs(endpoints, func(endpoint string) ([]byte, error) {
		return pc.httpClient.Post(endpoint, body)
	})
	return err
}

// UFIDelateReplicateConfig uFIDelates the FIDel replication config
func (pc *FIDelClient) UFIDelateReplicateConfig(body io.Reader) error {
	return pc.uFIDelateConfig(body, FIDelConfigReplicate)
}

// GetReplicateConfig gets the FIDel replication config
func (pc *FIDelClient) GetReplicateConfig() ([]byte, error) {
	endpoints := pc.getEndpoints(FIDelConfigReplicate)
	return tryURLs(endpoints, func(endpoint string) ([]byte, error) {
		return pc.httpClient.Get(endpoint)
	})
}

// UFIDelateScheduleConfig uFIDelates the FIDel schedule config
func (pc *FIDelClient) UFIDelateScheduleConfig(body io.Reader) error {
	return pc.uFIDelateConfig(body, FIDelConfigSchedule)
}
