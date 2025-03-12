# Project Variables
PROJECT_NAME = pro
VERSION = v1.0.0
BUILD_DIR = dist

# List of supported OS and architectures
OS_LIST = linux darwin windows
ARCH_LIST = amd64 arm64

# Default target (runs tests)
.PHONY: all
all: test build

# 🛠 Build binary for the current system
.PHONY: build
build:
	@echo "🔨 Building $(PROJECT_NAME)..."
	@go build -o $(BUILD_DIR)/$(PROJECT_NAME) main.go

# 🛠 Build binaries for all platforms
.PHONY: build-all
build-all:
	@echo "🚀 Building $(PROJECT_NAME) for all platforms..."
	@mkdir -p $(BUILD_DIR)
	@for OS in $(OS_LIST); do \
		for ARCH in $(ARCH_LIST); do \
			echo "🔨 Building for $$OS/$$ARCH..."; \
			GOOS=$$OS GOARCH=$$ARCH go build -o $(BUILD_DIR)/$(PROJECT_NAME)-$$OS-$$ARCH main.go; \
		done; \
	done

# ✅ Run tests
.PHONY: test
test:
	@echo "✅ Running tests..."
	@go test ./cmd -v

# 🏗️ Clean build artifacts
.PHONY: clean
clean:
	@echo "🗑️ Cleaning up..."
	@rm -rf $(BUILD_DIR)

# 📦 Create a GitHub release
.PHONY: release
release: build-all
	@echo "🚀 Releasing $(VERSION)..."
	@git tag $(VERSION)
	@git push origin $(VERSION)
	@gh release create $(VERSION) $(BUILD_DIR)/* --title "$(PROJECT_NAME) $(VERSION)" --notes "Release $(VERSION)"

# 🏗 Run all steps before a release
.PHONY: prepare-release
prepare-release: clean test build-all