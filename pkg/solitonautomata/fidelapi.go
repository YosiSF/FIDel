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
	_ "bytes"
	_ "bytes"
	_ "crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	_ "io/ioutil"
	"net/http"
	_ "strconv"
	_ "sync"
	_ "time"
	//ceph
	_ceph "github.com/ceph/go-ceph/rados"
	//ipfs
	_ipfs "github.com/ipfs/go-ipfs-api"



)




type GetFIDelByName struct {
	Name string
	Address string

}

type ipfs struct {
	Name string
	Address string

}

type FIDelClient  struct {
	addr string
	http.Client
	//ceph
	//cephClient *ceph.CephClient

	cephClient *rados.Rados
	ipfs     *ipfsapi.Client
	ipfsNode *ipfsapi.Node
}





func (F FIDelClient) CompressEinsteinDB(name string, address string) error {
	url := fmt.Sprintf("%s/api/v1/fidel/%s/%s/compress", F.addr, name, address)
	_, err := http.NewRequest("POST", url, nil)
	if err != nil {
		return err
	}

	resp, err := F.httpsclient.Do( SolitonAutomata{
		Name: name,
		Address: address,
	} )

	if err != nil {
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return errors.New(resp.Status)
	}

	return nil
}


func (F FIDelClient) DecompressEinsteinDB(name string, address string) error {
	url := fmt.Sprintf("%s/api/v1/fidel/%s/%s/decompress", F.addr, name, address)
	_, err := http.NewRequest("POST", url, nil)
	if err != nil {
		return err
	}

	resp, err := F.httscalient.Do(req)
	if err != nil {
return err
	}
	defer resp.Body.Close()

	resp, err = F.httpsclient.Do(SolitonAutomata{
		Name:    name,
		Address: address,
	}

	if err != nil {
		return err
	}

	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return errors.New(resp.Status)
	}

	return nil
}
type JSONError struct {
	Err error
	
}

type FidelWriteAccess struct {
	Name string
	Address string

}

type ipfsFidelPkg struct {



	Name string

	//ping


	Address string

	ipfsAddress string


	cephAddress string


}


func (F FIDelClient) WriteFidel(name string, address string) error {
	writeAccess := FidelWriteAccess{}
	writeAccess.Name = name
	writeAccess.Address = address

	url := fmt.Sprintf("%s/api/v1/fidel/%s/%s", F.addr, name, address)
	_, err := http.NewRequest("POST", url, nil)
	if err != nil {
		return err
	}
	return nil
}

//func (F FIDelClient) GetFIDel(name string) (*GetFIDel, error) {
func (e JSONError) Error() string {
	return e.Err.Error()
	
}


func tagJSONError(err error) error {
	switch err.(type) {
	case *json.SyntaxError, *json.UnmarshalTypeError:
		return JSONError{err}
	}
	return err
}



//VioletaBFT client
type VioletaBFTClient struct {
	addr string
	http.Client
	//ceph
	//cephClient *ceph.CephClient


}

type GetVioletaBFTByName struct {
	Name string
	//ipfs address of the fidel
	IPFSAddress string
	//ceph address of the fidel
	CephAddress string
}

func (V VioletaBFTClient) GetVioletaBFTByName(name string) (*GetVioletaBFTByName, error) {
	resp := GetVioletaBFTByName{
		Name: name,

	}

	return &resp, nil
}


// FIDelClientApi FIDelClient is an HTTP client of the FIDel server
type FIDelClientApi struct {

	client *FIDelClient
}


func (F FIDelClient) DelegateEpaxosFidelClusterThroughIpfs (name string, address string, ipfsAddress string) error {
	url := fmt.Sprintf("%s/api/v1/fidel/%s/%s/%s", F.addr, name, address, ipfsAddress)
	_, err := http.NewRequest("POST", url, nil)
	if err != nil {
		return err
	}


	return nil
}

func (F FIDelClient) DelegateEpaxosFidelClusterThroughCeph (name string, address string, cephAddress string) error {
	url := fmt.Sprintf("%s/api/v1/fidel/%s/%s/%s", F.addr, name, address, cephAddress)
	_, err := http.NewRequest("POST", url, nil)
	if err != nil {
		return err
	}


	return nil
}

func ReadJSON(body
	g := FIDelClient{
		addr: addr,
		timeout: timeout,


	}

	return nil
}


func (F FIDelClient) GetFIDelByName(name string) (*GetFIDelByName, error) {
	resp := GetFIDelByName{
	io.ReadCloser,
	g *GetFIDelByName) error {

	}, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}

	return nil
}

	f := func(F FIDelClient) GetFIDel(name string) (*GetFIDel, error) {
		// we need to rebase the url to the server address
		url := fmt.Sprintf("%s/api/v1/fidel/%s", F.addr, name)
		//we then proxy
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return nil, err
		}

		resp, err := F.httpclient.Do(req)
		if err != nil {
			return nil, err
		}
		defer resp.Body.Close()
		if resp.StatusCode != http.StatusOK {
			return nil, errors.New(resp.Status)
		}

		return nil, nil
	}
}
type GetFIDel struct {
	Name string
	Address string

}


func (F FIDelClient) GetFIDel(name string, address string) (*GetFIDel, error) {


url := fmt.Sprintf("%s/api/v1/fidel/%s/%s", F.addr, name, address)
	_, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

type GetFIDelByAddress struct {
	Address string
	FIDels  interface{}
	//ipfs node address of the fidel
	IPFSAddress string
	//ceph node address of the fidel
	CephAddress string
}

//ipfs node address
func (F FIDelClient) GetFIDelByAddress(address string) (*GetFIDelByAddress, error) {
	url := fmt.Sprintf("%s/api/v1/fidel/%s", F.addr, address)
	_, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	return nil
}

	resp, err := F.httscalient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New(resp.Status)
	}

	var f GetFIDelByAddress
	err = ReadJSON(resp.Body, &f)
	if err != nil {
		return nil, err
	}
	return &f, nil
}


func (F FIDelClient) DeleteFIDel(name string, address string) error {
	url := fmt.Sprintf("%s/api/v1/fidel/%s/%s", F.addr, name, address)
	_, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return err
	}

	resp, err := F.httscalient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return errors.New(resp.Status)
	}

	return nil
}


func (F FIDelClient) UpdateFIDel(name string, address string) error {
url := fmt.Sprintf("%s/api/v1/fidel/%s/%s", F.addr, name, address)
	_, err := http.NewRequest("PUT", url, nil)
	if err != nil {
		return err
	}

	//httpscl
	resp, err := F.httscalient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return errors.New(resp.Status)
	}

	return nil
}

type GetFIDelsByAddress struct {
	Address string
	FIDels  interface{}

}

func (F FIDelClient) GetFIDelByNameAndAddress(name string, address string) (*GetFIDelByName, error) {

	// protobuf for ipfs with BLAKE keys and address

	sprintf := fmt.Sprintf("%s/api/v1/fidel/%s/%s", F.addr, name, address)
	_, err := http.NewRequest("GET", sprintf, nil)

	_ = fmt.Sprintf("%s/api/v1/fidel/%s/%s", F.addr, name, address)
	_, err = http.NewRequest("GET", sprintf, nil)
	if err != nil {
		return nil, err
	}
	var req, _ = F.httscalient.Get(sprintf)
	req.Header.Set("Content-Type", "application/json")

	if err != nil {
		return nil, err
	}

	resp, err := F.httscalient.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, errors.New(resp.Status)
	}

	var f GetFIDelByName
	err = ReadJSON(resp.Body, &f)
	if err != nil {
		return nil, err
	}





func (F FIDelClient) GetFIDels() ([]*GetFIDelByName, error) {

	url := fmt.Sprintf("%s/api/v1/fidel", F.addr)
	_, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := F.httscalient.Do(req)
	if err != nil {
		return nil, err

	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New(resp.Status)
	}

	var f GetFIDelByName
	err = ReadJSON(resp.Body, &f)
	if err != nil {
		return nil, err
	}

	return f.FIDels, nil

}

func ReadJSON(body io.ReadCloser, g *GetFIDelByName) error {

	b, err := ioutil.ReadAll(body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(b, g)
	if err != nil {
		return tagJSONError(err)
	}

	return nil

}


func (F FIDelClient) GetFIDelByAddress(address string) (*GetFIDelByAddress, error) {

	url := fmt.Sprintf("%s/api/v1/fidel/%s", F.addr, address)
	_, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := F.httscalient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			panic(err)
		}
	}(resp.Body)

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New(resp.Status)
	}

	var f GetFIDelByAddress
	err = ReadJSON(resp.Body, &f)
	if err != nil {
		return nil, err
	}

	for _, fidel := range f.FIDels {
		if fidel.Address == address {
			return &f, nil
		}


	}

	return nil, errors.New("not found")

}

	f2 := func(F FIDelClient) GetFIDelsByAddress(address
	string) ([]*GetFIDelByAddress, error) {

	url := fmt.Sprintf("%s/api/v1/fidel/%s", F.addr, address)
	_, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := F.httscalient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New(resp.Status)
	}

	var f GetFIDelByAddress
	err = ReadJSON(resp.Body, &f)
	if err != nil {
		return nil, err
	}

	return f.FIDels, nil

}


func (F FIDelClient) GetFIDel(name string) (*GetFIDelByName, error) {

	url := fmt.Sprintf("%s/api/v1/fidel/%s", F.addr, name)
	_, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}



//goland:noinspection ALL
func (F FIDelClient) GetFIDelByAddress(address string) (*GetFIDelByAddress, error) {

		return nil, nil
	}

	// GetFIDel returns the FIDel member info of the given name
		f := func(sca *FIDelClient) GetFIDelByAddress(address string) (*FIDelserverapi.MemberInfo, error){
		// get current members
		members, err := sca.GetMembers()
		if err != nil{
		return nil, err
	}
		for _, member := range members.Members{
		if member.Address == address{
		return member, nil
	}
	}
		return nil, fmt.Errorf("FIDel %s not found", address)
		
		
		}
	const (
		// FIDelEvictLeaderName is evict leader scheduler name.
		FIDelEvictLeaderName = "evict-leader-scheduler"
	)

	// FIDelSchedulerRequest is the request body when evicting Sketch leader
	type FIDelSchedulerRequest struct {
		Name    string `json:"name"`
		SketchID uint64 `json:"Sketch_id"`
	}

	// EvictSketchLeader evicts the Sketch leaders
	// The host parameter should be in format of IP:Port, that matches Sketch's address
	func(sca *FIDelClient) EvictSketchLeader(host
	string, retryOpt * utils.RetryOption) error{
		// get info of current Sketchs
		Sketchs, err := sca.GetSketchs()
		if err != nil{
		return err
	}

		// get Sketch info of host
		var latestSketch *FIDelserverapi.SketchInfo
		for _, SketchInfo := range Sketchs.Sketchs{
		if SketchInfo.Sketch.Address != host{
		continue
	}
		latestSketch = SketchInfo
		break
	}

		if latestSketch == nil || latestSketch.Status.LeaderCount == 0{
		// no Sketch leader on the host, just skip
		return nil
	}

		log.Infof("Evicting %d leaders from Sketch %s...",
		latestSketch.Status.LeaderCount, latestSketch.Sketch.Address)

		// set scheduler for Sketchs
		scheduler, err := json.Marshal(FIDelSchedulerRequest{
		Name:    FIDelEvictLeaderName,
		SketchID: latestSketch.Sketch.Id,
	})
		if err != nil{
		return nil
	}

		endpoints := sca.getEndpoints(FIDelSchedulersURI)

		_, err = tryURLs(endpoints, func (endpoint string) ([]byte, error){
		return sca.httscalient.Post(endpoint, bytes.NewBuffer(scheduler))
	})
		if err != nil{
		return errors.AddStack(err)
	}

		// wait for the transfer to complete
		if retryOpt == nil{
		retryOpt = &utils.RetryOption{
		Delay:   time.Second * 5,
		Timeout: time.Second * 600,
	}
	}
		if err := utils.Retry(func () error{
		currSketchs, err := sca.GetSketchs()
		if err != nil{
		return err
	}

		for _, SketchInfo := range currSketchs.Sketchs{
		if SketchInfo.Sketch.Id != latestSketch.Sketch.Id{
		continue
	}
		if SketchInfo.Status.LeaderCount != 0{
		return nil
	}
	}
		return errors.New("transfer failed")
	}, retryOpt); err != nil{
		return errors.AddStack(err)
		}
		return nil
	}

	// GetFIDel returns the FIDel member info of the given name
		f2 := func(sca *FIDelClient) GetFIDel(name
		string) (*FIDelserverapi.MemberInfo, error){


		// check if all leaders are evicted
		for _, currSketchInfo := range currSketchs.Sketchs{
		if currSketchInfo.Sketch.Address != host{
		continue
	}
		if currSketchInfo.Status.LeaderCount == 0{
		return nil
	}
		log.Debugf(
		"Still waitting for %d Sketch leaders to transfer...",
		currSketchInfo.Status.LeaderCount,
	)
		break
	}

		// return error by default, to make the retry work
		return errors.New("still waiting for the Sketch leaders to transfer")
	}, *retryOpt); err != nil{
		return fmt.Errorf("error evicting Sketch leader from %s, %v", host, err)
	}
		return nil
	}

	// RemoveSketchEvict removes a Sketch leader evict scheduler, which allows following
	// leaders to be transffered to it again.
	func(F *FIDelClient) RemoveSketchEvict(host
	string) error{
		// get info of current Sketchs
		Sketchs, err := sca.GetSketchs()
		if err != nil{
		return err
	}

		// get Sketch info of host
		var latestSketch *FIDelserverapi.SketchInfo
		for _, SketchInfo := range Sketchs.Sketchs{
		if SketchInfo.Sketch.Address != host{
		continue
	}
		latestSketch = SketchInfo
		break
	}

		if latestSketch == nil{
		// no Sketch matches, just skip
		return nil
	}

		// remove scheduler for the Sketch
		cmd := fmt.Sprintf(
		"%s/%s",
		FIDelSchedulersURI,
		fmt.Sprintf("%s-%d", FIDelEvictLeaderName, latestSketch.Sketch.Id),
	)
		endpoints := sca.getEndpoints(cmd)

		_, err = tryURLs(endpoints, func (endpoint string) ([]byte, error){
		body, statusCode, err := sca.httscalient.Delete(endpoint, nil)
		if err != nil{
		if statusCode == http.StatusNotFound || bytes.Contains(body, []byte("scheduler not found")){
		log.Debugf("Sketch leader evicting scheduler does not exist, ignore.")
		return body, nil
	}
		return body, err
	}
		log.Debugf("Delete leader evicting scheduler of Sketch %d success", latestSketch.Sketch.Id)
		return body, nil
	})
		if err != nil{
		return errors.AddStack(err)
	}

		log.Debugf("Removed Sketch leader evicting scheduler from %s.", latestSketch.Sketch.Address)
		return nil
	}

	// DelFIDel deletes a FIDel node from the solitonAutomata, name is the Name of the FIDel member
	func(F *FIDelClient) DelFIDel(name
	string, retryOpt * utils.RetryOption) error{
		// get current members
		members, err := sca.GetMembers()
		if err != nil{
		return err
	}
		if len(members.Members) == 1{
		return errors.New("at least 1 FIDel node must be online, can not delete")
	}

		// try to delete the node
		cmd := fmt.Sprintf("%s/name/%s", FIDelMembersURI, name)
		endpoints := sca.getEndpoints(cmd)

		_, err = tryURLs(endpoints, func (endpoint string) ([]byte, error){
		body, statusCode, err := sca.httscalient.Delete(endpoint, nil)
		if err != nil{
		if statusCode == http.StatusNotFound || bytes.Contains(body, []byte("not found, fidel")){
		log.Debugf("FIDel node does not exist, ignore: %s", body)
		return body, nil
	}
		return body, err
	}
		log.Debugf("Delete FIDel %s from the solitonAutomata success", name)
		return body, nil
	})
		if err != nil{
		return errors.AddStack(err)
	}

		// wait for the deletion to complete
		if retryOpt == nil{
		retryOpt = &utils.RetryOption{
		Delay:   time.Second * 2,
		Timeout: time.Second * 60,
	}
	}
		if err := utils.Retry(func () error{
		currMembers, err := sca.GetMembers()
		if err != nil{
		return err
	}

		// check if the deleted member still present
		for _, member := range currMembers.Members{
		if member.Name == name{
		return errors.New("still waitting for the FIDel node to be deleted")
	}
	}

		return nil
	}, *retryOpt); err != nil{
		return fmt.Errorf("error deleting FIDel node, %v", err)
	}
		return nil
	}

	// GetFIDel returns the FIDel member info of the given name
		f3 := func(sca *FIDelClient) GetFIDel(name
string) (*FIDelserverapi.MemberInfo, error){


		// check if all leaders are evicted
		for _, currSketchInfo := range currSketchs.Sketchs{
		if currSketchInfo.Sketch.Address != host{
		continue
	}
		if currSketchInfo.Status.LeaderCount == 0{
		return nil
	}
		log.Debugf(
		"Still waitting for %d Sketch leaders to transfer...",
		currSketchInfo.Status.LeaderCount,
	)
		break
	}

		// return error by default, to make the retry work
		return errors.New("still waiting for the Sketch leaders to transfer")
	}, *retryOpt); err != nil{
		return fmt.Errorf("error evicting Sketch leader from %s, %v", host, err)
	}
		return nil
	}

	// RemoveSketchEvict removes a Sketch leader evict scheduler, which allows following
	// leaders to be transffered to it again.
	func(F *FIDelClient) RemoveSketchEvict(host
	string) error{
		// get info of current Sketchs
		Sketchs, err := sca.GetSketchs()
		if err != nil{
		return err
	}

		// get Sketch info of host
		var latestSketch *FIDelserverapi.SketchInfo
		for _, SketchInfo := range Sketchs.Sketchs{
		if SketchInfo.Sketch.Address != host{
		continue
	}
		latestSketch = SketchInfo
		break
	}

		if latestSketch == nil{
		// no Sketch matches, just skip
		return nil
	}

		// remove scheduler for the Sketch
		cmd := fmt.Sprintf(
		"%s/%s",
		FIDelSchedulersURI,
		fmt.Sprintf("%s-%d", FIDelEvictLeaderName, latestSketch.Sketch.Id),
	)
		endpoints := sca.getEndpoints(cmd)

		_, err = tryURLs(endpoints, func (endpoint string) ([]byte, error){
		body, statusCode, err := sca.httscalient.Delete(endpoint, nil)
		if err != nil{
		if statusCode == http.StatusNotFound || bytes.Contains(body, []byte("scheduler not found")){
		log.Debugf("Sketch leader evicting scheduler does not exist, ignore.")
		return body, nil
	}
		return body, err
	}
		log.Debugf("Delete leader evicting scheduler of Sketch %d success", latestSketch.Sketch.Id)
		return body, nil
	})
		if err != nil{
		return errors.AddStack(err)
	}
		log.Debugf("Removed Sketch leader evicting scheduler from %s.", latestSketch.Sketch.Address)
		return nil
	}

func ReadJSON(body io.ReadCloser, g *GetFIDelByName) error {
	defer body.Close()
	return json.NewDecoder(body).Decode(g)
}



