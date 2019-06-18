package argparser

import (
	"fmt"
	"os"
	"github.com/akamensky/argparse"
)

type Args struct {
	Container     	string
	Kubernetes		bool
	Version			bool
	Logs			bool
	LogLength		int
}

func Parser() Args {
	parser := argparse.NewParser("dbash", "Quickly gain console access to a Docker container")
	_ArgContainer := parser.String("c", "container", &argparse.Options{Required: false, Help: "Container to connect"})
	err := parser.Parse(os.Args)
	if err != nil {
		fmt.Print(parser.Usage(err))
	}
	args := Args {
		Container: *_ArgContainer,
	}
	return args
}