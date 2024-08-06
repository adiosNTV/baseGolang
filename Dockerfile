# Stage 1: Build stage
FROM golang:1.22.1-alpine AS builder

WORKDIR /app

# Copy and download dependencies
COPY go.mod go.sum ./
RUN go mod download && go mod verify

# Copy the source code and set environments
COPY ./cmd ./cmd
COPY ./db ./db
COPY ./internal ./internal
COPY ./cache ./cache
COPY ./utils ./utils

COPY ./.env ./.env
ENV GO111MODULE=on
ENV GOCACHE=/root/.cache/go-build

# Builds the application as a staticly linked one, to allow it to run on alpine.
RUN --mount=type=cache,target="/root/.cache/go-build" CGO_CFLAGS_ALLOW=-Xpreprocessor GOOS=linux go build -a -installsuffix cgo -o apiserver ./cmd/api

# Stage 2: Final stage
FROM alpine:edge
COPY --from=builder ["/app/apiserver", "/app/.env" ,"/"]

# Set the entrypoint command
ENTRYPOINT ["/apiserver"]
EXPOSE 3000
