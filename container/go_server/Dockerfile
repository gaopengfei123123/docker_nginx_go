# 编译容器
FROM golang:1-alpine AS builder

ARG APP_PATH
WORKDIR /go/src/go-server/
COPY main.go .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .

# 运行容器
FROM alpine:latest AS runner
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /go/src/go-server/app .
EXPOSE 9090
CMD [ "./app" ]