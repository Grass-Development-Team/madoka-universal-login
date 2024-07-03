package serializer

import "errors"

type Error struct {
	Code int
	Msg  string
	Raw  error
}

func (err *Error) Error() string {
	return err.Msg
}

func NewErrorFromResponse(res *Response) *Error {
	return &Error{
		Code: res.Code,
		Msg:  res.Msg,
		Raw:  errors.New(res.Error),
	}
}

// The three-digit code uses the original meaning of the http response code
// The four-digit code was custom code of the program
const (
	// Success
	CodeOK = 200
	// Bad Request
	CodeBadRequest = 400
	// Unauthorized
	CodeUnauthorized = 401
	// Not Found
	CodeNotFound = 404
)
