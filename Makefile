BINARY=infoApp
SERVICE_NAME=order-service
DOCKER_FILE_NAME=Dockerfile
VERSION=1.0

build:
	set GOOS=linux&& set GOARCH=amd64&& go build -o ${BINARY}
	@echo Order-Service is up...


docker-build:
	docker build -f ./${DOCKER_FILE_NAME} -t ${SERVICE_NAME}:${VERSION} .
	@echo Order-Service is up...