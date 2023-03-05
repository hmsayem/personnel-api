.DEFAULT_GOAL := build
BIN_FILE=littleu
build:
	@go build -o "${BIN_FILE}"
clean:
	go clean
	rm --force "cp.out"
	rm --force nohup.out
test:
	go test
check:
	go test
cover:
	go test -coverprofile cp.out
	go tool cover -html=cp.out
run:
	./"${BIN_FILE}"
lint:
	golangci-lint run