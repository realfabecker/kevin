FROM golang:1.25.0-alpine3.21 AS base
RUN apk add --no-cache --update \
    git \
    make \
    nodejs \
    npm \
    openjdk17-jdk \
    && addgroup -g 1000 -S nonroot \
    && adduser -D -G nonroot -u 1000 nonroot

FROM base AS dev

USER nonroot
ENV GOPATH="/home/nonroot"
ENV PATH="${GOPATH}/bin:${PATH}"

RUN GODEBUG=http2client=0,netdns=go+1m GOPROXY=https://proxy.golang.org,direct go install "golang.org/x/tools/gopls@latest" \
    && GODEBUG=http2client=0,netdns=go+1m GOPROXY=https://proxy.golang.org,direct go install "honnef.co/go/tools/cmd/staticcheck@latest" \
    && GODEBUG=http2client=0,netdns=go+1m GOPROXY=https://proxy.golang.org,direct go install "github.com/go-delve/delve/cmd/dlv@latest"

WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go mod download