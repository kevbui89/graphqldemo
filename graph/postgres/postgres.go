package postgres

import (
	"context"
	"errors"
	"fmt"

	"github.com/go-pg/pg/v10"
)

type DBLogger struct{}

func (d DBLogger) BeforeQuery(ctx context.Context, q *pg.QueryEvent) (context.Context, error) {
	return ctx, nil
}

func (d DBLogger) AfterQuery(ctx context.Context, q *pg.QueryEvent) error {
	var err error
	b, err := q.FormattedQuery()
	if err != nil {
		errors.New("could not format query")
	}

	fmt.Println(string(b))
	return nil
}

func New(opts *pg.Options) *pg.DB {
	return pg.Connect(opts)
}
