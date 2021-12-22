package ginhandler

// HttpErrorCode http errorcode list
var HttpErrorCode = struct {
	InvalidQueryParameter string
	InternalFailure       string
}{
	InvalidQueryParameter: "InvalidQueryParameter",
	InternalFailure:       "InternalFailure",
}
