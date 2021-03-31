FROM golang:latest
RUN go get -u golang.org/x/tools/cmd/cover
WORKDIR /src