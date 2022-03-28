package main

import (
	"github.com/gorilla/mux"
	"github.com/imemir/gofood/pkg/envext"
	"github.com/imemir/gofood/pkg/service"
	"github.com/imemir/gofood/services/processing/internal/app/order"
	"log"
	"net"
	"net/http"
)

var (
	configs *service.Configs
	Name    = "processing"
)

func init() {
	configs = new(service.Configs)
	if err := envext.Load(configs); err != nil {
		log.Fatal("cannot load service configs: ", err)
	}
}

func main() {
	service.Serve(configs.HttpPort, func(listener net.Listener) error {
		router := mux.NewRouter()
		return http.Serve(listener, router)
	})
	{
		order.Consumer.Start()
	}
	service.Start(Name)
}
