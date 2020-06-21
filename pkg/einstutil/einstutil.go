//Copyright 2020 WHTCORPS INC ALL RIGHTS RESERVED.
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

package einstutil

import (
	"encoding/json"
	"fmt"
	"sync"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/YosiSF/FIDel/errcode"
	"github.com/YosiSF/FIDel/log"
	"github.com/YosiSF/pkg/errors"
	"github.com/unrolled/render"
)

func DeferClose(c io.Closer, err *error) {
	if cerr := c.Close(); cerr != nil && *err == nil {
		*err = errors.WithStack(cerr)
	}
}

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


func ReadJSON(r io.ReadCloser, data interface{}) error {
	var err error
	defer DeferClose(r, &err)
	b, err := ioutil.ReadAll(r)
	if err != nil {
		return errors.WithStack(err)
	}

	err = json.Unmarshal(b, data)
	if err != nil {
		return tagJSONError(err)
	}

	return err
}

// FieldError connects an error to a particular field
type FieldError struct {
	error
	field string
}


func ParseUint64VarsField(vars map[string]string, varName string) (uint64, *FieldError) {
	str, ok := vars[varName]
	if !ok {
		return 0, &FieldError{field: varName, error: fmt.Errorf("field %s not present", varName)}
	}
	parsed, err := strconv.ParseUint(str, 10, 64)
	if err == nil {
		return parsed, nil
	}
	return parsed, &FieldError{field: varName, error: err}
}

// ReadJSONRespondError writes json into data.
// On error respond with a 400 Bad Request
func ReadJSONRespondError(rd *render.Render, w http.ResponseWriter, body io.ReadCloser, data interface{}) error {
	err := ReadJSON(body, data)
	if err == nil {
		return nil
	}
	var errCode errcode.ErrorCode
	if jsonErr, ok := errors.Cause(err).(JSONError); ok {
		errCode = errcode.NewInvalidInputErr(jsonErr.Err)
	} else {
		errCode = errcode.NewInternalErr(err)
	}
	ErrorResp(rd, w, errCode)
	return err
}

// ErrorResp Respond to the client about the given error, integrating with errcode.ErrorCode.
//
// Important: if the `err` is just an error and not an errcode.ErrorCode (given by errors.Cause),
// then by default an error is assumed to be a 500 Internal Error.
//
// If the error is nil, this also responds with a 500 and logs at the error level.
func ErrorResp(rd *render.Render, w http.ResponseWriter, err error) {
	if err == nil {
		log.Error("nil is given to errorResp")
		rd.JSON(w, http.StatusInternalServerError, "nil error")
		return
	}
	if errCode := errcode.CodeChain(err); errCode != nil {
		w.Header().Set("TiDB-Error-Code", errCode.Code().CodeStr().String())
		rd.JSON(w, errCode.Code().HTTPCode(), errcode.NewJSONFormat(errCode))
	} else {
		rd.JSON(w, http.StatusInternalServerError, err.Error())
	}
}

