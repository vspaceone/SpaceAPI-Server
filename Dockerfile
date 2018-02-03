FROM golang:1.9

WORKDIR /go/src/github.com/vspaceone/SpaceAPI-Server
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

ENTRYPOINT ["SpaceAPI-Server", "serve"]