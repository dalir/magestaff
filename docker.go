package magestaff

import (
	"context"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

const defaultDockerAPIVersion = "v1.40"

func BuildDocker(dockerFilePath string) (err error) {
	cli, err := client.NewEnvClient()
	if err != nil {
		return
	}
	opt := types.ImageBuildOptions{
		Dockerfile: dockerFilePath,
	}
	_, err = cli.ImageBuild(context.Background(), nil, opt)
	return
}
