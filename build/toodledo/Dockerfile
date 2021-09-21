FROM golang:1.16.5-buster as builder
# build-env contains all dependencies
#

WORKDIR app

ENV GOPROXY="https://goproxy.cn,https://mirrors.aliyun.com/goproxy,direct"
ENV GOPRIVATE=""
ENV GOFLAGS=""

RUN export CGO_ENABLED=0 GOOS=linux GOARCH=amd64

# all go mod id downloaded in build-env image
COPY Makefile .
COPY go.mod .
COPY go.sum .
RUN git init
RUN make download

COPY . .
RUN make build

FROM alpine:3.7
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.tuna.tsinghua.edu.cn/g' /etc/apk/repositories
RUN apk update && apk add --no-cache bash make curl git util-linux util-linux-doc binutils findutils readline

WORKDIR /

COPY --from=builder /app/bin/toodledo /toodledo

ENTRYPOINT ["/toodledo"]
