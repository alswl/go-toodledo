FROM golang:1.22-buster as builder

RUN sed -i "s@http://.*.debian.org@https://mirrors.tuna.tsinghua.edu.cn@g" /etc/apt/sources.list
RUN apt update && apt install -y bash make curl git binutils findutils libreadline7 upx && apt clean && rm -rf /var/apt/cache/*

WORKDIR /app

ENV GOPROXY="https://goproxy.cn,https://mirrors.aliyun.com/goproxy,direct"
ENV GOPRIVATE=""
ENV GOFLAGS=""

# all go mod id downloaded in build-env image
COPY Makefile .
COPY go.mod .
COPY go.sum .
RUN make download
