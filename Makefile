APP_NAME=go-api
VERSION=1.0.0

DOCKER_REGISTRY=docker.io
IMAGE_NAME=$(DOCKER_REGISTRY)/$(DOCKER_USERNAME)/$(APP_NAME)
TAG=$(VERSION)


# Build the Docker image with semver tagging
docker-build:
	docker build -t $(APP_NAME):$(TAG) .
	docker tag $(APP_NAME):$(TAG) $(IMAGE_NAME):$(TAG)

# Run the container with env variables
docker-run:
	docker run --env-file .env.local -p 8080:8080 $(APP_NAME):$(VERSION)
	
docker-push: 
	docker push $(IMAGE_NAME):$(TAG)

# Clean up (optional)
docker-clean:
	docker rmi $(APP_NAME):$(VERSION)
	
run: 
	docker-compose up --build
	
test: 
	go test ./routes
	
lint:
	golangci-lint run
	
build: 
	go build -o bin/app main.go
