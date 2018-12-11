# 运行容器
FROM alpine:latest AS runner
RUN apk --no-cache add ca-certificates
WORKDIR /www/site/
COPY ./app/app ./app
EXPOSE 9090
CMD [ "./app" ]