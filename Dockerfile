# syntax=docker/dockerfile:1
# Following best practices mentioned in:
# https://www.saybackend.com/blog/02-golang-dockerfile/

###################################
# STAGE 1: Install Dependencies   #
###################################
FROM golang:1.22.4-bookworm AS deps

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

###################################
# STAGE 2: Build Application      #
###################################
FROM golang:1.22.4-bookworm AS builder

WORKDIR /app

COPY --from=deps /go/pkg /go/pkg
COPY . .

# Enable them if you need them
# Setting CGO_ENABLED=0 and GOOS=linux ensures that the binary is statically linked and compatible with Linux.
# ENV CGO_ENABLED=0
# ENV GOOS=linux

# Using -ldflags="-w -s" to strip the binary of debug information and symbols, which helps in reducing the size of the binary.
RUN go build -ldflags="-w -s" -o main ./cmd/api/

###################################
# STAGE 3: Run Application        #
###################################
# https://martinheinz.dev/blog/92
FROM debian:bookworm-slim

RUN apt-get update && apt-get install -y curl

WORKDIR /app

# Create a non-root user and group
RUN groupadd -r appuser && useradd -r -g appuser appuser

# Copy the built application
COPY --from=builder /app/main .

# Change ownership of the application binary
RUN chown appuser:appuser /app/main

# Switch to the non-root user
USER appuser

# Set NODE_ENV again to ensure the app runs in production mode
ENV ENV=production

CMD ["./main"]
