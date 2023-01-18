package db

import "github.com/notnulldev/e-com-2/types"

type ProductDb interface {
	GetProductByName(productName string) (types.Product, error)
	GetProductsWithCategories(categories []types.ProductCategory) ([]types.Product, error)
	GetProductWithName(name string) (types.Product, error)
	GetProductsContainsName(contains string) ([]types.Product, error)
}

// TODO: create SQLFilter per query
