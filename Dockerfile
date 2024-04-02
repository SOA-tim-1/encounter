# Build stage
FROM golang:alpine AS build-env

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o main .

# Run stage
FROM alpine:latest AS final

WORKDIR /app

# Set environment variables
ENV DATABASE_HOST=host.docker.internal
ENV DATABASE_PORT=5432
ENV DATABASE_SCHEMA=explorer-v1
ENV DATABASE_SCHEMA_NAME=encounters
ENV DATABASE_USERNAME=postgres
ENV DATABASE_PASSWORD=super

# Copy the binary from the build stage
COPY --from=build-env /app/main .

CMD ["./main"]