BUILD_DIR = build

# Auth Microservice

AUTH_BINARY_NAME = auth
AUTH_BINARY_DIR = platform/auth

build-auth:
	cd $(AUTH_BINARY_DIR) && go build -o $(BUILD_DIR)/$(AUTH_BINARY_NAME) main.go

run-auth:
	cd $(AUTH_BINARY_DIR) && go run main.go

test-auth:
	cd $(AUTH_BINARY_DIR) && go test ./tests

clean-auth:
	cd $(AUTH_BINARY_DIR) && rm -rf ./$(BUILD_DIR)