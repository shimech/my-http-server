package app

import (
	"strings"
)

type Context struct {
	params map[string]string
}

func NewContext(route *Route, request *Request) *Context {
	params := make(map[string]string)
	for i, path := range route.Paths {
		if strings.HasPrefix(path, "{") && strings.HasSuffix(path, "}") {
			name := strings.TrimPrefix(path, "{")
			name = strings.TrimSuffix(name, "}")
			params[name] = request.Paths[i]
		}
	}
	return &Context{
		params: params,
	}
}

func (c *Context) Param(name string) string {
	if param, exists := c.params[name]; exists {
		return param
	} else {
		return ""
	}
}
