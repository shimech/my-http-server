package main

import (
	"fmt"
	"net"
	"net/http"
	"os"

	"github.com/shimech/my-http-server/app"
)

// Ensures gofmt doesn't remove the "net" and "os" imports above (feel free to remove this!)
var _ = net.Listen
var _ = os.Exit

func main() {
	a := app.New()
	a.Register(http.MethodGet, "/", func(c app.Context) (*app.Response, error) {
		return app.NewResponse(http.StatusOK, "OK", "text/plain", "Hello, World!"), nil
	})
	a.Register(http.MethodGet, "/echo/{message}", func(c app.Context) (*app.Response, error) {
		message := c.Param("message")
		return app.NewResponse(http.StatusOK, "OK", "text/plain", message), nil
	})
	err := a.Start("4221")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
