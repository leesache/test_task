# Stage 1:
FROM golang:1.23.3-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download && go mod verify
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o main .

# Stage 2:
FROM alpine:latest
RUN apk add --no-cache bash
WORKDIR /root/
# Copy the binary from the builder stage
COPY --from=builder /app/migrations ../../migrations
COPY --from=builder /app/main .
EXPOSE 9090
CMD ["./main"]