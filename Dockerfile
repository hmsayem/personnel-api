# Start from the latest golang base image
FROM golang:latest as builder

# Add maintainer info
LABEL maintainer="Hossain Mahmud <hmsayem@gmail.com>"

#Set the current working directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o rest-server .

# Start a new stage from scratch
FROM alpine:latest

WORKDIR /root

COPY --from=builder /app/rest-server .

# Command to run the executable
CMD ["./rest-server"]
