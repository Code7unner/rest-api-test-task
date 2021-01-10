FROM golang:1.14 AS builder
ENV GO111MODULE=on

ARG ENVCONSUL_VERSION=0.9.3

# first stage [test and build]
ADD deployments /go/builder
WORKDIR /go/builder
RUN apt-get update \
    && apt-get install -y curl
RUN curl -sf -LO https://github.com/hashicorp/envconsul/archive/v${ENVCONSUL_VERSION}.tar.gz \
    && tar -xvf v${ENVCONSUL_VERSION}.tar.gz \
    && cd envconsul-${ENVCONSUL_VERSION} \
    && make linux/amd64 \
    && mv pkg/linux_amd64/envconsul /usr/local/bin/envconsul

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -mod=mod -o service main.go

# final stage
FROM alpine
ENV SERVICE_PORT=8080
COPY --from=builder /usr/local/bin/envconsul /usr/local/bin/envconsul
COPY --from=builder /go/builder/service /go/bin/service
RUN apk update \
    && apk add tzdata \
    && mkdir /lib64 \
    && ln -s /lib/libc.musl-x86_64.so.1 /lib64/ld-linux-x86-64.so.2
WORKDIR /go/bin
ENTRYPOINT ["/usr/local/bin/envconsul"]