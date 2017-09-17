FROM golang:alpine

# add git support to golang alpine
RUN apk add --update --no-cache git

ADD . /go/src/github.com/RomanosTrechlis/auto-blog-builder/fetch

WORKDIR /artifacts

RUN go install github.com/RomanosTrechlis/auto-blog-builder/fetch

ENTRYPOINT /go/bin/fetch

EXPOSE 50051