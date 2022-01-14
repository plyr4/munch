up: run-build-docker

run: build run-binary

run-build-docker: linux run-docker

clean:
	@echo "### Cleaning up"
	@go mod tidy
	@go vet ./...
	@go fmt ./...

build:
	@echo
	@echo "### Building static release/munch binary"
	go build -o release/munch github.com/plyr4/munch/cmd/munch

build-static-ci:
	@echo
	@echo "### Building CI static release/munch binary"
	@go build -a \
		-ldflags '-s -w -extldflags "-static" ${LD_FLAGS}' \
		-o release/munch \
		github.com/plyr4/munch/cmd/munch

linux:
	@echo
	@echo "### Building static release/munch binary for linux"
	GOOS=linux CGO_ENABLED=0 \
		go build -o release/munch github.com/plyr4/munch/cmd/munch

docker:
	@echo "### Building Docker image"
	docker build .

run-binary:
	@echo
	@echo "### Running munch server"
	./release/munch server

run-docker:
	@echo
	@echo "### Rebuilding and running munch docker"
	docker-compose up --build -d