package types

import (
	"gorm.io/gorm"
)

type Currency int

const (
	CURRENCY_PLN  Currency = iota
	CURRENCY_EURO          = iota
)

type ProductCategory int

const (
	PRODUCT_CATEGORY_ELECTRONICS = iota
	PRODUCT_CATEGORY_HOUSING     = iota
	PRODUCT_CATEGORY_HEALTH      = iota
)

type ProductPrice struct {
	gorm.Model
	Value    int
	Currency Currency
}

type Product struct {
	gorm.Model
	Name         string
	Description  string
	ImagesUrls   []string
	PreviewImage string
	Title        string
	Categories   []ProductCategory
	Price        ProductPrice
	Seller       User
	SellerID     uint
}
