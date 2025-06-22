GO := $(shell which go)

.PHONY: run

run:
	@$(GO) run .

