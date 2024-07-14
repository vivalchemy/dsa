# Directories and files
TEST_DIR := $(shell find . -type f -name '*_test.go' -exec dirname {} \; | sort -u)
TMP_DIR := ./tmp
BUILD_DIR := ./build
APP_NAME := dsa
GO_FILES := $(shell find . -type f -name '*.go')

# Conditional verbosity
ifdef VERBOSE
	VERBOSE_FLAG := -v
else
	VERBOSE_FLAG :=
endif

.PHONY: all build clean dev preview test tidy

# Default target
all: preview
	@echo "Running $@"

# Build target
build: $(BUILD_DIR)/$(APP_NAME)
	@echo "Running $@"

# Clean target
clean:
	@echo "Running $@"
	@rm -rf $(TMP_DIR)

# Development target
dev: $(TMP_DIR)/$(APP_NAME)
	@echo "Running $@"
	@$<

# Preview the 1st dependency
preview: build
	@echo "Running $@"
	@$(BUILD_DIR)/$(APP_NAME)

# Test target
test:
	@echo "Running $@"
	@go $@ $(VERBOSE_FLAG) $(TEST_DIR)

# Tidy target
tidy:
	@echo "Running $@"
	@go mod $@

# create the directory and build the output binary in the mentioned file
%/$(APP_NAME): $(GO_FILES)
	@echo "Running $@"
	@mkdir -p $(@D)
	@go build -o $@ .
