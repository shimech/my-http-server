package app

import (
	"fmt"
	"strings"

	"github.com/shimech/my-http-server/util"
)

type Request struct {
	Method string
	Path   string
}

func parseRequest(request string) (*Request, error) {
	lines := strings.Split(request, util.CRLF)
	if len(lines) < 1 {
		return nil, fmt.Errorf("invalid request")
	}
	form := lines[0]
	parts := strings.Split(form, " ")
	if len(parts) < 2 {
		return nil, fmt.Errorf("invalid request")
	}
	return &Request{
		Method: parts[0],
		Path:   parts[1],
	}, nil
}
