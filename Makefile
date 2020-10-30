run:
	reflex -r '\.go' -s -- sh -c "go run main.go"

test: 
	go test

build:
	go build