package video_service

import (
	"context"
	"errors"

	pb "github.com/ngunthor/papercup-task/api/gen/v1"
	"github.com/ngunthor/papercup-task/internal/pkg/components/annotation"
	"github.com/ngunthor/papercup-task/internal/pkg/domain"
	"github.com/ngunthor/papercup-task/internal/pkg/storage"
	"github.com/ngunthor/papercup-task/internal/pkg/youtube"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type annotations interface {
	CreateAnnotation(ctx context.Context, annotationMeta *domain.AnnotationMeta) (*domain.Annotation, error)
	GetAnnotations(ctx context.Context, videoID string) ([]*domain.Annotation, error)
	UpdateAnnotation(ctx context.Context, annotationID int64, annotationMeta *domain.AnnotationMeta) error
	DeleteAnnotation(ctx context.Context, annotationID int64) error
}

type videos interface {
	CreateVideo(ctx context.Context, videoID, title, description string) error
	DeleteVideo(ctx context.Context, videoID string) error
}

type videoService struct {
	pb.UnimplementedVideoServiceServer

	annotations annotations
	videos      videos
}

func NewVideoService(annotations annotations, videos videos) *videoService {
	return &videoService{
		annotations: annotations,
		videos:      videos,
	}
}

func (s *videoService) CreateVideo(ctx context.Context, req *pb.CreateVideoRequest) (*pb.CreateVideoResponse, error) {
	if err := validateVideo(req.GetVideo()); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "%s", err.Error())
	}

	video := req.GetVideo()

	if err := s.videos.CreateVideo(ctx, video.GetID(), video.GetTitle(), video.GetDescription()); err != nil {
		switch {
		case errors.Is(err, youtube.ErrVideoNotFound):
			return nil, status.Errorf(codes.NotFound, "%s", err.Error())
		case errors.Is(err, storage.ErrEntityAlreadyExist):
			return nil, status.Errorf(codes.AlreadyExists, "%s", err.Error())
		default:
			return nil, status.Errorf(codes.Internal, "%s", err.Error())
		}
	}

	return &pb.CreateVideoResponse{Video: &pb.Video{
		ID:          video.GetID(),
		Title:       video.GetTitle(),
		Description: video.GetDescription(),
	}}, nil
}

func (s *videoService) DeleteVideo(ctx context.Context, req *pb.DeleteVideoRequest) (*pb.DeleteVideoResponse, error) {
	if err := s.videos.DeleteVideo(ctx, req.GetVideoID()); err != nil {
		switch {
		case errors.Is(err, storage.ErrNotFound):
			return nil, status.Errorf(codes.NotFound, "%s", err.Error())
		default:
			return nil, status.Errorf(codes.Internal, "%s", err.Error())
		}
	}

	return &pb.DeleteVideoResponse{}, nil
}

func (s *videoService) CreateAnnotation(ctx context.Context, req *pb.CreateAnnotationRequest) (*pb.CreateAnnotationResponse, error) {
	if err := validateAnnotationMeta(req.GetAnnotationMeta()); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "%s", err.Error())
	}

	createdAnnotation, err := s.annotations.CreateAnnotation(ctx, fromProtoToDomainAnnotationMeta(req.GetAnnotationMeta()))
	if err != nil {
		switch {
		case errors.Is(err, annotation.ErrInvalidAnnotationTimeRange):
			return nil, status.Errorf(codes.InvalidArgument, "%s", err.Error())
		case errors.Is(err, storage.ErrEntityAlreadyExist):
			return nil, status.Errorf(codes.AlreadyExists, "%s", err.Error())
		default:
			return nil, status.Errorf(codes.Internal, "%s", err.Error())
		}
	}

	return &pb.CreateAnnotationResponse{Annotation: fromDomainToProtoAnnotation(createdAnnotation)}, nil
}

func (s *videoService) GetAnnotations(ctx context.Context, req *pb.GetAnnotationsRequest) (*pb.GetAnnotationsResponse, error) {
	annotationList, err := s.annotations.GetAnnotations(ctx, req.GetVideoID())
	if err != nil {
		switch {
		case errors.Is(err, storage.ErrNotFound):
			return nil, status.Errorf(codes.NotFound, "%s", err.Error())
		default:
			return nil, status.Errorf(codes.Internal, "%s", err.Error())
		}
	}

	return &pb.GetAnnotationsResponse{Annotations: fromDomainToProtoAnnotationList(annotationList)}, nil
}

func (s *videoService) UpdateAnnotation(ctx context.Context, req *pb.UpdateAnnotationRequest) (*pb.UpdateAnnotationResponse, error) {
	if err := validateAnnotationMeta(req.GetAnnotation().GetAnnotationMeta()); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "%s", err.Error())
	}

	if err := s.annotations.UpdateAnnotation(ctx, req.GetAnnotation().GetID(), fromProtoToDomainAnnotationMeta(req.GetAnnotation().GetAnnotationMeta())); err != nil {
		switch {
		case errors.Is(err, storage.ErrNotFound):
			return nil, status.Errorf(codes.NotFound, "%s", err.Error())
		case errors.Is(err, annotation.ErrInvalidAnnotationTimeRange):
			return nil, status.Errorf(codes.InvalidArgument, "%s", err.Error())
		default:
			return nil, status.Errorf(codes.Internal, "%s", err.Error())
		}
	}

	return &pb.UpdateAnnotationResponse{}, nil
}

func (s *videoService) DeleteAnnotation(ctx context.Context, req *pb.DeleteAnnotationRequest) (*pb.DeleteAnnotationResponse, error) {
	if err := s.annotations.DeleteAnnotation(ctx, req.GetAnnotationID()); err != nil {
		switch {
		case errors.Is(err, storage.ErrNotFound):
			return nil, status.Errorf(codes.NotFound, "%s", err.Error())
		default:
			return nil, status.Errorf(codes.Internal, "%s", err.Error())
		}
	}

	return &pb.DeleteAnnotationResponse{}, nil
}
