package app

import (
	"fmt"
	"net"
	"os"
)

type Handler func(c Context) (*Response, error)

type App struct {
	handlers map[*Route]Handler
}

func New() *App {
	return &App{
		handlers: make(map[*Route]Handler),
	}
}

func (a *App) Start(port string) error {
	l, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%s", port))
	if err != nil {
		return err
	}
	defer l.Close()

	for {
		c, err := l.Accept()
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
		go a.exec(c)
	}
}

func (a *App) exec(c net.Conn) error {
	request := make([]byte, 1024)
	_, err := c.Read(request)
	if err != nil {
		return err
	}

	_, err = c.Write([]byte(a.handle(string(request))))
	if err != nil {
		return err
	}

	return c.Close()
}

func (a *App) Register(method string, path string, handler Handler) {
	a.handlers[NewRoute(method, path)] = handler
}

func (a *App) handle(request string) string {
	r, err := parseRequest(request)
	if err != nil {
		// Recipients of an invalid request-line SHOULD respond with either a 400 (Bad Request) error or a 301 (Moved Permanently) redirect with the request-target properly encoded.
		// @see https://datatracker.ietf.org/doc/html/rfc9112#section-3.2
		return BAD_REQUEST.stringify()
	}

	handler, c, exists := a.findHandler(r)
	if !exists {
		return NOT_FOUND.stringify()
	}

	response, err := handler(*c)
	if err != nil {
		return INTERNAL_SERVER_ERROR.stringify()
	}
	return response.stringify()
}

func (a *App) findHandler(request *Request) (Handler, *Context, bool) {
	for route, handler := range a.handlers {
		if route.match(request) {
			return handler, NewContext(route, request), true
		}
	}
	return nil, nil, false
}
