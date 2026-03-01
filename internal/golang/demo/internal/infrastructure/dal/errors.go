package dal

import (
	"errors"

	"github.com/jackc/pgx/v5"
)

var (
	ErrFailedCollectExactlyOneRow = errors.New("failed query")
	ErrFailedQuery                = errors.New("failed query")
	ErrFailedExec                 = errors.New("failed exec")
)

func IsNotFound(err error) bool {
	if err == nil {
		return false
	}

	return errors.Is(err, pgx.ErrNoRows)
}
