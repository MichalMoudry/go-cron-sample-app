install_dependencies:
	go mod tidy

test:
	go test -v ./...