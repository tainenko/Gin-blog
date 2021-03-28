# Gin-blog [![rcard](https://goreportcard.com/badge/github.com/tainenko/gin-blog)](https://goreportcard.com/report/github.com/tainenko/gin-blog) [![License](http://img.shields.io/badge/license-mit-blue.svg?style=flat-square)](https://raw.githubusercontent.com/tainenko/gin-blog/master/LICENSE)

A blog back-end API developed based on Gin

## Installation

```
$ git clone https://github.com/tainenko/Gin-blog.git
```

## How to run

You can run this project with docker or docker-compose.yml

```
$ docker build -t gin-blog . 
$ docker run --link mysql:mysql --link redis:redis -p 8000:8000 gin-blog-docker
$ docker run --name mysql -p 3306:3306 -e MYSQL_ROOT_PASSWORD=rootroot -v ~/data/docker-mysql:/var/lib/mysql -d mysql
docker run --name redis  -p 6379:6379 -d redis
```

Or execute on the local environment that has been set up

### Required
- Mysql
- Redis

### Configuration
You may need to modify the configuration: `conf/app.ini`
```
[database]
Type = mysql
User = root
Password = 
Host = 127.0.0.1:3306
Name = blog
TablePrefix = blog_

[redis]
Host = 127.0.0.1:6379
Password =
MaxIdle = 30
MaxActive = 30
IdleTimeout = 200
...
```

### Run

```
$ cd $GOPATH/src/gin-blog
$ go run main.go 
```

Project information and existing API

```
[GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
 - using env:	export GIN_MODE=release
 - using code:	gin.SetMode(gin.ReleaseMode)

[GIN-debug] GET    /api/v1/tags              --> github.com/tainenko/gin-blog/routers/api/v1.GetTags (4 handlers)
[GIN-debug] POST   /api/v1/tags              --> github.com/tainenko/gin-blog/routers/api/v1.AddTag (4 handlers)
[GIN-debug] PUT    /api/v1/tags/:id          --> github.com/tainenko/gin-blog/routers/api/v1.EditTag (4 handlers)
[GIN-debug] DELETE /api/v1/tags/:id          --> github.com/tainenko/gin-blog/routers/api/v1.DeleteTag (4 handlers)
[GIN-debug] GET    /api/v1/articles          --> github.com/tainenko/gin-blog/routers/api/v1.GetArticles (4 handlers)
[GIN-debug] GET    /api/v1/articles/:id      --> github.com/tainenko/gin-blog/routers/api/v1.GetArticle (4 handlers)
[GIN-debug] POST   /api/v1/articles          --> github.com/tainenko/gin-blog/routers/api/v1.AddArticle (4 handlers)
[GIN-debug] PUT    /api/v1/articles/:id      --> github.com/tainenko/gin-blog/routers/api/v1.EditArticle (4 handlers)
[GIN-debug] DELETE /api/v1/articles/:id      --> github.com/tainenko/gin-blog/routers/api/v1.DeleteArticle (4 handlers)

Listening port is 8000
Actual pid is 4393
```

## Features

- RESTful API
- Gorm
- logging
- Jwt-go
- Gin
- App configurable
- Cron
- Redis