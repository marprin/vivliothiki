pkgs          = $(shell go list ./... | grep -v /tests | grep -v /vendor/ | grep -v /common/)

test:
	@echo " >> running tests"
	@go test  -cover $(pkgs)
