# cncamp04

## Module01
基于 Channel 编写一个简单的单线程生产者消费者模型：
队列：
队列长度 10，队列元素类型为 int
生产者：
每 1 秒往队列中放入一个类型为 int 的元素，队列满时生产者可以阻塞
消费者：
每一秒从队列中获取一个元素并打印，队列为空时消费者阻塞

## Module02
编写一个 HTTP 服务器，大家视个人不同情况决定完成到哪个环节，但尽量把 1 都做完：
接收客户端 request，并将 request 中带的 header 写入 response header
读取当前系统的环境变量中的 VERSION 配置，并写入 response header
Server 端记录访问日志包括客户端 IP，HTTP 返回码，输出到 server 端的标准输出
当访问 localhost/healthz 时，应返回 200

## Module03
构建本地镜像
编写 Dockerfile 将练习 2.2 编写的 httpserver 容器化
将镜像推送至 docker 官方镜像仓库
通过 docker 命令本地启动 httpserver
通过 nsenter 进入容器查看 IP 配置
```
docker build -f Dockerfile . -t zzzzzsy/cncamp04:1.0
docker login -u zzzzzsy
docker push zzzzzsy/cncamp04:1.0

# dockerhub image https://hub.docker.com/repository/docker/zzzzzsy/cncamp04
```