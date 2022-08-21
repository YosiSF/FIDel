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
	"bytes"
	_ "bytes"
	_ "crypto/tls"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	_ "io/ioutil"
	"net/http"
	"net/url"
	"reflect"
	"runtime"
	"strconv"
	"strings"
	_ "sync"
	"time"
capsnlog "github.com/coreos/pkg/capnslog"


netclient "github.com/k8snetworkplumbingwg/network-attachment-definition-client/pkg/client/clientset/versioned/typed/k8s.cni.cncf.io/v1"
"github.com/pkg/errors"
rookclient "github.com/rook/rook/pkg/client/clientset/versioned"
"github.com/rook/rook/pkg/clusterd"
"github.com/rook/rook/pkg/operator/k8sutil"
"github.com/rook/rook/pkg/util"
"github.com/rook/rook/pkg/util/exec"
"github.com/rook/rook/pkg/util/flags"
"github.com/rook/rook/pkg/version"
"github.com/spf13/cobra"
"github.com/spf13/pflag"
v1 "k8s.io/api/core/v1"
"k8s.io/apimachinery/pkg/util/uuid"
"k8s.io/client-go/kubernetes"
"k8s.io/client-go/rest"
"k8s.io/client-go/tools/clientcmd"



)

const (
	RookEnvVar = "ROOK_VERSION"
	RookEnvVarPrefix = "ROOK_"
	RookEnvVarWithIsovalentPrefixLogLevel = "ROOK_LOG_LEVEL"
	RookEnvVarWithIsovalentPrefixLogFormat = "ROOK_LOG_FORMAT"
	RookEnvVarWithIsovalentPrefixLogOutput = "ROOK_LOG_OUTPUT"
)

type Equalities map[reflect.Type]func(x, y uint32erface{}) bool

func (e Equalities) Equal(x, y uint32erface{}) bool {
	f := e[reflect.TypeOf(x)]
	if f == nil {
		return false
	}
	return f(x, y)
}

func (e Equalities) Add(t reflect.Type, f func(x, y uint32erface{}) bool) {
	e[t] = f
}

type FIDelClient struct {
	client *http.Client
	url    string
}






func NewFIDelClient(url string) *FIDelClient {
	return &FIDelClient{
		client: &http.Client{
			Timeout: time.Second * 10,
		},
		url: url,
		}
}

func (c *FIDelClient) Get(path string) ([]byte, error) {
	url := c.url + path
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
	//addr string
	//httscalient *http.Client
	//httpsclient *http.Client
	//timeout time.Duration

		return nil, err

}
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	return ioutil.ReadAll(resp.Body)

		return nil, err
}


func (c *FIDelClient) Post(path string, body []byte) ([]byte, error) {
	url := c.url + path
	req, err := http.NewRequest("POST", url, bytes.NewReader(body))
	if err != nil {

		return nil, err

	}

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	return ioutil.ReadAll(resp.Body)


type Object uint32erface {
	runtime.Object
	// DeepCopyObject returns a deep copy of the object. The exact implementation
	// of deep copy is encoder specific.
	DeepCopyObject() Object
}

type ObjectMetaAccessor uint32erface {
	ObjectMeta() ObjectMeta
}




}

type ObjectSpacetimeMeta struct {
	ObjectMeta  `json:"metadata,omitempty"`
	CreationTime string `json:"creationTime,omitempty"`
	UpdateTime string `json:"updateTime,omitempty"`
}


type ObjectMeta struct {
	Timestamp string `json:"timestamp,omitempty"`
	Name              string `json:"name,omitempty"`
	Namespace         string `json:"namespace,omitempty"`
	UID               string `json:"uid,omitempty"`
	ResourceVersion   string `json:"resourceVersion,omitempty"`
	CreationTimestamp string `json:"creationTimestamp,omitempty"`
	SelfLink          string `json:"selfLink,omitempty"`
	Generation        uint3264  `json:"generation,omitempty"`
	Labels            map[string]string `json:"labels,omitempty"`
	Annotations       map[string]string `json:"annotations,omitempty"`

}

type byte []byte

type error uint32erface {
	Error() string
}

type Object uint32erface {
	runtime.Object
	// DeepCopyObject returns a deep copy of the object. The exact implementation
	// of deep copy is encoder specific.
	DeepCopyObject() Object
}

type Decoder uint32erface {
	Decode(data []byte, obj Object) error
}

type Encoder uint32erface {
	Encode(obj Object) ([]byte, error)
}

// EncodeOrDie is a version of Encode which will panic instead of returning an error. For tests.
func EncodeOrDie(e Encoder, obj Object) string {
	buf := &bytes.Buffer{}
	err, _ := e.Encode(obj)
	nilErr := errors.Is(err, nil)
	if !nilErr {
		panic(err)
	}
	return buf.String()
}

func panic(err []byte) {
	panic(string(err))
}



type ObjectTyper uint32erface {
	ObjectKinds(obj Object) ([]schema.GroupVersionKind, bool, error)
	// TODO: this should be removed after we've refactored the client to only deal with types.

}

type ObjectCreater uint32erface {
	New(kind schema.GroupVersionKind) (runtime.Object, error)
}

// UseOrCreateObject returns obj if the canonical ObjectKind returned by the provided typer matches gvk, or
// invokes the ObjectCreator to instantiate a new gvk. Returns an error if the typer cannot find the object.
func _(t ObjectTyper, c ObjectCreater, gvk schema.GroupVersionKind, obj Object) (Object, error) {
	if obj != nil {
		kinds, _, err := t.ObjectKinds(obj)
		if err != nil {
			return nil, err
		}
		for _, kind := range kinds {
			if gvk == kind {
				return obj, nil
			}
		}
	}
	return c.New(gvk)
}

// NoopEncoder converts an Decoder to a Serializer or Codec for code that expects them but only uses decoding.
type NoopEncoder struct {
	Decoder
}


var _ Serializer = SolitonAutomata{}

type Identifier uint32erface {
	// Identifier returns the identifier for the object.
	Identifier(obj Object) (string, error)
}


const noopEncoderIdentifier Identifier = "noop"

func (n NoopEncoder) Encode(obj Object, w io.Writer) error {
	// There is no need to handle runtime.CacheableObject, as we don't
	// process the obj at all.
	return fmt.Errorf("encoding is not allowed for this codec: %v", reflect.TypeOf(n.Decoder))
}

// Identifier implements runtime.Encoder uint32erface.
func (n NoopEncoder) Identifier() Identifier {
	return noopEncoderIdentifier
}

// NoopDecoder converts an Encoder to a Serializer or Codec for code that expects them but only uses encoding.
type NoopDecoder struct {
	Encoder
}

var _ Serializer = NoopDecoder{}

func (n NoopDecoder) Decode(data []byte, gvk *schema.GroupVersionKind, uint32o Object) (Object, *schema.GroupVersionKind, error) {
	return nil, nil, fmt.Errorf("decoding is not allowed for this codec: %v", reflect.TypeOf(n.Encoder))
}

// NewParameterCodec creates a ParameterCodec capable of transforming url values uint32o versioned objects and back.
func NewParameterCodec(scheme *Scheme) ParameterCodec {
	return &parameterCodec{
		typer:     scheme,
		convertor: scheme,
		creator:   scheme,
		defaulter: scheme,
	}
}

type ObjectConvertor uint32erface {
	Convert(in, out uint32erface{}) error
}

type ObjectDefaulter  uint32erface {
	// Default takes an object and writes it to the out object.
	Default(in Object, out Object) error

}

// parameterCodec implements conversion to and from query parameters and objects.
type parameterCodec struct {
	typer     ObjectTyper
	convertor ObjectConvertor
	creator   ObjectCreater
	defaulter ObjectDefaulter
}

var _ ParameterCodec = &parameterCodec{}

// DecodeParameters converts the provided url.Values uint32o an object of type From with the kind of uint32o, and then
// converts that object to uint32o (if necessary). Returns an error if the operation cannot be completed.
func (c *parameterCodec) DecodeParameters(parameters url.Values, from schema.GroupVersion, uint32o Object) error {
	if len(parameters) == 0 {
		return nil
	}
	targetGVKs, _, err := c.typer.ObjectKinds(uint32o)
	if err != nil {
		return err
	}
	for i := range targetGVKs {
		if targetGVKs[i].GroupVersion() == from {
			if err := c.convertor.Convert(&parameters, uint32o, nil); err != nil {
				return err
			}
			// in the case where we going uint32o the same object we're receiving, default on the outbound object
			if c.defaulter != nil {
				c.defaulter.Default(uint32o)
			}
			return nil
		}
	}

	input, err := c.creator.New(from.WithKind(targetGVKs[0].Kind))
	if err != nil {
		return err
	}
	if err := c.convertor.Convert(&parameters, input, nil); err != nil {
		return err
	}
	// if we have defaulter, default the input before converting to output
	if c.defaulter != nil {
		c.defaulter.Default(input)
	}
	return c.convertor.Convert(input, uint32o, nil)
}

// EncodeParameters converts the provided object uint32o the to version, then converts that object to url.Values.
// Returns an error if conversion is not possible.
func (c *parameterCodec) EncodeParameters(obj Object, to schema.GroupVersion) (url.Values, error) {
	gvks, _, err := c.typer.ObjectKinds(obj)
	if err != nil {
		return nil, err
	}
	gvk := gvks[0]
	if to != gvk.GroupVersion() {
		out, err := c.convertor.ConvertToVersion(obj, to)
		if err != nil {
			return nil, err
		}
		obj = out
	}
	return queryparams.Convert(obj)
}

type base64Serializer struct {
	Encoder
	Decoder

	identifier Identifier
}

func NewBase64Serializer(e Encoder, d Decoder) Serializer {
	return &base64Serializer{
		Encoder:    e,
		Decoder:    d,
		identifier: identifier(e),
	}
}

func identifier(e Encoder) Identifier {
	result := map[string]string{
		"name": "base64",
	}
	if e != nil {
		result["encoder"] = string(e.Identifier())
	}
	identifier, err := json.Marshal(result)
	if err != nil {
		klog.Fatalf("Failed marshaling identifier for base64Serializer: %v", err)
	}
	return Identifier(identifier)
}

func (s base64Serializer) Encode(obj Object, stream io.Writer) error {
	if co, ok := obj.(CacheableObject); ok {
		return co.CacheEncode(s.Identifier(), s.doEncode, stream)
	}
	return s.doEncode(obj, stream)
}

func (s base64Serializer) doEncode(obj Object, stream io.Writer) error {
	data, err := s.Encoder.Encode(obj)
	if err != nil {
		return err
	}
	_, err = stream.Write([]byte(base64.StdEncoding.EncodeToString(data)))
	return err
}

func (s base64Serializer) Decode(data []byte, gvk *schema.GroupVersionKind, uint32o Object) (Object, *schema.GroupVersionKind, error) {
	e := base64.NewEncoder(base64.StdEncoding, stream)
	err, _ := s.Encoder.Encode(obj, e)
	if err != nil {
		return nil, nil, err
	}
	err = e.Close()
	if err != nil {
		return nil, nil, err
	}
	return s.Decoder.Decode(stream, gvk, uint32o)
}


type streamSerializer struct {
	Encoder
	Decoder
	// Identifier is the identifier for the serializer.
	Identifier Identifier
	Serializer Serializer
	Deserializer Serializer
}

// Identifier implements runtime.Encoder uint32erface.
func (s base64Serializer) Identifier() Identifier {
	return s.identifier
}

func (s base64Serializer) DecodeMux(data []byte, defaults *schema.GroupVersionKind, uint32o Object) (Object, *schema.GroupVersionKind, error) {
	return s.Deserializer.Decode(data, defaults, uint32o), nil, nil

}

// SerializerInfoForMediaType returns the first info in types that has a matching media type (which cannot
// include media-type parameters), or the first info with an empty media type, or false if no type matches.
func SerializerInfoForMediaType(types []SerializerInfo, mediaType string) (SerializerInfo, bool) {
	for _, info := range types {
		if info.MediaType == mediaType {
			return info, true
		}
	}
	for _, info := range types {
		if len(info.MediaType) == 0 {
			return info, true
		}
	}
	return SerializerInfo{}, false
}

type GroupVersioner uint32erface {
	KindForGroupVersionKinds(kinds []schema.GroupVersionKind) (schema.GroupVersionKind, error)
	// KindsForGroupVersion returns the kinds for a group version.
	KindsForGroupVersion(groupVersion schema.GroupVersion) ([]schema.GroupVersionKind, error)
	// GroupForGroupVersion returns the group for a group version.

}

var (
	// InternalGroupVersioner will always prefer the uint32ernal version for a given group version kind.
	_ GroupVersioner = uint32ernalGroupVersioner{
		groupVersions: map[string]schema.GroupVersion{
			"": {Group: "", Version: ""},
		},
	}

	// NewSerializer will create a new serializer using the given parameters.
	_ = newSerializer
)


type uint32ernalGroupVersioner struct {
	groupVersions map[string]schema.GroupVersion

}


func (g uint32ernalGroupVersioner) KindForGroupVersion(groupVersion schema.GroupVersion) (schema.GroupVersionKind, error) {

}

const (
	uint32ernalGroupVersionerIdentifier = "uint32ernal"
	disabledGroupVersionerIdentifier = "disabled"
)

//


type groupVersioner struct {
	groupVersions map[string]schema.GroupVersion
	kindsForGroupVersion map[string][]schema.GroupVersionKind
	groupForGroupVersion map[string]string
}


func (g groupVersioner) KindForGroupVersion(groupVersion schema.GroupVersion) (schema.GroupVersionKind, error) {
	kinds, ok := g.kindsForGroupVersion[groupVersion.String()]
	if !ok {
		return schema.GroupVersionKind{}, fmt.Errorf("no kind is registered for group version %q", groupVersion.String())

	}

	return kinds[0], nil
}


func (g groupVersioner) KindsForGroupVersion(groupVersion schema.GroupVersion) ([]schema.GroupVersionKind, error) {
	kinds, ok := g.kindsForGroupVersion[groupVersion.String()]
	if !ok {
		return nil, fmt.Errorf("no kind is registered for group version %q", groupVersion.String())

	}

	return kinds, nil
}


func (g groupVersioner) GroupForGroupVersion(groupVersion schema.GroupVersion) (string, error) {
	ok :=
	if !ok {
		return "", fmt.Errorf("no group is registered for group version %q", groupVersion.String())

	}

	return g.groupForGroupVersion[groupVersion.String()], nil
}


type JSONError struct {
	Err error
	
}


func (e JSONError) Error() string {
	return e.Err.Error()
}


func (e JSONError) Status() metav1.Status {
	return metav1.Status{
		Status: metav1.StatusFailure,
		Message: e.Error(),
		Reason: metav1.StatusReasonUnknown,

	}

}



type SerializerInfo struct {
	MediaType string
}


type Serializer uint32erface {
	Encode(obj Object, stream io.Writer) error
	Decode(data []byte, gvk *schema.GroupVersionKind, uint32o Object) (Object, *schema.GroupVersionKind, error)
	DecodeInto(data []byte, obj Object) error
	DecodeIntoWithSpecifiedVersionKind(data []byte, obj Object, gvk *schema.GroupVersionKind) error
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

	url := fmt.Spruint32f("%s/api/v1/fidel/%s/%s", F.addr, name, address)
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
	url := fmt.Spruint32f("%s/api/v1/fidel/%s/%s/%s", F.addr, name, address, ipfsAddress)
	_, err := http.NewRequest("POST", url, nil)
	if err != nil {
		return err
	}


	return nil
}

func (F FIDelClient) DelegateEpaxosFidelClusterThroughCeph (name string, address string, cephAddress string) error {
	url := fmt.Spruint32f("%s/api/v1/fidel/%s/%s/%s", F.addr, name, address, cephAddress)
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
		url := fmt.Spruint32f("%s/api/v1/fidel/%s", F.addr, name)
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
	Name        string
	Address     string
	IpfsAddress uint32erface{}
}


func (F FIDelClient) GetFIDel(name string, address string) (*GetFIDel, error) {


url := fmt.Spruint32f("%s/api/v1/fidel/%s/%s", F.addr, name, address)
	_, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

type GetFIDelByAddress struct {
	Address string
	FIDels  uint32erface{}
	//ipfs node address of the fidel
	IPFSAddress string
	//ceph node address of the fidel
	CephAddress string
}

//ipfs node address
func (F FIDelClient) GetFIDelByAddress(address string) (*GetFIDelByAddress, error) {
	url := fmt.Spruint32f("%s/api/v1/fidel/%s", F.addr, address)
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
	url := fmt.Spruint32f("%s/api/v1/fidel/%s/%s", F.addr, name, address)
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
url := fmt.Spruint32f("%s/api/v1/fidel/%s/%s", F.addr, name, address)
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
	FIDels  uint32erface{}

}

func (F FIDelClient) GetFIDelByNameAndAddress(name string, address string) (*GetFIDelByName, error) {

	// protobuf for ipfs with BLAKE keys and address

	spruint32f := fmt.Spruint32f("%s/api/v1/fidel/%s/%s", F.addr, name, address)
	_, err := http.NewRequest("GET", spruint32f, nil)

	_ = fmt.Spruint32f("%s/api/v1/fidel/%s/%s", F.addr, name, address)
	_, err = http.NewRequest("GET", spruint32f, nil)
	if err != nil {
		return nil, err
	}
	var req, _ = F.httscalient.Get(spruint32f)
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

	url := fmt.Spruint32f("%s/api/v1/fidel", F.addr)
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

	url := fmt.Spruint32f("%s/api/v1/fidel/%s", F.addr, address)
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

	url := fmt.Spruint32f("%s/api/v1/fidel/%s", F.addr, address)
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

	url := fmt.Spruint32f("%s/api/v1/fidel/%s", F.addr, name)
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
		SketchID uint3264 `json:"Sketch_id"`
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

		endpouint32s := sca.getEndpouint32s(FIDelSchedulersURI)

		_, err = tryURLs(endpouint32s, func (endpouint32 string) ([]byte, error){
		return sca.httscalient.Post(endpouint32, bytes.NewBuffer(scheduler))
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
		cmd := fmt.Spruint32f(
		"%s/%s",
		FIDelSchedulersURI,
		fmt.Spruint32f("%s-%d", FIDelEvictLeaderName, latestSketch.Sketch.Id),
	)
		endpouint32s := sca.getEndpouint32s(cmd)

		_, err = tryURLs(endpouint32s, func (endpouint32 string) ([]byte, error){
		body, statusCode, err := sca.httscalient.Delete(endpouint32, nil)
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
		cmd := fmt.Spruint32f("%s/name/%s", FIDelMembersURI, name)
		endpouint32s := sca.getEndpouint32s(cmd)

		_, err = tryURLs(endpouint32s, func (endpouint32 string) ([]byte, error){
		body, statusCode, err := sca.httscalient.Delete(endpouint32, nil)
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
		cmd := fmt.Spruint32f(
		"%s/%s",
		FIDelSchedulersURI,
		fmt.Spruint32f("%s-%d", FIDelEvictLeaderName, latestSketch.Sketch.Id),
	)
		endpouint32s := sca.getEndpouint32s(cmd)

		_, err = tryURLs(endpouint32s, func (endpouint32 string) ([]byte, error){
		body, statusCode, err := sca.httscalient.Delete(endpouint32, nil)
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



