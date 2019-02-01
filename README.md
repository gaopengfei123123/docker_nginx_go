# docker-nginx-go

在容器在通过nginx代理go的示例

运行:

确保本地 8088端口可用
```
docker-compose up
```

运行 `http://localhost:8088/hello`看一下效果


###tips:

因为go的服务和nginx属于两个容器, 而且ip其实并不固定, 因此需要用到nginx 的 `upstream`方法进行基于容器别名的转发, 而不是ip转发

详情参考 `nginx/conf.d/default.conf` 和 `nginx/conf.d/upstream.conf`

###使用https

可以参考 `nginx/conf.d/default.conf.example`里面的写法, 证书文件放在`nginx/cert/`底下

### 如何新增一个go服务
1. 先把代码文件放到 `www`目录下, 和 `$GOPATH/src` 一样, 而且要把外部依赖文件放到vendor包里面
2. 然后在 `docker-compose.yml` 参考`example_server`写一个server
3. 然后在`nginx/site` 底下加上相应的nginx配置, **要加上upstream的配置, 参考`nginx/conf.d/upstream.cof`

