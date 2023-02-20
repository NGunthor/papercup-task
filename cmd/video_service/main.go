package main

import (
	"context"
	"github.com/ngunthor/papercup-task/config"
	"github.com/ngunthor/papercup-task/internal/services/video_service"

	_ "github.com/lib/pq"
	servicepb "github.com/ngunthor/papercup-task/api/gen/v1"
	"github.com/ngunthor/papercup-task/internal/pkg/components/annotation"
	"github.com/ngunthor/papercup-task/internal/pkg/components/video"
	"github.com/ngunthor/papercup-task/internal/pkg/storage"
	youtubeAdapter "github.com/ngunthor/papercup-task/internal/pkg/youtube"
	rkboot "github.com/rookie-ninja/rk-boot"
	rkgrpc "github.com/rookie-ninja/rk-grpc/boot"
	"github.com/xlab/closer"
	"google.golang.org/grpc"
)

// Application entrance.
func main() {
	// Create a new boot instance.
	boot := rkboot.NewBoot()
	defer closer.Close()

	// ***************************************
	// ******* Register GRPC & Gateway *******
	// ***************************************

	// Get grpc entry with name
	grpcEntry := boot.GetEntry("video_service").(*rkgrpc.GrpcEntry)
	// Register grpc registration function
	grpcEntry.AddRegFuncGrpc(registerVideoService)
	// Register grpc-video_service registration function
	grpcEntry.AddRegFuncGw(servicepb.RegisterVideoServiceHandlerFromEndpoint)

	// Bootstrap
	boot.Bootstrap(context.Background())

	// Wait for shutdown sig
	boot.WaitForShutdownSig(context.Background())
}

// Implementation of [type GrpcRegFunc func(server *grpc.Server)]
func registerVideoService(server *grpc.Server) {
	db := config.MustConnectToPostgres()
	store := storage.New(db)

	annotationsComponent := annotation.NewAnnotationComponent(store)
	videoComponent := video.NewVideoComponent(store, youtubeAdapter.NewYoutubeAdapter(config.MustConnectToYoutube()))

	servicepb.RegisterVideoServiceServer(server, video_service.NewVideoService(annotationsComponent, videoComponent))
}
