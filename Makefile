BINARY_NAME=jwtauth

build:
	@go build -o ${BINARY_NAME} src/main.go

run: build
	export GIN_MODE=release && ./${BINARY_NAME}

dev: build
	export GIN_MODE=debug && ./${BINARY_NAME}

test:
	go test -v -coverprofile=coverage.out ./...

coverage: test
	go tool cover -html=coverage.out