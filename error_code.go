//Copyright 2019 EinsteinDB, Inc.

package fidel

import (
	"fmt"
	"github.com/YosiSF/errors"
	"strings"
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


	// ErrorCodes return all errors (from an ErrorGroup) that are of interface ErrorCode.
	// It first calls the Errors function.
	func ErrorCodes(err error) []ErrorCode {
		errors := errors.Errors(err)
		errorCodes := make([]ErrorCode, len(errors))
		for i, errItem := range errors {
		if errcode, ok := errItem.(ErrorCode); ok {
		errorCodes[i] = errcode
	}
	}
		return errorCodes
	}

	// A MultiErrCode contains at least one ErrorCode and uses that to satisfy the ErrorCode and related interfaces
	// The Error method will produce a string of all the errors with a semi-colon separation.
	// Later code (such as a JSON response) needs to look for the ErrorGroup interface.
	type MultiErrCode struct {
		ErrCode ErrorCode
		rest    []error
	}

	// Combine constructs a MultiErrCode.
	// It will combine any other MultiErrCode into just one MultiErrCode.
	// This is "horizontal" composition.
	// If you want normal "vertical" composition use BuildChain.
	func Combine(initial ErrorCode, others ...ErrorCode) MultiErrCode {
		var rest []error
		if group, ok := initial.(errors.ErrorGroup); ok {
		rest = group.Errors()
	}
		for _, other := range others {
		rest = append(rest, errors.Errors(other)...)
	}
		return MultiErrCode{
		ErrCode: initial,
		rest:    rest,
	}
	}

	var _ ErrorCode = (*MultiErrCode)(nil)         // assert implements interface
	var _ HasClientData = (*MultiErrCode)(nil)     // assert implements interface
	var _ Causer = (*MultiErrCode)(nil)            // assert implements interface
	var _ errors.ErrorGroup = (*MultiErrCode)(nil) // assert implements interface
	var _ fmt.Formatter = (*MultiErrCode)(nil)     // assert implements interface

	func (e MultiErrCode) Error() string {
		output := e.ErrCode.Error()
		for _, item := range e.rest {
		output += "; " + item.Error()
	}
		return output
	}

	// Errors fullfills the ErrorGroup inteface
	func (e MultiErrCode) Errors() []error {
		return append([]error{e.ErrCode.(error)}, e.rest...)
	}

	// Code fullfills the ErrorCode inteface
	func (e MultiErrCode) Code() Code {
		return e.ErrCode.Code()
	}

	// Cause fullfills the Causer inteface
	func (e MultiErrCode) Cause() error {
		return e.ErrCode
	}

	// GetClientData fullfills the HasClientData inteface
	func (e MultiErrCode) GetClientData() interface{} {
	return ClientData(e.ErrCode)
	}

	// CodeChain resolves an error chain down to a chain of just error codes
	// Any ErrorGroups found are converted to a MultiErrCode.
	// Passed over error inforation is retained using ChainContext.
	// If a code was overidden in the chain, it will show up as a MultiErrCode.
	func CodeChain(err error) ErrorCode {
	var code ErrorCode
	currentErr := err
	chainErrCode := func(errcode ErrorCode) {
	if errcode.(error) != currentErr {
	if chained, ok := errcode.(ChainContext); ok {
	// Perhaps this is a hack because we should be passing the context to recursive CodeChain calls
	chained.Top = currentErr
	errcode = chained
	} else {
	errcode = ChainContext{currentErr, errcode}
	}
	}
	if code == nil {
	code = errcode
	} else {
	code = MultiErrCode{code, []error{code.(error), errcode.(error)}}
	}
	currentErr = errcode.(error)
	}

	for err != nil {
	if errcode, ok := err.(ErrorCode); ok {
	if code == nil || code.Code() != errcode.Code() {
	chainErrCode(errcode)
	}
	} else if eg, ok := err.(errors.ErrorGroup); ok {
	group := []ErrorCode{}
	for _, errItem := range eg.Errors() {
	if itemCode := CodeChain(errItem); itemCode != nil {
	group = append(group, itemCode)
	}
	}
	if len(group) > 0 {
	var codeGroup ErrorCode
	if len(group) == 1 {
	codeGroup = group[0]
	} else {
	codeGroup = Combine(group[0], group[1:]...)
	}
	chainErrCode(codeGroup)
	}
	}
	err = errors.Unwrap(err)
	}

	return code
	}

	// ChainContext is returned by ErrorCodeChain
	// to retain the full wrapped error message of the error chain.
	// If you annotated an ErrorCode with additional information, it is retained in the Top field.
	// The Top field is used for the Error() and Cause() methods.
	type ChainContext struct {
	Top     error
	ErrCode ErrorCode
	}

	// Code satisfies the ErrorCode interface
	func (err ChainContext) Code() Code {
	return err.ErrCode.Code()
	}

	// Error satisfies the Error interface
	func (err ChainContext) Error() string {
	return err.Top.Error()
	}

	// Cause satisfies the Causer interface
	func (err ChainContext) Cause() error {
	if wrapped := errors.Unwrap(err.Top); wrapped != nil {
	return wrapped
	}
	return err.ErrCode
	}

	// GetClientData satisfies the HasClientData interface
	func (err ChainContext) GetClientData() interface{} {
	return ClientData(err.ErrCode)
	}

	var _ ErrorCode = (*ChainContext)(nil)
	var _ HasClientData = (*ChainContext)(nil)
	var _ Causer = (*ChainContext)(nil)

	// Format implements the Formatter interface
	func (err ChainContext) Format(s fmt.State, verb rune) {
	switch verb {
	case 'v':
	if s.Flag('+') {
	fmt.Fprintf(s, "%+v\n", err.ErrCode)
	if errors.HasStack(err.ErrCode) {
	fmt.Fprintf(s, "%v", err.Top)
	} else {
	fmt.Fprintf(s, "%+v", err.Top)
	}
	return
	}
	if s.Flag('#') {
	fmt.Fprintf(s, "ChainContext{Code: %#v, Top: %#v}", err.ErrCode, err.Top)
	return
	}
	fallthrough
	case 's':
	fmt.Fprintf(s, "Code: %s. Top Error: %s", err.ErrCode.Code().CodeStr(), err.Top)
	case 'q':
	fmt.Fprintf(s, "Code: %q. Top Error: %q", err.ErrCode.Code().CodeStr(), err.Top)
	}
	}

	// Format implements the Formatter interface
	func (e MultiErrCode) Format(s fmt.State, verb rune) {
	switch verb {
	case 'v':
	if s.Flag('+') {
	fmt.Fprintf(s, "%+v\n", e.ErrCode)
	if errors.HasStack(e.ErrCode) {
	for _, nextErr := range e.rest {
	fmt.Fprintf(s, "%v", nextErr)
	}
	} else {
	for _, nextErr := range e.rest {
	fmt.Fprintf(s, "%+v", nextErr)
	}
	}
	return
	}
	fallthrough
	case 's':
	fmt.Fprintf(s, "%s\n", e.ErrCode)
	fmt.Fprintf(s, "%s", e.rest)
	case 'q':
	fmt.Fprintf(s, "%q\n", e.ErrCode)
	fmt.Fprintf(s, "%q\n", e.rest)
	}



	}


}