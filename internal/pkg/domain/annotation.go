package domain

import "time"

type AnnotationMeta struct {
	VideoID string        `db:"video_id"`
	Start   time.Duration `db:"start_at"`
	End     time.Duration `db:"end_at"`
	Type    string        `db:"annotation_type"`
	Notes   string        `db:"notes"`
}

type Annotation struct {
	ID              int64 `db:"id"`
	*AnnotationMeta `db:",inline"`
	CreatedAt       time.Time `db:"created_at"`
}
