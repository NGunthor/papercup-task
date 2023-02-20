package video_service

import (
	pb "github.com/ngunthor/papercup-task/api/gen/v1"
	"github.com/ngunthor/papercup-task/internal/pkg/domain"
	"strconv"
	"strings"
	"time"
)

func fromProtoToDomainAnnotationMeta(meta *pb.AnnotationMeta) *domain.AnnotationMeta {
	return &domain.AnnotationMeta{
		VideoID: meta.GetVideoID(),
		Start:   getDurationFromRequest(meta.GetStart()),
		End:     getDurationFromRequest(meta.GetEnd()),
		Type:    meta.GetType(),
		Notes:   meta.GetNotes(),
	}
}

func fromDomainToProtoAnnotationList(annotations []*domain.Annotation) []*pb.Annotation {
	resultAnnotations := make([]*pb.Annotation, 0, len(annotations))

	for _, annotation := range annotations {
		resultAnnotations = append(resultAnnotations, fromDomainToProtoAnnotation(annotation))
	}

	return resultAnnotations
}

func fromDomainToProtoAnnotation(annotation *domain.Annotation) *pb.Annotation {
	return &pb.Annotation{
		ID: annotation.ID,
		AnnotationMeta: &pb.AnnotationMeta{
			VideoID: annotation.VideoID,
			Start:   getDurationString(annotation.Start),
			End:     getDurationString(annotation.End),
			Type:    annotation.Type,
			Notes:   annotation.Notes,
		},
	}
}

func getDurationFromRequest(duration string) time.Duration {
	split := strings.Split(duration, ":")

	var (
		hours   int
		minutes int
		seconds int
	)

	switch len(split) {
	case 2:
		minutes, _ = strconv.Atoi(split[0])
		seconds, _ = strconv.Atoi(split[1])
	case 3:
		hours, _ = strconv.Atoi(split[0])
		minutes, _ = strconv.Atoi(split[1])
		seconds, _ = strconv.Atoi(split[2])
	}

	return time.Duration(hours)*time.Hour + time.Duration(minutes)*time.Minute + time.Duration(seconds)*time.Second
}

func getDurationString(duration time.Duration) string {
	var ()

	return duration.String()
}
