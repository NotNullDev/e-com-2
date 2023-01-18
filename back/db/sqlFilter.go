package db

type SQLFilter interface {
	GetWhereClause() string
}
