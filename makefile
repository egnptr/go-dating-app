export PKGS=$(shell go list ./...)

lint:
	@golangci-lint run $(PKGS) --disable=errcheck --timeout 10m

test:
	@go test -v -cover -race $(PKGS)

build:
	@go build -v -o ./app/go-dating-app ./app/app.go

run:
	make build
	@./app/go-dating-app

race: 
	@go run -race ./app/app.go