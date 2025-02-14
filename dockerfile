FROM golang:1.23

RUN go install github.com/go-swagger/go-swagger/cmd/swagger@latest
