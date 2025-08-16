package app

import "net/http"

type Request struct {
	Method string
	Path   string
}

func parseRequest(request string) (*Request, error) {
	var _ = request
	return &Request{
		Method: http.MethodGet,
		Path:   "/",
	}, nil
}
