all: clean fmt mattsu-coin-client

mattsu-coin-client: *.go
	go build -o mattsu-coin-client

clean:
	go clean

test:
	go test github.com/mattsu6666/mattsu-coin-client/...

fmt:
	gofmt -s -w .
