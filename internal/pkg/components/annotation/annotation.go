package annotation

import (
	"context"
	"errors"
	"github.com/ngunthor/papercup-task/internal/pkg/domain"
	"time"
)

type annotationComponent struct {
	storage storage
}

type storage interface {
	GetVideo(ctx context.Context, videoID string) (*domain.Video, error)
	CreateAnnotation(ctx context.Context, meta *domain.AnnotationMeta) (*domain.Annotation, error)
	GetAnnotationsByVideoID(ctx context.Context, videoID string) ([]*domain.Annotation, error)
	UpdateAnnotation(ctx context.Context, annotationID int64, meta *domain.AnnotationMeta) error
	DeleteAnnotation(ctx context.Context, annotationID int64) error
}

func NewAnnotationComponent(storage storage) *annotationComponent {
	return &annotationComponent{storage: storage}
}

var (
	ErrInvalidAnnotationTimeRange = errors.New("annotation's startTime and endTime should be in video duration range")
)

func (c *annotationComponent) CreateAnnotation(ctx context.Context, annotationMeta *domain.AnnotationMeta) (*domain.Annotation, error) {
	video, err := c.storage.GetVideo(ctx, annotationMeta.VideoID)
	if err != nil {
		return nil, err
	}

	if !isValidAnnotationTimeRange(video.Duration, annotationMeta.Start, annotationMeta.End) {
		return nil, ErrInvalidAnnotationTimeRange
	}

	return c.storage.CreateAnnotation(ctx, annotationMeta)
}

func (c *annotationComponent) GetAnnotations(ctx context.Context, videoID string) ([]*domain.Annotation, error) {
	return c.storage.GetAnnotationsByVideoID(ctx, videoID)
}

func (c *annotationComponent) UpdateAnnotation(ctx context.Context, annotationID int64, annotationMeta *domain.AnnotationMeta) error {
	video, err := c.storage.GetVideo(ctx, annotationMeta.VideoID)
	if err != nil {
		return err
	}

	if !isValidAnnotationTimeRange(video.Duration, annotationMeta.Start, annotationMeta.End) {
		return ErrInvalidAnnotationTimeRange
	}

	return c.storage.UpdateAnnotation(ctx, annotationID, annotationMeta)
}

func (c *annotationComponent) DeleteAnnotation(ctx context.Context, annotationID int64) error {
	return c.storage.DeleteAnnotation(ctx, annotationID)
}

func isValidAnnotationTimeRange(videoDuration, start, end time.Duration) bool {
	if start < 0 || end < 0 {
		return false
	}

	if start > end {
		return false
	}

	if start > videoDuration || end > videoDuration {
		return false
	}

	return true
}
