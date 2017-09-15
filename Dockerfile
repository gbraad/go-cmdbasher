FROM fedora:26

MAINTAINER Gerard Braad <me@gbraad.nl>

ENV GOPATH /workspace
RUN dnf install -y gcc golang git

COPY . /workspace/src/github.com/gbraad/go-cmdbasher
WORKDIR /workspace/src/github.com/gbraad/go-cmdbasher

RUN go get -v -d ./...
RUN go test
RUN GOOS=windows go install -v ./cmd/basher
