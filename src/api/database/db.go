package database

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/k2glyph/blog/src/config"
)

// DB Connect
func Connect() (*gorm.DB, error) {
	db, err := gorm.Open("mysql", config.DBURL)
	if err != nil {
		return nil, err
	}
	return db, nil
}
