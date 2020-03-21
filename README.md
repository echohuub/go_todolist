## Go 语言 Demo

一个用 Go 语言和 [Gin](https://github.com/gin-gonic/gin) 框架实现的简单 TodoList ，包含了对 MySql 数据库的操作。

![](/screenshot/index.png)

### 使用

#### 创建数据库

最简单的方式是使用 Docker ，比如：
```shell script
$ docker run --name mysql \
-p 8993:3306 \
-e MYSQL_ROOT_PASSWORD=qwer1234 \
-d mysql:8.0.3
```
或者使用 Docker Compose
```shell script
$ cat > docker-compose.yaml <<EOF
version: '3'

services:
  mysql:
    image: mysql:8.0.3
    ports:
      - 8993:3306
    volumes:
      - ~/workspace/docker/apps/mysql-8.0.3/conf:/etc/mysql/conf.d
      - ~/workspace/docker/apps/mysql-8.0.3/data:/var/lib/mysql
    environment:
      - TZ=Asia/Shanghai
      - MYSQL_ROOT_PASSWORD=qwer1234
EOF
```
```shell script
$ docker-compose up -d mysql
```

然后创建数据库 `golang`

#### 启动

代码根目录下执行
```shell script
$ go run main.go
```

```shell script
[GIN-debug] [WARNING] Creating an Engine instance with the Logger and Recovery middleware already attached.

[GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
 - using env:	export GIN_MODE=release
 - using code:	gin.SetMode(gin.ReleaseMode)

[GIN-debug] Loaded HTML Templates (3): 
	- 
	- favicon.ico
	- index.html

[GIN-debug] GET    /static/*filepath         --> github.com/gin-gonic/gin.(*RouterGroup).createStaticHandler.func1 (3 handlers)
[GIN-debug] HEAD   /static/*filepath         --> github.com/gin-gonic/gin.(*RouterGroup).createStaticHandler.func1 (3 handlers)
[GIN-debug] GET    /                         --> go_todo/controller.IndexHandler (3 handlers)
[GIN-debug] POST   /v1/todo                  --> go_todo/controller.CreateTodoHandler (3 handlers)
[GIN-debug] GET    /v1/todo                  --> go_todo/controller.GetTodoListController (3 handlers)
[GIN-debug] PUT    /v1/todo/:id              --> go_todo/controller.UpdateTodoController (3 handlers)
[GIN-debug] DELETE /v1/todo/:id              --> go_todo/controller.DeleteTodoController (3 handlers)
[GIN-debug] Environment variable PORT is undefined. Using port :8080 by default
[GIN-debug] Listening and serving HTTP on localhost:8080
[GIN] 2020/03/21 - 23:54:40 | 200 |    4.272671ms |       127.0.0.1 | GET      "/"
[GIN] 2020/03/21 - 23:54:40 | 304 |     265.094µs |       127.0.0.1 | GET      "/static/js/app.007f9690.js"
[GIN] 2020/03/21 - 23:54:41 | 200 |    1.612225ms |       127.0.0.1 | GET      "/v1/todo"
[GIN] 2020/03/21 - 23:54:42 | 200 |   15.484888ms |       127.0.0.1 | DELETE   "/v1/todo/7"
[GIN] 2020/03/21 - 23:54:44 | 200 |    9.329674ms |       127.0.0.1 | PUT      "/v1/todo/8"
[GIN] 2020/03/21 - 23:54:45 | 200 |     355.351µs |       127.0.0.1 | GET      "/"
[GIN] 2020/03/21 - 23:54:45 | 200 |    1.234355ms |       127.0.0.1 | GET      "/v1/todo"
[GIN] 2020/03/21 - 23:54:45 | 200 |      358.64µs |       127.0.0.1 | GET      "/"
[GIN] 2020/03/21 - 23:54:45 | 200 |   11.159539ms |       127.0.0.1 | GET      "/v1/todo"
[GIN] 2020/03/21 - 23:54:47 | 200 |   13.730053ms |       127.0.0.1 | DELETE   "/v1/todo/8"
[GIN] 2020/03/21 - 23:54:47 | 200 |    8.232487ms |       127.0.0.1 | DELETE   "/v1/todo/9"
[GIN] 2020/03/21 - 23:54:48 | 200 |     483.045µs |       127.0.0.1 | GET      "/"
[GIN] 2020/03/21 - 23:54:49 | 200 |    1.649549ms |       127.0.0.1 | GET      "/v1/todo"
[GIN-debug] [WARNING] Creating an Engine instance with the Logger and Recovery middleware already attached.

[GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
 - using env:	export GIN_MODE=release
 - using code:	gin.SetMode(gin.ReleaseMode)

[GIN-debug] Loaded HTML Templates (3): 
	- 
	- favicon.ico
	- index.html

[GIN-debug] GET    /static/*filepath         --> github.com/gin-gonic/gin.(*RouterGroup).createStaticHandler.func1 (3 handlers)
[GIN-debug] HEAD   /static/*filepath         --> github.com/gin-gonic/gin.(*RouterGroup).createStaticHandler.func1 (3 handlers)
[GIN-debug] GET    /                         --> go_todo/controller.IndexHandler (3 handlers)
[GIN-debug] POST   /v1/todo                  --> go_todo/controller.CreateTodoHandler (3 handlers)
[GIN-debug] GET    /v1/todo                  --> go_todo/controller.GetTodoListController (3 handlers)
[GIN-debug] PUT    /v1/todo/:id              --> go_todo/controller.UpdateTodoController (3 handlers)
[GIN-debug] DELETE /v1/todo/:id              --> go_todo/controller.DeleteTodoController (3 handlers)
[GIN-debug] Environment variable PORT is undefined. Using port :8080 by default
[GIN-debug] Listening and serving HTTP on localhost:8080
[GIN] 2020/03/21 - 23:55:12 | 200 |       429.9µs |       127.0.0.1 | GET      "/"
[GIN] 2020/03/21 - 23:55:12 | 200 |    1.434125ms |       127.0.0.1 | GET      "/v1/todo"
[GIN] 2020/03/21 - 23:55:17 | 200 |    9.450225ms |       127.0.0.1 | POST     "/v1/todo"
[GIN] 2020/03/21 - 23:55:17 | 200 |    1.067368ms |       127.0.0.1 | GET      "/v1/todo"
[GIN] 2020/03/21 - 23:55:20 | 200 |   10.002092ms |       127.0.0.1 | POST     "/v1/todo"
[GIN] 2020/03/21 - 23:55:20 | 200 |    1.292865ms |       127.0.0.1 | GET      "/v1/todo"
[GIN] 2020/03/21 - 23:55:31 | 200 |    4.469636ms |       127.0.0.1 | POST     "/v1/todo"
[GIN] 2020/03/21 - 23:55:31 | 200 |    1.481128ms |       127.0.0.1 | GET      "/v1/todo"
[GIN] 2020/03/21 - 23:55:33 | 200 |    7.823372ms |       127.0.0.1 | PUT      "/v1/todo/12"

```