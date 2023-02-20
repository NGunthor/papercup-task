package youtube

import (
	"context"
	"errors"
	"fmt"
	"google.golang.org/api/youtube/v3"
	"strings"
	"time"
)

type youtubeAdapter struct {
	api *youtube.Service
}

func NewYoutubeAdapter(api *youtube.Service) *youtubeAdapter {
	return &youtubeAdapter{api: api}
}

var (
	ErrVideoNotFound = errors.New("video not found in youtube")
)

func (a *youtubeAdapter) GetVideoDuration(ctx context.Context, videoID string) (time.Duration, error) {

	response, err := a.api.Videos.List([]string{"contentDetails"}).Id(videoID).Context(ctx).Do()
	if err != nil {
		return 0, err
	}

	if len(response.Items) == 0 {
		return 0, ErrVideoNotFound
	}

	durationStr := response.Items[0].ContentDetails.Duration
	duration, err := parseDuration(durationStr)
	if err != nil {
		return 0, fmt.Errorf("error parsing duration from string (%s): %s", durationStr, videoID)
	}

	return duration, nil
}

func parseDuration(durationStr string) (time.Duration, error) {
	// Replace "PT" at the beginning of the duration string with ""
	durationStr = strings.Replace(durationStr, "PT", "", 1)

	// Replace all "H", "M", "S" with "h", "m", "s" respectively
	durationStr = strings.ReplaceAll(durationStr, "H", "h")
	durationStr = strings.ReplaceAll(durationStr, "M", "m")
	durationStr = strings.ReplaceAll(durationStr, "S", "s")

	// Parse the duration string using time.ParseDuration()
	duration, err := time.ParseDuration(durationStr)
	if err != nil {
		return duration, fmt.Errorf("error parsing duration: %s", err)
	}

	return duration, nil
}
