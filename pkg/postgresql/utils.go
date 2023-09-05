package postgresql

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/jackc/pgconn"
)

func ParsePgError(err error) error {
	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) {
		pgErr = err.(*pgconn.PgError)

		return fmt.Errorf(
			"database error. message:%s, detail:%s, where:%s, sqlstate:%s",
			pgErr.Message,
			pgErr.Detail,
			pgErr.Where,
			pgErr.SQLState(),
		)
	}

	return err
}

func PrettySQL(q string) string {
	return strings.ReplaceAll(strings.ReplaceAll(q, "\t", ""), "\n", " ")
}

func DoWithAttempts(fn func() error, maxAttempts int, delay time.Duration) error {
	var err error

	for maxAttempts > 0 {
		if err = fn(); err != nil {
			time.Sleep(delay)
			maxAttempts--

			continue
		}

		return nil
	}

	return err
}
