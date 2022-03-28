package repository

import (
	"github.com/imemir/gofood/pkg/envext"
	"github.com/imemir/gofood/pkg/gormext"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

var (
	Orders *orderRepository
)

func init() {
	var (
		configs gormext.Config
		db      *gorm.DB
		err     error
	)
	configs = gormext.Config{}
	if err = envext.Load(&configs); err != nil {
		log.Fatal("cannot load repository env")
	}

	db, err = gormext.Open(configs)
	if err != nil {
		log.Fatal(err)
	}

	if err = db.Transaction(func(tx *gorm.DB) error {
		if err = tx.AutoMigrate(
			new(Order),
		); err != nil {
			return err
		}
		return nil
	}); err != nil {
		log.WithError(err).Fatal("can not migration database")
	}

	Orders = &orderRepository{
		db: db,
	}
}
