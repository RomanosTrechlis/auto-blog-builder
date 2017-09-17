FROM golang:alpine

# add git support to golang alpine
RUN apk add --update --no-cache git

ADD . /go/src/github.com/RomanosTrechlis/auto-blog-builder/generate

WORKDIR /artifacts

RUN go install github.com/RomanosTrechlis/auto-blog-builder/generate

ENTRYPOINT /go/bin/generate

EXPOSE 50051