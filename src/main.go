package main

import (
	"fmt"

	"./packages/docker"
	"./packages/argparser"
	"./packages/prompts"
)

func main() {
	fmt.Println("Dbash v2.0.0")
	args := argparser.Parser()
	if args.Container != `` && args.Logs == true {
		docker.ContainerLogs(args.Container, args.LogLength)
	} else if args.Container != `` {
		docker.ContainerConsole(args.Container)
	} else {
		containers := docker.GetContainers()
		container := prompts.DockerPrompts(containers)
		if args.Logs == true {
			docker.ContainerLogs(container.ID, args.LogLength)
		} else {
			docker.ContainerConsole(container.ID)
		}
	}
}
