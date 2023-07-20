# 项目部署（Docker）
##  前端

1) 有如下两种方法获取获取镜像

* 从**DockerHub**拉取

```
	docker pull 
```

* 根据**Dockerfile**自行构建


```shell
	# 将用源代码构建好的dist文件夹放在本GitHub项目/Docker/frontend中，在该路径下
	docker build -t frontend:latest .
```

2) 接着可以运行服务
```shell
	docker run -p [PORT]:80 frontend:latest
	# 将[PORT]替换为主机的某个可用端口，服务将会通过这个端口进行通信
```