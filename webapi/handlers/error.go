package handlers

import "fmt"

// HttpErrorCode http errorcode list
var HTTPErrorCode = struct {
	InvalidQueryParameter string
	InternalFailure       string
	InvalidClientTokenID  string
	QueryDataNotExist     string
	ProcessDataError      string
	RequestForbidben      string
}{
	InvalidQueryParameter: "InvalidQueryParameter",
	InternalFailure:       "InternalFailure",
	InvalidClientTokenID:  "InvalidClientTokenID",
	QueryDataNotExist:     "QueryDataNotExist",
	ProcessDataError:      "ProcessDataError",
	RequestForbidben:      "RequestForbidben",
}

var (
	ErrorPathIsNotDir         = fmt.Errorf("path is not dir ")
	ErrorPathIsNotRegularFile = fmt.Errorf("path is not regular file ")
	ErrorInvalidRequestPath   = fmt.Errorf("invalid http request path")
	ErrorInvalidParameters    = fmt.Errorf("parameters are invalid")
)
