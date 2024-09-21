package store

import (
	"context"
	"github.com/jackc/pgx/v5"
)

type SQLX struct {
	Client *pgx.Conn
}

func NewSQL(dataSourceName string, ctx context.Context) (store *pgx.Conn, err error) {
	store, err = pgx.Connect(ctx, dataSourceName)
	if err != nil {
		return nil, err
	}

	defer func(conn *pgx.Conn, ctx context.Context) {
		err := conn.Close(ctx)
		if err != nil {
			return
		}
	}(store, ctx)

	return store, nil
}

//func (s *SQLX) Close(ctx context.Context) error {
//	if s.Client != nil {
//		return s.Client.Close(ctx)
//	}
//	return nil
//}
