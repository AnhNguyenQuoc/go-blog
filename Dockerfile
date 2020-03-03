FROM golang:latest as go-app

LABEL maintainer="nguyenquocanh121296@gmail.com"

ENV GOPATH=/go

ENV PATH=$GOPATH/bin:$PATH

RUN mkdir -p $GOPATH/src/app

WORKDIR $GOPATH/src/app

COPY . .

RUN go get github.com/tools/godep && godep restore

RUN go get github.com/githubnemo/CompileDaemon

RUN go build -o main .

VOLUME $GOPATH/src/app

EXPOSE 3000

ENTRYPOINT CompileDaemon --build="go build -o main ." --command=./main