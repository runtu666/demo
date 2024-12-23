FROM jfrog.foxitsoftware.com/docker-hub-proxy/golang:alpine AS builder

RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories
RUN apk update --no-cache && apk add --no-cache tzdata
LABEL stage=gobuilder

ENV CGO_ENABLED 0
ENV GOPROXY https://goproxy.cn,direct

WORKDIR /build

ADD go.mod .
ADD go.sum .
COPY . .
RUN go mod download
RUN go build -ldflags="-s -w" -o /app/cmd  api/main.go


FROM jfrog.foxitsoftware.com/docker-hub-proxy/alpine

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=builder /usr/share/zoneinfo/Asia/Shanghai /usr/share/zoneinfo/Asia/Shanghai
ENV TZ Asia/Shanghai

WORKDIR /app
COPY --from=builder /app/cmd /app/cmd
COPY job/etc /app/etc

CMD ["./cmd", "-f", "etc/job.yaml"]
