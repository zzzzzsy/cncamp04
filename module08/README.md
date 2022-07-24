## 需求完成
- 通过增加livenessProb以及readinessProb提供了探活以及优雅启动的能力
- 使用`gopkg.in/tylerb/graceful.v1`为程序提供了优雅终止的能力
- 资源需求和 QoS 保证
- 通过`github.com/sirupsen/logrus`包丰富了程序的日志等级，默认`info`，可通过配置文件修改日志等级，提供debug日志
- 增加Service，Nginx Ingress
- 通过letsencrypt签发的证书保证服务的通讯安全
- 增加configmap使配置与代码分离
- 更新自动化部署文件Makefile

## Pre Request
### 设置Dockerhub登录信息(可选)
```
# 如果需要push image则需要设置登录dockerhub 登录信息
export DOCKER_PASSWORD=FIXME
export DOCKER_USER=FIXME
```

### 安装NGINX ingress controller
```
kubectl apply -f https://raw.githubusercontent.com/kubernetes/ingress-nginx/controller-v1.3.0/deploy/static/provider/cloud/deploy.yaml
```

### 安装cert-manager
```
helm repo add jetstack https://charts.jetstack.io

helm repo update

kubectl apply -f https://github.com/cert-manager/cert-manager/releases/download/v1.8.2/cert-manager.crds.yaml

helm install \
  cert-manager jetstack/cert-manager \
  --namespace cert-manager \
  --create-namespace \
  --version v1.8.2

# 安装成功
❯ k get pod -n cert-manager

NAME                                       READY   STATUS    RESTARTS   AGE
cert-manager-67599dcc49-nwr86              1/1     Running   0          3m19s
cert-manager-cainjector-58c8955c5d-hm847   1/1     Running   0          3m19s
cert-manager-webhook-d698b4885-kkbff       1/1     Running   0          3m19s
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

#根据`module08/manifests/`里到文件 发布到k8s集群
#发布前请确保您当前的k8s集群是目标集群
make apply

#清理
#删除cncamp namespace下所有资源
make cleanup
```

## 测试

### 第二部分
- 验证certificate安装
![Alt text](../img/cert.jpg?raw=true "letsencrypt")

- 通过浏览器最终验证
![Alt text](../img/securehttps.jpg?raw=true "final-test")

### 第一部分

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

#日志分级信息
time="2022-07-17 10:57:33" level=debug msg="Response header Accept: */*\n"
time="2022-07-17 10:57:33" level=debug msg="Response header Connection: close\n"
time="2022-07-17 10:57:33" level=debug msg="Response header User-Agent: kube-probe/1.23\n"
time="2022-07-17 10:57:33" level=info msg="Request is from 172.17.0.1:57046\n"
time="2022-07-17 10:57:33" level=info msg="Response code is 200\n"
time="2022-07-17 10:57:33" level=debug msg="Response header Connection: close\n"
time="2022-07-17 10:57:33" level=info msg="Request is from 172.17.0.1:57048\n"
time="2022-07-17 10:57:33" level=info msg="Response code is 200\n"
^C
```
