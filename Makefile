all: mattsu-coin-client

mattsu-coin-client: *.go
	go build -o mattsu-coin-client
