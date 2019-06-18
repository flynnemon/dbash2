package main

import (
	"fmt"
	"strings"
	"os/exec"
	//"log"
	"os"

	"github.com/manifoldco/promptui"
	"github.com/ryanuber/columnize"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"golang.org/x/net/context"
	"github.com/akamensky/argparse"
)

func DockerContainers() []string {
	output := []string{}

	cli, err := client.NewClientWithOpts(client.WithVersion("1.39"))
	if err != nil { panic(err) }

	containers, err := cli.ContainerList(context.Background(), types.ContainerListOptions{})
	if err != nil { panic(err) }

	for _, container := range containers {
		containerName := strings.Replace(container.Names[0], "/", "", -1)
		containerInfo := fmt.Sprintln(containerName, "|", container.Image, "|", container.Created)
		output = append(output, containerInfo)
	}
	result := columnize.SimpleFormat(output)
	results := strings.Split(result, "\n")
	return results
}

func CommandPrep(container string, cmdPath string) error {
	cmd := exec.Command("docker", "exec", "-it", container, cmdPath)
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	return err
}

func ContainerConsole(result string) {
	container := strings.Split(result, ` `)[0]
	err := CommandPrep(container, "/bin/bash")
	if err != nil {
		CommandPrep(container, "/bin/sh")
	}
}

func main() {
	parser := argparse.NewParser("dbash", "Quickly gain console access to a Docker container")
	Container := parser.String("c", "container", &argparse.Options{Required: false, Help: "Container to connect"})
	Logs := parser.String("l", "logs", &argparse.Options{Required: false, Help: "Show logs", Default: false})
	err := parser.Parse(os.Args)
	if err != nil {
		fmt.Print(parser.Usage(err))
	}
	fmt.Println(*Logs)
	fmt.Println(*Container)
	container := *Container
	if container != `` {
		ContainerConsole(container)
	} else {
		items := DockerContainers()
		prompt := promptui.Select{
			Label: "Select a running container",
			Items: 	items,
		}

		_, result, err := prompt.Run()
		
		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			return
		}
		ContainerConsole(result)
	}
}