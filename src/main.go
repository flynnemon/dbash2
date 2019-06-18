package main

import (
	"fmt"

	"./packages/docker"
	"./packages/argparser"
	"./packages/prompts"
)

func main() {
	fmt.Println("Dbash v2.0.0\n")
	args := argparser.Parser()
	if args.Container != `` {
		docker.ContainerConsole(args.Container)
	} else {
		containers := docker.GetContainers()
		container := prompts.DockerPrompts(containers)
		docker.ContainerConsole(container.ID)
	}
}
