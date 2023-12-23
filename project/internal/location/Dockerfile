FROM golang:1.17-alpine

WORKDIR /app

COPY /go.mod ./go.mod
COPY /go.sum ./go.sum

RUN go mod download

COPY /cmd/location/main.go ./cmd/location/main.go
COPY /internal/location/ ./internal/location/
COPY /modals/ ./modals/
COPY /env.dev ./env.dev

RUN go build -o /location-service ./cmd/location/main.go

EXPOSE 8080

CMD ["/location-service"]