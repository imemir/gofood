package service

import (
	"fmt"
	"github.com/imemir/gofood/pkg/netext"
	"net"
)

var (
	serves = make(map[netext.Port]ServeFunc)
)

func Serve(port netext.Port, f ServeFunc) {
	serves[port] = f
}

type ServeFunc func(listener net.Listener) error

func (serve ServeFunc) Listen(port netext.Port) error {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%v", port))
	if err != nil {
		return err
	}
	defer func() {
		_ = listener.Close()
	}()
	return serve(listener)
}
