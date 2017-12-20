FROM golang:1.9.2-alpine AS builder

WORKDIR /go/src/github.com/porty/pubserve
COPY . .
RUN CGO_ENABLED=0 go-wrapper install -ldflags '-extldflags "-static"'

FROM scratch

COPY --from=builder /go/bin/pubserve /
ENTRYPOINT ["/pubserve"]
