# docker-nginx-go

在容器在通过nginx代理go的示例

运行:

确保本地 8088端口可用
```
docker-compose up
```

运行 `http://localhost:8088/hello`看一下效果


###tips:

因为go的服务和nginx属于两个容器, 而且ip其实并不固定, 因此需要用到nginx 的 `upstream`方法进行tcp转发, 而不是http转发

详情参考 `nginx/conf.d/default.conf` 和 `nginx/conf.d/upstream.conf`



