FROM golang:1.13 AS builder
ENV CGO_ENABLED 0
WORKDIR /go/src/app
ADD . .
RUN go build -mod vendor -o /redis-stability-check

FROM scratch
COPY --from=builder /redis-stability-check /redis-stability-check
CMD ["/redis-stability-check"]
