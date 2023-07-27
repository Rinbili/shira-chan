# 项目部署（Docker）
**注：容器每次启动都会自动拉取最新的前后端文件**
## 数据库(爆炸状态💣💣💣)
1) 有两种方式获取Docker镜像
* 直接从DockerHub拉取

```shell
	docker pull azusaing/db:latest
```

* 也可在/Docker/backend/database中自行构建镜像:
```shell
	# 先将mysql.preseed中两个'0000'的值修改为新密码，否则默认密码0000!!!
	# 接着将start.sh中的两个'0000'修改为新密码
	docker build -t azusaing/db:latest .
```

2) 获取镜像后即可启动服务，如果是下载我们的镜像记得自行修改默认密码0000以防安全问题。

```shell
	docker run -p [PORT]:3306 azusaing/db:latest
	# 将[PORT]的值替换为主机的某个可用端口，数据库服务将会通过这个端口进行通信
```

## 后端（✔️✔️✔️工作状态）
1) 同样可以从/Docker/backend/server自行构建镜像，但是推荐直接拉取

```shell
	docker pull azusaing/server:latest
```

2) 服务配置运行

```shell
	docker run -p [PORT]:8080 azusaing/server:latest
	# 将[PORT]改为与前端通信的端口

	docker exec -it [DOCKER_ID] /bin/bash
	# 将[CONTAINER_ID]改为镜像实例id，在新的终端下运行便会进入容器内部

	cd shira-chan
	vim ./config/config.yml
	# 修改配置文件，注意每个冒号后都要加一个空格

	go run ./cmd/main/main.go
	# 使用这条指令之后服务便会启动在上文[PORT]端口
```

##  前端(✔️✔️✔️工作状态)

1) 同样可以从/Docker/frontend自行构建镜像，但是推荐直接拉取

* 从**DockerHub**拉取

```shell
	docker pull azusaing/frontend:latest
```


2) 接着可以选择直接运行服务，或者配置负载均衡后再运行服务
* 两个选择都需要先配置Docker服务
```shell
	docker run -p [PORT]:80 azusaing/frontend:latest
	# 将[PORT]替换为主机的某个可用端口，网页服务将会通过这个端口进行通信

	docker exec -it [CONTAINER_ID] /bin/bash
	# 将[CONTAINER_ID]替换为容器ID，进入容器
```

* 假如不需要配置负载均衡服务器
```shell
	rm /etc/nginx/sites-enabled/balancer.conf
	# 删除掉不需要的负载均衡文件
	nginx
	# 此时服务即运行在本机[PORT]端口
```

* 假如需要配置负载均衡，在配置了几个上述的下游服务器的基础上，配置负载均衡服务器
```shell
	rm /etc/nginx/sites-enabled/server.conf
	# 删除不需要的文件

	vim /etc/nginx/sites-enabled/balancer.conf
	# 进入负载均衡配置文件添加下游服务器，':wq'保存退出
	
	nginx
	# 此时服务即运行在本机[PORT]端口，作为所有web服务器的入口
```
