package argparser

import (
	"fmt"
	"os"
	"github.com/akamensky/argparse"

	"../../models"
)

func Parser() models.Args {
	parser := argparse.NewParser("dbash", "Quickly gain console access to a Docker container")
	_ArgContainer := parser.String("c", "container", &argparse.Options{Required: false, Help: "Container to connect"})
	_ArgLogs := parser.Flag("l", "logs", &argparse.Options{Required: false, Help: "follow container logs", Default: false})
	_ArgLogLength := parser.String("L", "log-length", &argparse.Options{Required: false, Help: "Lines to go back in the log", Default: "500"})
	err := parser.Parse(os.Args)
	if err != nil {
		fmt.Print(parser.Usage(err))
	}
	args := models.Args {
		Container: *_ArgContainer,
		LogLength: *_ArgLogLength,
		Logs: *_ArgLogs,
	}
	return args
}