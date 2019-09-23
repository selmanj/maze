PKGS := $(shell go list ./... | grep -v /vendor)

maze: 
	go build github.com/selmanj/maze/cmd/maze


.PHONY: test
test:
	go test $(PKGS)