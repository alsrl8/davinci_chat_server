# Use the official Golang image to create a build artifact.
# This is based on Debian and sets the GOPATH to /go.
FROM golang:1.20 as builder

# Copy local code to the container image.
WORKDIR /app
COPY . .

# Build the Go app.
RUN CGO_ENABLED=0 GOOS=linux go build -v -o server

# Use the official Alpine image for a lean production container.
# https://hub.docker.com/_/alpine
FROM alpine:latest
RUN apk --no-cache add ca-certificates

# Copy the binary to the production image from the builder stage.
COPY --from=builder /app/server /server

# Run the web service on container startup.
CMD ["/server"]
