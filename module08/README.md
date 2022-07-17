## 需求完成
- 通过增加livenessProb以及readinessProb提供了探活以及优雅启动的能力
- 配置文件与可执行文件分离
- 使用`gopkg.in/tylerb/graceful.v1`为程序提供了优雅终止的能力
- 资源需求和 QoS 保证
- 通过`github.com/sirupsen/logrus`包丰富了程序的日志等级，默认`info`，可通过配置文件修改日志等级，提供debug日志

## Pre Request
### 设置Dockerhub登录信息(可选)
```
# 如果需要push image则需要设置登录dockerhub 登录信息
export DOCKER_PASSWORD=FIXME
export DOCKER_USER=FIXME
```

## 执行代码
```
#本地测试
make test

#生成docker镜像
make docker-build

#上传镜像
make docker-push

#本地启动docker镜像
make docker-run

#根据配置文件`module08/conf/config.json` 重新生成deployment文件
#deployment模版在`module08/template`目录
#目前支持 1.修改端口  2.镜像版本
make manifests

#根据`module08/manifests/`里到文件 发布到k8s集群
#发布前请确保您当前的k8s集群是目标集群
make apply
```

## 测试

```
#根据当前service的端口，将流量转发到本地进行测试
#目前service的端口配置在80
#将流量转发到本地的8443
kubectl port-forward svc/httpclient-service 8443:80

#测试结果
❯ curl http://localhost:8443/ip
127.0.0.1:55482%
❯ curl http://localhost:8443/healthz
Hello LiveRamp SRE%
```
