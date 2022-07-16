FROM golang:alpine3.16 as builder

WORKDIR /workspace

COPY . .
RUN go build -o httpclient


FROM alpine:3.16

WORKDIR /usr/app

ENV CONFIG_PATH=/usr/app/config.json

COPY --from=builder /workspace/httpclient ./
COPY module08/conf/config.json ./

ENTRYPOINT ["/usr/app/httpclient"]