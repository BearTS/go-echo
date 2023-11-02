package seeds

import (
	"github.com/BearTS/go-echo/pkg/tables"
	"gorm.io/gorm"
)

func Users(db *gorm.DB) error {
	// Seed 1
	err := db.Create(&tables.Users{
		PID:        "usr_1",
		Email:      "admin@admin.com",
		Password:   []byte("admin"),
		IsVerified: true,
	}).Error

	if err != nil {
		return err
	}

	return nil
}
