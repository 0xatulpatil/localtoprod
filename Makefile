APP_NAME=go-api
VERSION=1.0.0

# Build the Docker image with semver tagging
docker-build:
	docker build -t $(APP_NAME):$(VERSION) .

# Run the container with env variables
docker-run:
	docker run --env-file .env.local -p 8080:8080 $(APP_NAME):$(VERSION)

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
