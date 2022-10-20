FROM golang:alpine AS builder

LABEL stage=gobuilder

ENV CGO_ENABLED 0
ENV GOOS linux
ENV GOPROXY https://goproxy.cn,direct

WORKDIR /build/zero

ADD go.mod .
ADD go.sum .
RUN go mod download
COPY . .
RUN go build -ldflags="-s -w" -o /app/api ./app/app.go


FROM alpine

RUN mkdir -p /app/etc
ENV TZ Asia/Shanghai

WORKDIR /app
COPY --from=builder /app/api /app/api
COPY /app/etc/app.yaml /app/etc/app.yaml
COPY /app/account.json /app/account.json


EXPOSE 8888
CMD  ./api