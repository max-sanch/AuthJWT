FROM golang:1.16-buster

RUN go version
ENV GOPATH=/

WORKDIR /go/src/app
COPY . .

RUN go mod download
RUN go build -o auth_jwt ./cmd/main.go

CMD ["./auth_jwt"]