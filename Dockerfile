FROM golang:1.23 AS base

RUN mkdir -p /home/gopher \
    && groupadd --gid 1000 gopher \
    && useradd --uid 1000 --gid 1000 -m gopher \
    && chown -R 1000:1000 /home/gopher

USER gopher
ENV GOPATH="/home/gopher"
ENV PATH="${GOPATH}/bin:${PATH}"

FROM base AS dev
RUN go install  "golang.org/x/tools/gopls@latest"
RUN go install "github.com/tpng/gopkgs@latest"
RUN go install "github.com/ramya-rao-a/go-outline@latest"
RUN go install "honnef.co/go/tools/cmd/staticcheck@latest"
RUN go install "github.com/go-delve/delve/cmd/dlv@latest"
RUN go install "github.com/swaggo/swag/cmd/swag@latest"
WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go mod download