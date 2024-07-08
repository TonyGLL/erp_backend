# Stage 1: Build stage
FROM golang:1.22-alpine AS build

# Set the working directory
WORKDIR /app

# Copy and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code
COPY . .

# Build the Go application
RUN CGO_ENABLED=0 GOOS=linux go build -o main .

# Stage 2: Final stage
FROM alpine:edge

# Set the working directory
WORKDIR /app

# Copy the binary from the build stage
COPY --from=build /app/main .

# Copy the environment configuration file
ARG CONFIG_FILE=dev.env
COPY ${CONFIG_FILE} /app/${CONFIG_FILE}

# Set the timezone and install CA certificates
RUN apk --no-cache add ca-certificates tzdata

# Set the entrypoint command
ENV CONFIG_FILE=${CONFIG_FILE}
ENTRYPOINT ["./main"]
