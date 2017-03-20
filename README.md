# TURL
TURL is a tiny url service written in [go](https://golang.org/)

## Introduction
* use [redis](http://redis.io) for backend storage.
* use [gin](https://github.com/gin-gonic/gin) for web service.
* use [hashids](http://www.hashids.org) for id hashing.

## Setup
```
go get github.com/wangzn/turl
```

## Example Usage

### Start server

```
> ./server
```

### Add new url

```
> curl -v -X POST localhost:8080/new -Furl=https://www.google.com/
```

### Get url from a tiny key

```
> curl -v localhost:8080/MGOQen
```


