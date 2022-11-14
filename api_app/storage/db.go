package storage

import (
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/pkg/errors"
	"golang.org/x/net/context"
	"log"
	"time"
)

func ConnectToPostgres(ctx context.Context, url string) (*pgxpool.Pool, error) {
	var err error
	for i := 0; i < 5; i++ {
		p, err := pgxpool.Connect(ctx, url)
		if err != nil || p == nil {
			time.Sleep(3 * time.Second)
			continue
		}
		log.Println("Success connect to DB")
		return p, nil
	}
	return nil, errors.Wrap(err, "timed out waiting to connect postgres")
}
