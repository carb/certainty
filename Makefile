PROJECT_ROOT = github.com/carb/certainty

PROGS = cmd/explore/explore

explore: $(shell find . -not -path "./vendor/*" -name "*.go") glide.lock
	go build -o cmd/explore/explore cmd/explore/main.go

lint:: errcheck staticcheck unused

.PHONY: run
run: bins
	@./certainty
