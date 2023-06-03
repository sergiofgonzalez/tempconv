PATH_TO_CHROME := "/opt/google/chrome/chrome"


setup:
	go mod tidy
.PHONY: setup

format:
	go fmt ./...
.PHONY: format

clean:
	rm -rf bin/
.PHONY: clean

lint:
	golint ./...
.PHONY: lint

test:
	go clean -testcache && go test -race $$(go list ./... | grep -v tests | grep -v mocks) -coverprofile=coverage.out -covermode=atomic
.PHONY: test

test-v:
	go clean -testcache && go test -race -v $$(go list ./... | grep -v tests | grep -v mocks) -coverprofile=coverage.out -covermode=atomic
.PHONY: test-v

cover:
	go tool cover -html=coverage.out -o ./.coverage.html
	${PATH_TO_CHROME} file://${PWD}/.coverage.html &
.PHONY: cover

doc:
	godoc -http localhost:8080
.PHONY: doc

all:
	$(MAKE) format
	$(MAKE) list
	$(MAKE) test