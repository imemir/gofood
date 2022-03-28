package envext

import (
	"github.com/caarlos0/env/v6"
	"github.com/joho/godotenv"
)

func Load(v interface{}) error {
	err := godotenv.Load("configs/.env")
	if err != nil {
		return err
	}

	if err = env.Parse(v); err != nil {
		return err
	}

	return nil
}
