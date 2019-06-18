package main

import (
	"fmt"
	"strings"
	//"log"
	"time"
	"os/exec"
	"os"

	"github.com/manifoldco/promptui"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"golang.org/x/net/context"
	"github.com/akamensky/argparse"
)

type Container struct {
	Name     	string
	CreatedAt 	time.Time
	Image  		string
	ID			string
	State		string
}

func DockerContainers() []Container {
	output := []Container{}

	cli, err := client.NewClientWithOpts(client.WithVersion("1.39"))
	if err != nil { panic(err) }

	containers, err := cli.ContainerList(context.Background(), types.ContainerListOptions{})
	if err != nil { panic(err) }

	for _, container := range containers {
		cont := Container {
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

func main() {
	fmt.Println("Dbash v2.0.0\n")

	parser := argparse.NewParser("dbash", "Quickly gain console access to a Docker container")
	_ArgContainer := parser.String("c", "container", &argparse.Options{Required: false, Help: "Container to connect"})
	err := parser.Parse(os.Args)
	if err != nil {
		fmt.Print(parser.Usage(err))
	}
	ArgContainer := *_ArgContainer
	if ArgContainer != `` {
		ContainerConsole(ArgContainer)
	} else {
		containers := DockerContainers()
		templates := &promptui.SelectTemplates{
			Label:    "{{ . }}?",
			Active:   "\U000021E8 {{ .Name | green }}",
			Inactive: "  {{ .Name | white }}",
			Selected: "\U00002714 Connecting to {{ .Name | white }}",
			Details: `
------------- Selected Container --------------
{{ "Name:" | faint }}	{{ .Name }}
{{ "Image:" | faint }}	{{ .Image }}
{{ "Created At:" | faint }}	{{ .CreatedAt }}`,
		}

		searcher := func(input string, index int) bool {
			container := containers[index]
			name := strings.Replace(strings.ToLower(container.Name), " ", "", -1)
			input = strings.Replace(strings.ToLower(input), " ", "", -1)

			return strings.Contains(name, input)
		}
		
		prompt := promptui.Select{
			Label:     "Which container to access",
			Items:     DockerContainers(),
			Templates: templates,
			Size:      10,
			Searcher:  searcher,
		}

		i, _, err := prompt.Run()

		if err != nil {
			fmt.Printf("Exited %v\n", err)
			return
		}
		ContainerConsole(containers[i].ID)
	}
}
