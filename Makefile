.DEFAULT_GOAL=explain
.PHONY: explain
explain:
	#          ____   _____   ___   ___ ___  __
	#    /\   / __ \ / ____| |__ \ / _ \__ \/_ |
	#   /  \ | |  | | |         ) | | | | ) || |
	#  / /\ \| |  | | |        / /| | | |/ / | |
	# / ____ \ |__| | |____   / /_| |_| / /_ | |
	#/_/    \_\____/ \_____| |____|\___/____||_|
	#
	### Targets
	@cat Makefile* | grep -E '^[a-zA-Z_-]+:.*?## .*$$' | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.PHONY: install
install: ## Install dependencies
	go install github.com/golang/mock/mockgen@v1.6.0
	go get ./...

.PHONY: test
test: ## Test Go code
	go test ./...

.PHONY: generate-mocks
generate-mocks: ## Generate mocks from the interfaces
	go generate ./...
