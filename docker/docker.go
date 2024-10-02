package docker

import (
	"archive/tar"
	"context"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

func TarDirectory(sourceDir string) (io.ReadCloser, error) {
	pipeReader, pipeWriter := io.Pipe()

	go func() {
		tarWriter := tar.NewWriter(pipeWriter)
		defer pipeWriter.Close()
		defer tarWriter.Close()

		err := filepath.Walk(sourceDir, func(file string, fi os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			header, err := tar.FileInfoHeader(fi, fi.Name())
			if err != nil {
				return err
			}

			relPath, err := filepath.Rel(sourceDir, file)
			if err != nil {
				return err
			}
			header.Name = filepath.ToSlash(relPath)

			if err := tarWriter.WriteHeader(header); err != nil {
				return err
			}

			if fi.IsDir() {
				return nil
			}

			f, err := os.Open(file)
			if err != nil {
				return err
			}
			defer f.Close()

			_, err = io.Copy(tarWriter, f)
			return err
		})

		if err != nil {
			fmt.Println("Error tarring directory:", err)
			pipeWriter.CloseWithError(err)
			return
		}
	}()

	return pipeReader, nil
}

func BuildDockerImage(dockerfilePath, imageName string) error {
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return fmt.Errorf("failed to create Docker client: %v", err)
	}

	tarStream, err := TarDirectory(dockerfilePath)
	if err != nil {
		return fmt.Errorf("failed to tar Docker build context: %v", err)
	}
	defer tarStream.Close()

	buildOptions := types.ImageBuildOptions{
		Tags:       []string{imageName},
		Dockerfile: "Dockerfile",
		Remove:     true,
	}

	buildResponse, err := cli.ImageBuild(context.Background(), tarStream, buildOptions)
	if err != nil {
		return fmt.Errorf("failed to build Docker image: %v", err)
	}
	defer buildResponse.Body.Close()

	_, err = io.Copy(os.Stdout, buildResponse.Body)
	if err != nil {
		return fmt.Errorf("failed to stream Docker build output: %v", err)
	}

	return nil
}
