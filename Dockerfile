# Builder stage
FROM golang:1.22-alpine AS builder
WORKDIR /app
COPY go.mod go.sum /app/
RUN set -Eeux && \
    go mod download && \
    go mod verify
COPY . /app/
ENV PORT 4000
RUN GOOS=linux GOARCH=amd64 \
    go build \
    -trimpath \
    -o main cmd/main.go

# Final stage
FROM scratch
WORKDIR /app
COPY --from=builder /app/main .
EXPOSE 4000

# Use sh to execute the main binary
CMD ["./main"]
