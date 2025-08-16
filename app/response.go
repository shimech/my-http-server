package app

import (
	"fmt"
	"net/http"

	"github.com/shimech/my-http-server/util"
)

var (
	BAD_REQUEST           = NewResponse(http.StatusBadRequest, "Bad Request")
	NOT_FOUND             = NewResponse(http.StatusNotFound, "Not Found")
	INTERNAL_SERVER_ERROR = NewResponse(http.StatusInternalServerError, "Internal Server Error")
)

type Response struct {
	Status  int
	Message string
}

func NewResponse(status int, message string) *Response {
	return &Response{
		Status:  status,
		Message: message,
	}
}

func (r *Response) stringify() string {
	return fmt.Sprintf("HTTP/1.1 %d %s%s%s", r.Status, r.Message, util.CRLF, util.CRLF)
}
