# 编译环境
FROM golang:1.9-alpine AS builder
WORKDIR /go/src/app/
ARG APP_PATH
COPY ${APP_PATH} .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .


# 运行容器
FROM alpine:latest AS runner
RUN apk --no-cache add ca-certificates
WORKDIR /www/site/
COPY --from=builder /go/src/app/app .
EXPOSE 9090
CMD [ "./app" ]

# # 运行容器
# FROM alpine:latest AS runner
# RUN apk --no-cache add ca-certificates
# WORKDIR /www/site/
# ARG APP_PATH
# COPY ${APP_PATH} ./app
# EXPOSE 9090
# CMD [ "./app" ]