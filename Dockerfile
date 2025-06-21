# Stage 1: Build stage
FROM golang:alpine AS builder

# Setup environment variables
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /app

COPY . .

RUN go mod tidy

RUN go build -o moodle-api .

# Stage 2: Production stage
FROM alpine:latest

# Install required packages
RUN apk add --no-cache tzdata

ENV TZ=Asia/Jakarta

# Set timezone
RUN ln -sf /usr/share/zoneinfo/$TZ /etc/localtime && \
    echo $TZ > /etc/timezone

WORKDIR /app

# Create user for security
RUN addgroup -g 1001 binarygroup
RUN adduser -D -u 1001 -G binarygroup userapp

COPY --from=builder --chown=userapp:binarygroup /app/moodle-api .
#COPY --chown=userapp:binarygroup .env .env

USER userapp

EXPOSE 9083

CMD ["./moodle-api"]