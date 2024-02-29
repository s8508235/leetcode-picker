##@ Help

.PHONY: help run build run-medium run-normal run-hard

run: ## run picker with negative credit and difficulty easy,medium,normal,hard
	go run cmd/leetcode-picker/main.go

run-medium: ## run picker medium problems with rating over 80%
	go run cmd/leetcode-picker/main.go -l m -r 4

run-normal: ## run picker medium, hard problems with rating over 80%
	go run cmd/leetcode-picker/main.go -l n -r 4

run-hard: ## run picker hard problems with rating over 80%
	go run cmd/leetcode-picker/main.go -l h -r 4

help:  ## Display this help
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[0-9a-zA-Z_-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)
	
build:
	go build -o bin/leetcode-picker cmd/leetcode-picker/main.go

.DEFAULT_GOAL := help