package order

import (
	"github.com/imemir/gofood/pkg/envext"
	log "github.com/sirupsen/logrus"
)

type Configs struct {
	RedisAddress  string `env:"REDIS_ADDRESS,required"`
	RedisPassword string `env:"REDIS_PASSWORD,required"`
	RedisDB       int    `env:"REDIS_DB,required"`
}

var configs Configs

func init() {
	if err := envext.Load(&configs); err != nil {
		log.WithError(err).Fatal("can not load consumer configs")
	}
}
