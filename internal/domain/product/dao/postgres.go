package dao

import (
	psql "github.com/Amore14rn/888Starz/pkg/postgresql"
	sq "github.com/Masterminds/squirrel"
)

type ProductDAO struct {
	qb     sq.StatementBuilderType
	client psql.Client
}

func NewProductStorage(client psql.Client) *ProductDAO {
	return &ProductDAO{
		qb:     sq.StatementBuilder.PlaceholderFormat(sq.Dollar),
		client: client,
	}
}
