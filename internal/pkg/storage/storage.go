package storage

import (
	"github.com/jmoiron/sqlx"
	"time"
)

type storage struct {
	db *sqlx.DB
}

var timeNowUTC = func() time.Time {
	return time.Now().UTC()
}

func New(
	db *sqlx.DB,
) *storage {
	return &storage{
		db: db,
	}
}
