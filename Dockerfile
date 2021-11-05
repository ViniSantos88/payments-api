# Use the official Golang image to create a build artifact.
# This is based on Debian and sets the GOPATH to /go.
# https://hub.docker.com/_/golang
FROM golang:1.16 as builder

# Create and change to the app directory.
WORKDIR /app

# TODO: SECRET MANAGER 
#RUN go get github.com/GoogleCloudPlatform/berglas
#RUN go install github.com/GoogleCloudPlatform/berglas

# Retrieve application dependencies.
# This allows the container build to reuse cached dependencies.
COPY go.* ./
RUN go mod download


# Copy local code to the container image.
COPY . ./

# Build the binary.
RUN CGO_ENABLED=0 GOOS=linux go build -v -o server

# Use the official Alpine image for a lean production container.
# https://hub.docker.com/_/alpine
# https://docs.docker.com/develop/develop-images/multistage-build/#use-multi-stage-builds
FROM alpine:3
RUN apk add --no-cache ca-certificates

# Copy the binary to the production image from the builder stage.
COPY --from=builder /app/server /server
COPY --from=builder /app/swagger.yaml ./
# Run the web service on container startup.
#CMD ["/server"]
CMD ["/server"]