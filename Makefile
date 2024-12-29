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

## help: print this help message (DEFAULT)
.PHONY: help
help:
	@echo -e "Make commands for ${binary_name}\n"
	@echo 'Usage:'
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' |  sed -e 's/^/ /' | sort

.PHONY: confirm
confirm:
	@echo -n 'Are you sure? [y/N] ' && read ans && [ $${ans:-N} = y ]

.PHONY: no-dirty
no-dirty:
	@test -z "$(shell git status --porcelain)"

all: preview
	@echo "Running $@"

## build: Build target
build: $(BUILD_DIR)/$(APP_NAME)
	@echo "Running $@"

## clean: Clean target
clean:
	@echo "Running $@"
	@rm -rf $(TMP_DIR)

## dev: Development target
dev: $(TMP_DIR)/$(APP_NAME)
	@echo "Running $@"
	@$<

## preview: Preview the 1st dependency
preview: build
	@echo "Running $@"
	@$(BUILD_DIR)/$(APP_NAME)

## test: Test target
test:
	@echo "Running $@"
	@go $@ $(VERBOSE_FLAG) $(TEST_DIR)

## tidy: Tidy target
tidy:
	@echo "Running $@"
	@go mod $@

# create the directory and build the output binary in the mentioned file
%/$(APP_NAME): $(GO_FILES)
	@echo "Running $@"
	@mkdir -p $(@D)
	@go build -o $@ .
