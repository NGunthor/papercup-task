package video

import (
	"context"
	"github.com/ngunthor/papercup-task/internal/pkg/domain"
	"time"
)

type videoComponent struct {
	storage storage
	youtube youtubeAdapter
}

type storage interface {
	CreateVideo(ctx context.Context, video *domain.Video) error
	DeleteVideo(ctx context.Context, videoID string) error
}

type youtubeAdapter interface {
	GetVideoDuration(ctx context.Context, videoID string) (time.Duration, error)
}

func NewVideoComponent(storage storage, youtubeAdapter youtubeAdapter) *videoComponent {
	return &videoComponent{
		storage: storage,
		youtube: youtubeAdapter,
	}
}

func (c *videoComponent) CreateVideo(ctx context.Context, videoID, title, description string) error {
	videoDuration, err := c.youtube.GetVideoDuration(ctx, videoID)
	if err != nil {
		return err
	}

	return c.storage.CreateVideo(ctx, &domain.Video{
		ID:          videoID,
		Title:       title,
		Description: description,
		Duration:    videoDuration,
	})
}

func (c *videoComponent) DeleteVideo(ctx context.Context, videoID string) error {
	return c.storage.DeleteVideo(ctx, videoID)
}
