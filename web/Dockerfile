FROM golang:alpine

# add git support to golang alpine
RUN apk add --update --no-cache git

# Copy the local package files to the container's workspace.
ADD . /go/src/github.com/RomanosTrechlis/auto-blog-builder/web

# Adding workdir in order for the container to see static files
WORKDIR ../artifacts

COPY config.json ../artifacts/config.json

# Build the outyet command inside the container.
# (You may fetch or manage dependencies here,
# either manually or with a tool like "godep".)
RUN go install github.com/RomanosTrechlis/auto-blog-builder/web

# Run the outyet command by default when the container starts.
ENTRYPOINT /go/bin/web

# Document that the service listens on port 8080.
EXPOSE 8080
