package app

import (
	"fmt"
	"net"
)

type App struct {
	handlers map[Request]func() (*Response, error)
}

func New() *App {
	return &App{
		handlers: make(map[Request]func() (*Response, error)),
	}
}

func (a *App) Start(port string) error {
	l, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%s", port))
	if err != nil {
		return err
	}
	defer l.Close()

	c, err := l.Accept()
	if err != nil {
		return err
	}
	defer c.Close()

	request := make([]byte, 1024)
	_, err = c.Read(request)
	if err != nil {
		return err
	}

	_, err = c.Write([]byte(a.handle(string(request))))
	if err != nil {
		return err
	}

	return nil
}

func (a *App) Register(method string, path string, handler func() (*Response, error)) {
	a.handlers[Request{Method: method, Path: path}] = handler
}

func (a *App) handle(request string) string {
	r, err := parseRequest(request)
	if err != nil {
		return BAD_REQUEST.stringify()
	}
	if handler, exists := a.handlers[*r]; exists {
		response, err := handler()
		if err != nil {
			return INTERNAL_SERVER_ERROR.stringify()
		}
		return response.stringify()
	} else {
		return NOT_FOUND.stringify()
	}
}
