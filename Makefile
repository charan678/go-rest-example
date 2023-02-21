BINARY_NAME=go-rest-examples

build-mac:
	GOARCH=amd64 GOOS=darwin go build -o ${BINARY_NAME}-darwin main.go

build:
	GOARCH=amd64 GOOS=linux go build -o ${BINARY_NAME}-linux main.go

run: build
	./${BINARY_NAME}-linux

clean:
	go clean
	rm ${BINARY_NAME}-darwin
	rm ${BINARY_NAME}-linux
	rm ${BINARY_NAME}-windows

test:
	go test ./...

test_coverage:
	go test ./... -coverprofile=coverage.out

dep:
	go mod download

vet:
	go vet

lint:
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	golangci-lint run --enable-all

docker-run:
	docker build . -t go-rest-example
	docker run -d -p 8080:8080 go-rest-example

ci: vet build test