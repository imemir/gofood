package service

import "github.com/imemir/gofood/pkg/netext"

type Configs struct {
	HttpPort netext.Port `env:"HTTP_PORT,required"`
}
