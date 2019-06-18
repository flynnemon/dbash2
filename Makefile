build:
	go build -ldflags="-s -w" -o ./dist/dbash ./src/main.go
run:
	go run src/main.go
configure:
	go get github.com/manifoldco/promptui
	go get github.com/ryanuber/columnize
	go get github.com/docker/docker/api/types
	go get github.com/docker/docker/client
	go get golang.org/x/net/context
	go get github.com/akamensky/argparse
