FROM golang:1.24.0-alpine3.21 AS base

RUN apk add --no-cache git make
RUN addgroup -g 1000 -S gopher \
    && adduser -D -G gopher -u 1000 gopher

USER gopher
ENV GOPATH="/home/gopher"
ENV PATH="${GOPATH}/bin:${PATH}"

FROM base AS dev
RUN GODEBUG=http2client=0,netdns=go+1m GOPROXY=https://proxy.golang.org,direct go install "golang.org/x/tools/gopls@latest"
RUN GODEBUG=http2client=0,netdns=go+1m GOPROXY=https://proxy.golang.org,direct go install "honnef.co/go/tools/cmd/staticcheck@latest"
RUN GODEBUG=http2client=0,netdns=go+1m GOPROXY=https://proxy.golang.org,direct go install "github.com/go-delve/delve/cmd/dlv@latest"
WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go mod download