FROM golang:1.17.6-bullseye AS builder

WORKDIR /root/build

ADD go.mod .
ADD go.sum .
RUN go mod download

ADD . .
RUN go build -o serve github.com/ehsundar/insta/cmd/serve


FROM ubuntu:20.04 AS runner

WORKDIR /bin
COPY --from=builder /root/build/serve .

EXPOSE 8080
