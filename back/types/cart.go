package types

import "gorm.io/gorm"

type Cart struct {
	gorm.Model
	Products []Product
	Owner    User
	Name     string
}
