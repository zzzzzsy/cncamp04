FROM golang:alpine3.16 as builder

WORKDIR /workspace

COPY . .
RUN go build -o httpclient


FROM alpine:3.16

WORKDIR /usr/app

COPY --from=builder /workspace/httpclient ./

ENTRYPOINT ["/usr/app/httpclient"]