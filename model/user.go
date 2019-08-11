package model

import (
	"github.com/jinzhu/gorm"
)

// User contains credentials to access the webapp
type User struct {
	gorm.Model
	FirstName string
	LastName  string
	Email     string
}
