package config

import (
	"context"

	"google.golang.org/api/option"
	"google.golang.org/api/youtube/v3"
)

func MustConnectToYoutube() *youtube.Service {
	client, err := youtube.NewService(context.Background(), option.WithAPIKey("AIzaSyCHdIXYNIZEMA2kKuldD5ghdvFUG1W9xQw"))

	if err != nil {
		panic(err)
	}

	return client
}
