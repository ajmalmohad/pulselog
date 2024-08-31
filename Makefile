BUILD_DIR = build

# Identity Microservice

IDENTITY_BINARY_NAME = identity
IDENTITY_BINARY_DIR = platform/identity

build-identity:
	cd $(IDENTITY_BINARY_DIR) && go build -o $(BUILD_DIR)/$(IDENTITY_BINARY_NAME) main.go

run-identity:
	cd $(IDENTITY_BINARY_DIR) && go run main.go

test-identity:
	cd $(IDENTITY_BINARY_DIR) && go test ./tests

clean-identity:
	cd $(IDENTITY_BINARY_DIR) && rm -rf ./$(BUILD_DIR)

dev:
	docker-compose -f docker-compose.dev.yaml --env-file .env up