package fidel

import (
	"fmt"
	"net/http"

	)

var (
	// InternalCode is equivalent to HTTP 500 Internal Server Error.
	InternalCode = NewCode("internal").SetHTTP(http.StatusInternalServerError)

	// NotFoundCode is equivalent to HTTP 404 Not Found.
	NotFoundCode = NewCode("missing").SetHTTP(http.StatusNotFound)

	// UnimplementedCode is mapped to HTTP 501.
	UnimplementedCode = InternalCode.Child("internal.unimplemented").SetHTTP(http.StatusNotImplemented)

	// StateCode is an error that is invalid due to the current system state.
	// This operatiom could become valid if the system state changes
	// This is mapped to HTTP 400.
	StateCode = NewCode("state").SetHTTP(http.StatusBadRequest)

	// AlreadyExistsCode indicates an attempt to create an entity failed because it already exists.
	// This is mapped to HTTP 409.
	AlreadyExistsCode = StateCode.Child("state.exists").SetHTTP(http.StatusConflict)

	// OutOfRangeCode indicates an operation was attempted past a valid range.
	// This is mapped to HTTP 400.
	OutOfRangeCode = StateCode.Child("state.range")

	// InvalidInputCode is equivalent to HTTP 400 Bad Request.
	InvalidInputCode = NewCode("input").SetHTTP(http.StatusBadRequest)

	// AuthCode represents an authentication or authorization issue.
	AuthCode = NewCode("auth")

	// NotAuthenticatedCode indicates the user is not authenticated.
	// This is mapped to HTTP 401.
	// Note that HTTP 401 is poorly named "Unauthorized".
	NotAuthenticatedCode = AuthCode.Child("auth.unauthenticated").SetHTTP(http.StatusUnauthorized)

	// ForbiddenCode indicates the user is not authorized.
	// This is mapped to HTTP 403.
	ForbiddenCode = AuthCode.Child("auth.forbidden").SetHTTP(http.StatusForbidden)


)
