package service

import (
	"github.com/gorilla/mux"
	"github.com/imemir/gofood/pkg/dispose"
	"github.com/imemir/gofood/pkg/netext"
	log "github.com/sirupsen/logrus"
)

type Options struct {
	Name     string
	Host     string
	HttpPort string
	Router   *mux.Router
}

func Start(name string) {
	var exposes []netext.Port
	interrupt := make(chan error)
	for port, serve := range serves {
		exposes = append(exposes, port)
		go func(port netext.Port, serve ServeFunc) {
			interrupt <- serve.Listen(port)
		}(port, serve)
	}

	data := log.Fields{
		"service": name,
		"exposes": exposes,
	}
	log.WithFields(data).Info("service is started successfully!")

	interruptErr := <-interrupt
	if err := dispose.Close(); err != nil {
		log.WithError(err).Error("can not dispose")
	}
	if interruptErr == nil {
		log.WithFields(data).Panic("service interrupted")
	} else {
		log.WithFields(data).WithError(interruptErr).Fatal("service interrupted")
	}
}
