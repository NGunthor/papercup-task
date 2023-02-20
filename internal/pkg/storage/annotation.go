package storage

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/ngunthor/papercup-task/internal/pkg/domain"
)

func (s *storage) CreateAnnotation(ctx context.Context, meta *domain.AnnotationMeta) (*domain.Annotation, error) {
	annotation := &domain.Annotation{
		AnnotationMeta: meta,
		CreatedAt:      timeNowUTC(),
	}

	if err := s.db.GetContext(ctx, &annotation.ID,
		`	INSERT INTO "annotations"
					(video_id, start_at, end_at, annotation_type, notes, created_at)
				VALUES
    			    ($1, $2, $3, $4, $5, $6)
    			    RETURNING id`,
		meta.VideoID,
		meta.Start,
		meta.End,
		meta.Type,
		meta.Notes,
		annotation.CreatedAt,
	); err != nil {
		return nil, HandlePGError(err)
	}

	return annotation, nil
}

func (s *storage) GetAnnotationsByVideoID(ctx context.Context, videoID string) ([]*domain.Annotation, error) {
	var annotations []*domain.Annotation
	if err := s.db.SelectContext(ctx, &annotations,
		`	SELECT * FROM "annotations"
				WHERE video_id=$1`,
		videoID,
	); err != nil {
		return nil, HandlePGError(err)
	}

	if len(annotations) == 0 {
		return nil, ErrNotFound
	}

	return annotations, nil
}

func (s *storage) UpdateAnnotation(ctx context.Context, annotationID int64, meta *domain.AnnotationMeta) error {
	result, err := s.db.ExecContext(ctx,
		`	UPDATE "annotations"
				SET video_id=$2,
				    start_at=$3,
				    end_at=$4,
				    annotation_type=$5,
				    notes=$6
				WHERE id=$1`,
		annotationID,
		meta.VideoID,
		meta.Start,
		meta.End,
		meta.Type,
		meta.Notes,
	)
	if err != nil {
		return HandlePGError(err)
	}

	affected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if affected == 0 {
		return ErrNotFound
	}

	return nil
}

func (s *storage) DeleteAnnotation(ctx context.Context, annotationID int64) error {
	result, err := s.db.ExecContext(ctx,
		`DELETE FROM "annotations" WHERE id=$1`,
		annotationID,
	)
	if err != nil {
		return HandlePGError(err)
	}

	affected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if affected == 0 {
		return ErrNotFound
	}

	return nil
}

func deleteAnnotationsByVideoID(ctx context.Context, e sqlx.ExecerContext, videoID string) error {
	_, err := e.ExecContext(ctx,
		`DELETE FROM "annotations" WHERE video_id=$1`,
		videoID,
	)
	if err != nil {
		return err
	}

	return nil
}
