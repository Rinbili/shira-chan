# 项目部署（Docker）
##  前端

1) 有如下两种方法获取获取镜像

* 从[**DockerHub**]()拉取

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
## 数据库
1) 如上所述可以从[这里](https://baidu.com)获取打包好的镜像, 也可在/Docker/backend/database中自行构建:

```shell
	# 先进入mysql.preseed中修改[PASSWORD]的值为密码，否则默认密码0000!!!
	docker build -t db:latest .
```

2) 获取镜像后即可启动服务，如果是下载我们的镜像记得修改默认密码0000以防安全问题。

```shell
	docker run -p [PORT]:3306 db:latest
	# 将[PORT]的值替换为主机的某个可用端口，数据库服务将会通过这个端口进行通信
```  
