FROM golang:1.20.4-alpine

RUN go version
ENV GOPATH=/

COPY ./ ./

# build go app
RUN go mod download
RUN go build -o genesis-project ./cmd/main.go

CMD ["./genesis-project"]