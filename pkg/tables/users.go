package tables

import (
	"time"

	"gorm.io/gorm"
)

// A struct on which the methods are defined
type DB struct {
	gormDB *gorm.DB
}

func NewDB(gormDB *gorm.DB) *DB {
	return &DB{gormDB: gormDB}
}

type Users struct {
	ID  int    `gorm:"column:id;primaryKey;autoIncrement"`
	PID string `gorm:"column:pid;unique;not null;type:varchar(40)"`
	// Name               string `gorm:"column:name;not null;type:varchar(100)"`
	Email string `gorm:"column:email;unique;type:varchar(100)"`
	// Phone              string `gorm:"column:phone;not null;type:varchar(10)"`
	// RegistrationNumber string `gorm:"column:registration_number;type:varchar(9)"` // optional
	Password []byte `gorm:"column:password;not null"`
	// IsExternal         bool   `gorm:"column:is_external;not null;default:false"`
	IsVerified bool `gorm:"column:is_verified;not null;default:false"`
	// IsSandbox          bool   `gorm:"column:is_sandbox;not null;default:false"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (t *Users) TableName() string {
	return "users"
}

// Create a new user
func (db *DB) CreateUser(user *Users) error {
	user.PID = UUIDWithPrefix("usr")
	return db.gormDB.Create(user).Error
}

// Get a user by id
func (db *DB) GetUserByID(id int) (*Users, error) {
	var user Users
	err := db.gormDB.Where("id = ?", id).First(&user).Error
	return &user, err
}

// Get a user by pid
func (db *DB) GetUserByPID(pid string) (*Users, error) {
	var user Users
	err := db.gormDB.Where("pid = ?", pid).First(&user).Error
	return &user, err
}
