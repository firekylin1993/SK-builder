package db

import (
	"github.com/google/wire"
	"gorm.io/gorm"
)

var ProviderSet = wire.NewSet(
	NewData,
	NewMysql,
)

// Data .
type Data struct {
	Mysql *gorm.DB
}

func NewData(db *gorm.DB) (*Data, error) {
	return &Data{
		Mysql: db,
	}, nil
}
