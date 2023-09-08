package dao

import (
	"context"
	"github.com/Amore14rn/888Starz/internal/domain/user/model"
	"github.com/Amore14rn/888Starz/pkg/errors"
	psql "github.com/Amore14rn/888Starz/pkg/postgresql"
	"github.com/Amore14rn/888Starz/pkg/tracing"
	sq "github.com/Masterminds/squirrel"
	"strconv"
)

type UserDAO struct {
	qb     sq.StatementBuilderType
	client psql.Client
}

func NewUserStorage(client psql.Client) *UserDAO {
	return &UserDAO{
		qb:     sq.StatementBuilder.PlaceholderFormat(sq.Dollar),
		client: client,
	}
}

func (u *UserDAO) CreateUser(ctx context.Context, req *model.CreateUser) error {
	sql, args, err := u.qb.Insert("users").
		Columns("id", "first_name", "last_name", "full_name", "age", "is_married", "password").
		Values(req.ID, req.FirstName, req.LastName, req.FullName, req.Age, req.IsMarried, req.Password).ToSql()
	if err != nil {
		err = psql.ErrCreateQuery(err)
		tracing.Error(ctx, err)
	}
	tracing.SpanEvent(ctx, "Insert Product query")
	tracing.TraceVal(ctx, "sql", sql)
	for i, arg := range args {
		tracing.TraceIVal(ctx, "arg-"+strconv.Itoa(i), arg)
	}

	cmd, execErr := u.client.Exec(ctx, sql, args...)
	if execErr != nil {
		execErr = psql.ErrDoQuery(execErr)
		tracing.Error(ctx, execErr)

		return execErr
	}

	if cmd.RowsAffected() == 0 {
		return errors.New("nothing inserted")
	}

	return nil
}

func (u *UserDAO) CreateOrder(ctx context.Context, req *model.CreateOrder) error {
	sql, args, err := u.qb.Insert("orders").
		Columns("id", "user_id", "timestamp").
		Values(req.ID, req.UserID, req.Timestamp).ToSql()
	if err != nil {
		err = psql.ErrCreateQuery(err)
		tracing.Error(ctx, err)
	}
	tracing.SpanEvent(ctx, "Insert Product query")
	tracing.TraceVal(ctx, "sql", sql)
	for i, arg := range args {
		tracing.TraceIVal(ctx, "arg-"+strconv.Itoa(i), arg)
	}

	cmd, execErr := u.client.Exec(ctx, sql, args...)
	if execErr != nil {
		execErr = psql.ErrDoQuery(execErr)
		tracing.Error(ctx, execErr)

		return execErr
	}

	if cmd.RowsAffected() == 0 {
		return errors.New("nothing inserted")
	}

	return nil
}

//func (u *UserDAO) AddToOrder(ctx context.Context, req *model.AddToOrder) error {
//	sql, args, err := u.qb.Insert("order_products").
//		Columns("order_id", "product_id", "quantity", "price").
//		Values(req.ID, req.UserID, req.Products, req.Timestamp).ToSql()
//	if err != nil {
//		err = psql.ErrCreateQuery(err)
//		tracing.Error(ctx, err)
//	}
//	tracing.SpanEvent(ctx, "Insert Product query")
//	tracing.TraceVal(ctx, "sql", sql)
//	for i, arg := range args {
//		tracing.TraceIVal(ctx, "arg-"+strconv.Itoa(i), arg)
//	}
//
//	cmd, execErr := u.client.Exec(ctx, sql, args...)
//	if execErr != nil {
//		execErr = psql.ErrDoQuery(execErr)
//		tracing.Error(ctx, execErr)
//
//		return execErr
//	}
//
//	if cmd.RowsAffected() == 0 {
//		return errors.New("nothing inserted")
//	}
//
//	return nil
//}
//
//func (u *UserDAO) GetOrderByUserID(ctx context.Context, userID string) (*model.Order, error) {
//	sql, args, err := u.qb.Select("id", "user_id", "timestamp").
//		From("orders").
//		Where(sq.Eq{"user_id": userID}).
//		ToSql()
//	if err != nil {
//		err = psql.ErrCreateQuery(err)
//		tracing.Error(ctx, err)
//	}
//	tracing.SpanEvent(ctx, "Insert Product query")
//	tracing.TraceVal(ctx, "sql", sql)
//	for i, arg := range args {
//		tracing.TraceIVal(ctx, "arg-"+strconv.Itoa(i), arg)
//	}
//
//	var order model.Order
//	err = u.client.QueryRow(ctx, sql, args...).Scan(&order.ID, &order.UserID, &order.Timestamp)
//	if err != nil {
//		err = psql.ErrDoQuery(err)
//		tracing.Error(ctx, err)
//	}
//
//	return &order, nil
//}
