default:
	go build -ldflags="-s -w" -o ./dist/dbash ./src/main.go

install:
	cp dist/dbash /usr/local/bin

run:
	go run src/main.go