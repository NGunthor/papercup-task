package video_service

import (
	"errors"
	pb "github.com/ngunthor/papercup-task/api/gen/v1"
	"regexp"
)

var (
	ErrNilRequest              = errors.New("request can't be nil")
	ErrEmptyTitle              = errors.New("title can't be nil")
	ErrIncorrectDurationFormat = errors.New("incorrect duration format; should be in format \"mm:ss\"; for example 03:45")
)

func validateVideo(video *pb.Video) error {
	if video == nil {
		return ErrNilRequest
	}

	if video.Title == "" {
		return ErrEmptyTitle
	}

	//... whatever checks we want

	return nil
}

func validateAnnotationMeta(meta *pb.AnnotationMeta) error {
	if meta == nil {
		return ErrNilRequest
	}

	if match, _ := regexp.MatchString("(^\\d{2}|^\\d{2}:\\d{2}):\\d{2}$", meta.GetStart()); !match {
		return ErrIncorrectDurationFormat
	}

	if match, _ := regexp.MatchString("(^\\d{2}|^\\d{2}:\\d{2}):\\d{2}$", meta.GetEnd()); !match {
		return ErrIncorrectDurationFormat
	}

	//... whatever checks we want

	return nil
}
