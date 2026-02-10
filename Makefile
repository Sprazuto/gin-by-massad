# Makefile for LKE Application

# Colors for output
YELLOW = \033[33m
GREEN = \033[32m
RED = \033[31m
BLUE = \033[34m
RESET = \033[0m

# Version file
VERSION_FILE = VERSION

## RUN APPLICATION
run:
	@echo -e "$(BLUE)🚀 Running the application...$(RESET)"
	@go run *.go

## RUN TESTS
test:
	@echo -e "$(BLUE)🔍 Running tests...$(RESET)"
	@go test -v ./tests/*

## INSTALL SWAG CLI TOOL & PACKAGES
install_swag:
	@echo -e "$(YELLOW)📥 Checking Swag CLI availability...$(RESET)"
	@installed=0; \
	export PATH=$$PATH:$$(go env GOPATH)/bin; \
	if ! which swag >/dev/null 2>&1; then \
		echo -e "$(RED)❌ Swag CLI not found! Installing now...$(RESET)"; \
		go install github.com/swaggo/swag/cmd/swag@latest; \
		if which swag >/dev/null 2>&1; then \
			echo -e "$(GREEN)Swag CLI installed successfully.$(RESET)"; \
			installed=1; \
		else \
			echo -e "$(RED)Swag CLI installation failed.$(RESET)"; \
		fi; \
	else \
		echo -e "$(GREEN)Swag CLI is already installed.$(RESET)"; \
	fi; \
	if [ $$installed -eq 1 ]; then \
		echo -e "$(BLUE)🔄 Updating project dependencies for Swag...$(RESET)"; \
		go mod tidy && echo -e "$(GREEN)go mod tidy completed.$(RESET)" || echo -e "$(RED)go mod tidy failed.$(RESET)"; \
		go mod download && echo -e "$(GREEN)go mod download completed.$(RESET)" || echo -e "$(RED)go mod download failed.$(RESET)"; \
	fi; \
	echo -e "$(GREEN)✅ Swag installation check complete!$(RESET)"

## UPDATE VERSION
update_version:
	@echo -e "$(YELLOW)🔄 Updating version from $(VERSION_FILE)...$(RESET)"
	@version=`cat $(VERSION_FILE) | tr -d '\n\r'`; \
	echo -e "$(BLUE)Version: $$version$(RESET)"; \
	sed -i '' 's|// @version.*|// @version         '"$$version"'|' main.go; \
	sed -i '' 's|"patch":.*|"patch":   "'"$$version"'",|' controllers/app.go; \
	echo -e "$(GREEN)✅ Version updated in main.go and controllers/app.go$(RESET)"

## GENERATE API DOCUMENTATION
generate_docs: update_version install_swag
	@echo -e "$(YELLOW)📜 Generating API documentation using Swag...$(RESET)"
	@export PATH=$$PATH:$$(go env GOPATH)/bin && swag init
	@echo -e "$(GREEN)✅ API documentation generated successfully!$(RESET)"

## BUILD APPLICATION
build:
	@echo -e "$(BLUE)🔨 Building the application...$(RESET)"
	@go build -o lke-app .

## FULL BUILD WITH VERSION UPDATE
release: update_version generate_docs build
	@echo -e "$(GREEN)🎉 Build complete with updated version!$(RESET)"

## CLEAN BUILDS
clean:
	@echo -e "$(BLUE)🧹 Cleaning build files...$(RESET)"
	@rm -f lke-app
	@rm -rf docs

## INSTALL DEPENDENCIES
deps:
	@echo -e "$(BLUE)📦 Installing project dependencies...$(RESET)"
	@go mod download
	@go mod tidy

## SHOW HELP
help:
	@echo -e "$(GREEN)Available targets:$(RESET)"
	@echo "  run          - Run the application"
	@echo "  test         - Run tests"
	@echo "  build        - Build the application"
	@echo "  release      - Update version, generate docs, and build"
	@echo "  update_version - Update version from VERSION file"
	@echo "  generate_docs - Generate API documentation using Swag"
	@echo "  clean        - Clean build files"
	@echo "  deps         - Install dependencies"
	@echo "  install_swag - Install Swag CLI and dependencies"
	@echo "  help         - Show this help message"

.PHONY: run test install_swag update_version generate_docs build release clean deps help
