package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"ksp.sk/proboj/web/config"
)

var Db *gorm.DB

func OpenDatabase() error {
	var err error
	Db, err = gorm.Open(sqlite.Open(config.Configuration.Database), &gorm.Config{})
	if err != nil {
		return err
	}

	err = Db.AutoMigrate(&Player{}, &PlayerVersion{}, &Map{}, &Game{})
	if err != nil {
		return err
	}

	return nil
}
