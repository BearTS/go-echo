package pkg

import (
	"github.com/BearTS/go-echo/pkg/tables"
	"gorm.io/gorm"
)

type Migrate struct {
	TableName string
	Run       func(*gorm.DB) error
}

func AutoMigrate(db *gorm.DB) []Migrate {
	var users tables.Users

	usersM := Migrate{TableName: "users",
		Run: func(d *gorm.DB) error { return db.AutoMigrate(&users) }}

	return []Migrate{usersM}
}
