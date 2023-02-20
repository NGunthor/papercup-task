package domain

import "time"

type Video struct {
	ID          string        `db:"id"`
	Title       string        `db:"title"`
	Description string        `db:"description"`
	Duration    time.Duration `db:"duration"`
	CreatedAt   time.Time     `db:"created_at"`
}
