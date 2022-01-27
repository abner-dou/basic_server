# Building stage
FROM golang:1.16-alpine3.13 as builder

RUN mkdir - /go/src/basic_server
WORKDIR /go/src/basic_server

ENV GO111MODULE=on

COPY ./ ./

RUN apk add --no-cache git curl build-base bash openssh-client shadow


RUN go env -w GOFLAGS=-mod=mod

RUN go mod download

RUN go build -o /basic_server .

# Final image
FROM alpine:3.13.5

COPY --from=builder /basic_server .
COPY --from=builder /go/src/basic_server/html ./html

CMD ["./basic_server"]