package ProviderMysql

import (
	"esim/config"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewMysql(cfg *config.Configuration) (*gorm.DB, error) {
	//dsn := "host=" + cfg.HostMysql + " user=" + cfg.UserMysql + " password=" + cfg.PasswordMysql + " dbname=" + cfg.DBNameMysql + " port=" + cfg.PortMysql + " sslmode=disable TimeZone=UTC"
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", cfg.UserMysql, cfg.PasswordMysql,
		cfg.HostMysql, cfg.PortMysql, cfg.DBNameMysql)
	fmt.Println(dsn)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	//db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}
