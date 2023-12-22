FROM golang:1.16-alpine

WORKDIR /app

COPY /go.mod ./go.mod
COPY /go.sum ./go.sum

RUN go mod download

COPY /cmd/driver/main.go ./cmd/driver/main.go
COPY /internal/driver/ ./internal/driver/
COPY /modals/ ./modals/

RUN go build -o /driver-service ./cmd/driver/main.go

EXPOSE 8080

CMD ["/driver-service"]