package gormext

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Config struct {
	MysqlHost     string `env:"MYSQL_HOST,required"`
	MysqlPort     string `env:"MYSQL_PORT,required"`
	MysqlUsername string `env:"MYSQL_USERNAME,required"`
	MysqlPassword string `env:"MYSQL_PASSWORD,required,file"`
	MysqlDatabase string `env:"MYSQL_DATABASE,required"`
}

func Open(c Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True",
		c.MysqlUsername,
		c.MysqlPassword,
		c.MysqlHost,
		c.MysqlPort,
		c.MysqlDatabase,
	)
	return gorm.Open(mysql.New(mysql.Config{
		DSN:                       dsn,
		DefaultStringSize:         256,
		DisableDatetimePrecision:  true,
		DontSupportRenameIndex:    true,
		DontSupportRenameColumn:   true,
		SkipInitializeWithVersion: false,
	}), &gorm.Config{})
}
