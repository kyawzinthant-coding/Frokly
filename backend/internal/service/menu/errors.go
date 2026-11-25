package menu

import (
	"errors"

	"github.com/jackc/pgx/v5/pgconn"
)

const pgUniqueViolation = "23505"

var (
	ErrInvalidInput     = errors.New("invalid input")
	ErrNameTaken        = errors.New("name already taken")
	ErrCategoryNotFound = errors.New("category not found")
	ErrCategoryNotEmpty = errors.New("category has active products")
	ErrProductNotFound  = errors.New("product not found")
)

func isUniqueViolation(err error) bool {
	var pgErr *pgconn.PgError
	return errors.As(err, &pgErr) && pgErr.Code == pgUniqueViolation
}
