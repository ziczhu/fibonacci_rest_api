GOPACKAGES = $(shell go list ./...  | grep -v /vendor/)
DOCKER_TEST_TAG := ziczhu/fibonacci-test

default: start

build:
	@docker-compose build

start:
	@docker-compose up -d

logs:
	@docker-compose logs -f fibonacci

stop:
	@docker-compose down

docker-test:
	@docker build -t $(DOCKER_TEST_TAG) -f Dockerfile.test .
	@docker run $(DOCKER_TEST_TAG)

test:
	@go test -v -cover -coverprofile=coverage.out $(GOPACKAGES)
