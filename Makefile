build: go.sum
	go build ./...

install: go.sum
	go install ./cmd/secretdbd
	go install ./cmd/secretdbcli

go.sum: go.mod
	go mod verify
	go mod tidy
