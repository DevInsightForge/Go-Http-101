# Builder stage
FROM golang:1.22-alpine AS builder
WORKDIR /app

COPY go.mod go.sum /app/
RUN set -Eeux && \
    go mod download && \
    go mod verify

COPY . /app/
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 \
    go build -trimpath -o gohttp101 cmd/main.go

# Final stage
FROM scratch
WORKDIR /app

COPY --from=builder /app/gohttp101 /app/main
ENV ADDRESS=0.0.0.0
ENV PORT 4000
EXPOSE 4000

CMD ["./main"]
