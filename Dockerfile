FROM golang:latest as go-app

LABEL maintainer="nguyenquocanh121296@gmail.com"

ENV GOPATH=/go

ENV PATH=$GOPATH/bin:$PATH

RUN mkdir -p $GOPATH/src/github.com/AnhNguyenQuoc/go-blog

WORKDIR $GOPATH/src/github.com/AnhNguyenQuoc/go-blog

COPY go.mod .

COPY go.sum .

RUN go mod download

COPY . .

RUN go get github.com/githubnemo/CompileDaemon

RUN go build -o main .

VOLUME $GOPATH/src/github.com/AnhNguyenQuoc/go-blog

EXPOSE 3000

ENTRYPOINT CompileDaemon --build="go build -o main ." --command=./main