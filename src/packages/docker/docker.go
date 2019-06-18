package docker

import (
	"strings"
	"time"
	"os/exec"
	"os"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"golang.org/x/net/context"

	"../../models"
)


func GetContainers() []models.Container {
	output := []models.Container{}

	cli, err := client.NewClientWithOpts(client.WithVersion("1.39"))
	if err != nil { panic(err) }

	containers, err := cli.ContainerList(context.Background(), types.ContainerListOptions{})
	if err != nil { panic(err) }

	for _, container := range containers {
		cont := models.Container {
			Name: strings.Replace(container.Names[0], "/", "", -1),
			CreatedAt: time.Unix(container.Created, 0),
			Image: container.Image,
			ID: container.ID,
			State: container.State,
		}
		output = append(output, cont)
	}

	return output
}

func CommandPrep(_container string, _cmdPath string) error {
	cmd := exec.Command("docker", "exec", "-it", _container, _cmdPath)
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	err := cmd.Run()

	return err
}

func ContainerConsole(_container string) {
	err := CommandPrep(_container, "/bin/bash")
	if err != nil {
		CommandPrep(_container, "/bin/sh")
	}
}