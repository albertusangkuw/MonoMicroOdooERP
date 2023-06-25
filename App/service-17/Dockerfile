# syntax=docker/dockerfile:1
# Start from a Go image
FROM golang:1.20.4

# Set destination for COPY
WORKDIR /app

# Download Go modules
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code. Note the slash at the end, as explained in
# https://docs.docker.com/engine/reference/builder/#copy
COPY . .

# Build
RUN go build -o service-artifact

EXPOSE 1324

# Set the entrypoint command
CMD ["./service-artifact" , "-env=deploy.env" ]
