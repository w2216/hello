```
// 构建镜像
docker build -t hello ./ -f services/hello/api/Dockerfile

// 运行
docker-compose up -d
```