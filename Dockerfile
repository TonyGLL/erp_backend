# Stage 1: Build stage
FROM golang:1.22-alpine AS build

# Set the working directory
WORKDIR /app

# Copy the source code
COPY . .

# Installing build dependencies. 
RUN apk add --no-progress --no-cache gcc musl-dev

# Copy and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Build the Go application
RUN CGO_ENABLED=0 GOOS=linux go build -tags musl -ldflags '-extldflags "-static"' -o main .

# Stage 2: Final stage
FROM scratch

# Set the working directory
WORKDIR /app

# Copy the binary from the build stage
COPY --from=build /app/main .

# Copy the environment configuration file
ARG CONFIG_FILE=dev.env
COPY ${CONFIG_FILE} /app/${CONFIG_FILE}

# Set the entrypoint command
ENV CONFIG_FILE=${CONFIG_FILE}
ENTRYPOINT ["./main"]
