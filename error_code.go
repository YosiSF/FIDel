//Copyright 2019 EinsteinDB, Inc.

package fidel

import (
	"fmt"
	"strings"
	"github.com/YosiSF/errors"
)


//to avoid merge conflicts build a representation of the type of a particular error

type CodeStr string

func (str CodeStr) String() string { return string(str) }

//attached to parent as metadata sink
type Code struct {
	codeStr CodeStr
	Parent *Code
}

//Create a new top-level code
func NewCode(codeRep CodeStr) Code {
	code := Code{codeStr: codeRep}
	if err := code.checkCodePath(); err != nil {
		panic(err)
	}

	return code
}


func (code Code) Child(childStr CodeStr) Code {
	child := Code{codeStr: childStr, Parent: &code}
	if err := child.checkCodePath(); err != nil {
		panic(err)
	}

	paths := strings.Split(child.codeStr.String(), ".")
	child.codeStr = CodeStr(paths[len(paths)-1])
	return child
 }

func (code Code) findAncestor(test func(Code) bool) *Code {
	if test(code) {
		return &code
	}
	if code.Parent == nil {
		return nil
	}
	return (*code.Parent).findAncestor(test)
}

// IsAncestor looks for the given code in its ancestors.
func (code Code) IsAncestor(ancestorCode Code) bool {
	return nil != code.findAncestor(func(an Code) bool { return an == ancestorCode })
}

type ErrorCode interface {
	Error() string // The Error interface
	Code() Code
}

type Causer interface {
	Cause() error
}

func ClientData(errCode ErrorCode) interface{} {
	var data interface{} = errCode
	if hasData, ok := errCode.(HasClientData); ok {
		data = hasData.GetClientData()
	}
	return data
}

type JSONFormat struct {
	Code      CodeStr           `json:"code"`
	Msg       string            `json:"msg"`
	Data      interface{}       `json:"data"`
	Operation string            `json:"operation,omitempty"`
	Stack     errors.StackTrace `json:"stack,omitempty"`
	Others    []JSONFormat      `json:"others,omitempty"`
}

func OperationClientData(errCode ErrorCode) (string, interface{}) {
	op := Operation(errCode)
	data := ClientData(errCode)
	if op == "" {
		op = Operation(data)
	}
	return op, data
}

func NewJSONFormat(errCode ErrorCode) JSONFormat {
	// Gather up multiple errors.
	// We discard any that are not ErrorCode.
	errorCodes := ErrorCodes(errCode)
	others := make([]JSONFormat, len(errorCodes)-1)
	for i, err := range errorCodes[1:] {
		others[i] = NewJSONFormat(err)
	}

	op, data := OperationClientData(errCode)

	var stack errors.StackTrace
	if errCode.Code().IsAncestor(InternalCode) {
		stack = StackTrace(errCode)
	}

	return JSONFormat{
		Data:      data,
		Msg:       errCode.Error(),
		Code:      errCode.Code().CodeStr(),
		Operation: op,
		Stack:     stack,
		Others:    others,
	}



}