package dal

import (
	"time"

	"gorm.io/gorm"
)

type GomrConfigOrder struct {
	Dialector    gorm.Dialector
	Opts         gorm.Option
	MaxOpenConns int
	MaxIdleConns int
}

var GormOrder *gorm.DB

func initGormOrder(config *GomrConfigOrder) {

	var err error

	GormOrder, err = gorm.Open(config.Dialector, config.Opts)
	if err != nil {
		panic(err)
	}

	sqlDB, err := Gorm.DB()
	if err != nil {
		panic(err)
	}

	sqlDB.SetMaxOpenConns(config.MaxOpenConns)
	sqlDB.SetMaxIdleConns(config.MaxIdleConns)
	sqlDB.SetConnMaxLifetime(time.Hour)

	err = sqlDB.Ping()
	if err != nil {
		panic(err)
	}
}
