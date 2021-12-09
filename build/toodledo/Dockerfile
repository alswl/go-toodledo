FROM golang:1.16.5-buster as builder
# TODO build-env contains all dependencies

WORKDIR /app

ENV GOPROXY="https://goproxy.cn,https://mirrors.aliyun.com/goproxy,direct"
ENV GOPRIVATE=""
ENV GOFLAGS=""

# all go mod id downloaded in build-env image
COPY Makefile .
COPY go.mod .
COPY go.sum .
RUN make download

COPY . .
RUN make build GOOS=linux GOARCH=amd64

FROM golang:1.16.5-buster
RUN sed -i "s@http://.*.debian.org@https://mirrors.tuna.tsinghua.edu.cn@g" /etc/apt/sources.list
RUN apt update && apt install -y bash make curl git binutils findutils libreadline7 && apt clean && rm -rf /var/apt/cache/*

WORKDIR /

COPY --from=builder /app/bin/toodledo /toodledo

ENTRYPOINT ["/toodledo"]