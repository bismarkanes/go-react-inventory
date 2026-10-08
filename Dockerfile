FROM golang:1.26-alpine AS builder

WORKDIR /app

# Copy EVERYTHING first because 'go mod tidy' needs to scan your .go files
COPY . .

# Detects used/unused packages, updates go.mod, and downloads them
RUN go mod tidy

# Build the binary
RUN CGO_ENABLED=0 GOOS=linux go build -o main .

# Minimal runtime stage
FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/main .

CMD ["./main"]
