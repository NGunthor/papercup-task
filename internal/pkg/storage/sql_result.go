package storage

import (
	"database/sql"
	"errors"
	"github.com/jackc/pgx"
)

var (
	ErrNotFound           = errors.New("not found")
	ErrEntityAlreadyExist = errors.New("entity already exists")
)

// https://www.postgresql.org/docs/12/errcodes-appendix.html
const (
	// pgErrCodeUniqueViolation - нарушение уникальности.
	pgErrCodeUniqueViolation = "23505"
)

// HandlePGError handles error from pg
func HandlePGError(err error) error {
	if errors.Is(err, sql.ErrNoRows) {
		return ErrNotFound
	}

	var pgErr pgx.PgError
	if !errors.As(err, &pgErr) {
		return err
	}

	switch pgErr.Code {
	case pgErrCodeUniqueViolation:
		return ErrEntityAlreadyExist
	default:
		return err
	}
}
