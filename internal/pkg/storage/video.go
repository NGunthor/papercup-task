package storage

import (
	"context"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/ngunthor/papercup-task/internal/pkg/domain"
)

func (s *storage) CreateVideo(ctx context.Context, video *domain.Video) error {
	result, err := s.db.ExecContext(ctx,
		`	INSERT INTO "videos"
					(id, title, description, duration, created_at)
				VALUES
    			    ($1, $2, $3, $4, $5)`,
		video.ID,
		video.Title,
		video.Description,
		video.Duration,
		timeNowUTC(),
	)
	if err != nil {
		return HandlePGError(err)
	}

	affected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if affected == 0 {
		return fmt.Errorf("internal error during creating video: incorrect rowsAffected")
	}

	return nil
}

func (s *storage) GetVideo(ctx context.Context, videoID string) (*domain.Video, error) {
	var video domain.Video
	if err := s.db.GetContext(ctx, &video,
		`SELECT * from "videos" WHERE id=$1`,
		videoID,
	); err != nil {
		return nil, HandlePGError(err)
	}

	return &video, nil
}

func (s *storage) DeleteVideo(ctx context.Context, videoID string) error {
	tx, err := s.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	if err := deleteVideo(ctx, tx, videoID); err != nil {
		return err
	}

	if err := deleteAnnotationsByVideoID(ctx, tx, videoID); err != nil {
		return err
	}

	return tx.Commit()
}

func deleteVideo(ctx context.Context, e sqlx.ExecerContext, videoID string) error {
	result, err := e.ExecContext(ctx,
		`DELETE FROM "videos" WHERE id=$1`,
		videoID,
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
