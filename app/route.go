package app

import (
	"strings"
)

type Route struct {
	Method string
	Paths  []string
}

func NewRoute(method string, path string) *Route {
	return &Route{
		Method: method,
		Paths:  parsePath(path),
	}
}

func parsePath(path string) []string {
	p := strings.TrimPrefix(path, "/")
	return strings.Split(p, "/")
}

func (r *Route) match(request *Request) bool {
	if r.Method != request.Method {
		return false
	}
	if len(r.Paths) != len(request.Paths) {
		return false
	}
	for i, path := range r.Paths {
		if strings.HasPrefix(path, "{") && strings.HasSuffix(path, "}") {
			continue
		}
		if path != request.Paths[i] {
			return false
		}
	}
	return true
}
