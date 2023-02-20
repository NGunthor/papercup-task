# checkout-task

This is a simple application on grpc with provided swagger

### Deployment

Deployment makes docker container with the App(see: ./Dockerfile) and compose it with postgres container(see: ./docker-compose.yml)

Before deployment(see: ./Makefile `all`): runs tests, makes migrations, generates API, runs the App

### Project structure

```shell
├───api
│   └───v1
├───cmd
│   └───video_service
├───config
├───internal
│   ├───auth
│   ├───pkg
│   │   ├───components
│   │   │   ├───annotation
│   │   │   └───video
│   │   ├───domain
│   │   ├───storage
│   │   └───youtube
│   └───services
│       └───video_service
└───migration



```

### Run Application

```shell
git clone https://github.com/NGunthor/papercup-task
cd papercup-task
docker-compose up -d
```

##### Rebuild containers

`docker-compose build --no-cache`

### Add new endpoint
1) modify api/v1/ or create new .proto
```protobuf
service VideoService {
  // ...
  rpc GetVideo(GetVideoRequest) returns (GetVideoResponse) {};
}
```
2) `make generate`
3) add API handler ./internal/services
```go
func (s *videoService) GetVideo(ctx context.Context, in *pb.GetVideoRequest) (*pb.GetVideoResponse, error) {
	// ...
}
```
4) add source files ./internal/pkg

### API

address: `localhost:8080/sw`

API Accordingly to the Task

- CreateVideo
- DeleteVideo
- CreateAnnotation
- GetAnnotations
- UpdateAnnotation
- DeleteAnnotation