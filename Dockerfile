# syntax=docker/dockerfile:1

FROM golang:1.20

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . ./

# Build
RUN CGO_ENABLED=0 GOOS=linux go build -o /go/bin/restauAPI ./cmd/app/main.go

EXPOSE 8080

# Run
CMD ["/go/bin/restauAPI"]