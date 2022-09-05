FROM golang:1.14-buster

RUN go version
ENV GOPATH=/

COPY ./ ./


# build go app
RUN go mod download
RUN go build -o todo-app ./cmd/main.go

CMD ["./todo-app"]