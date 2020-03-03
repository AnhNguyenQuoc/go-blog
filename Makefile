GOCMD=go
GODEPCMD=godep
GOBUID=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GORUN=CompileDaemon --build="go build -o main ." --command=./main
BINARY_NAME=main
BINARY_UNIX=$(BINARY_NAME)_unix

all: test buid

build:
	$(GOBUILD) -o $(BINARY_NAME) . -v

test:
	$(GOTEST) -v ./...

clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)
	rm -f $(BINARY_UNIX)

run:
	$(GORUN)

build-linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(BINARY_NAME) -v .

docker-build:
	docker-compose up --build