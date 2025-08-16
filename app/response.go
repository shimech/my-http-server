package app

import (
	"fmt"
	"net/http"

	"github.com/shimech/my-http-server/util"
)

var (
	BAD_REQUEST           = NewResponse(http.StatusBadRequest, "Bad Request", "text/plain", "")
	NOT_FOUND             = NewResponse(http.StatusNotFound, "Not Found", "text/plain", "")
	INTERNAL_SERVER_ERROR = NewResponse(http.StatusInternalServerError, "Internal Server Error", "text/plain", "")
)

type Response struct {
	Status        int
	Message       string
	ContentType   string
	ContentLength int
	Body          string
}

func NewResponse(
	status int,
	message string,
	contentType string,
	body string,
) *Response {
	return &Response{
		Status:        status,
		Message:       message,
		ContentType:   contentType,
		ContentLength: len(body),
		Body:          body,
	}
}

func (r *Response) stringify() string {
	return fmt.Sprintf("HTTP/1.1 %d %s", r.Status, r.Message) + util.CRLF +
		fmt.Sprintf("Content-Type: %s", r.ContentType) + util.CRLF +
		fmt.Sprintf("Content-Length: %d", r.ContentLength) + util.CRLF +
		util.CRLF +
		r.Body
}
