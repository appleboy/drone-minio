# drone-minio

[![GoDoc](https://godoc.org/github.com/appleboy/drone-minio?status.svg)](https://godoc.org/github.com/appleboy/drone-minio)
[![Build Status](http://drone.wu-boy.com/api/badges/appleboy/drone-minio/status.svg)](http://drone.wu-boy.com/appleboy/drone-minio)
[![codecov](https://codecov.io/gh/appleboy/drone-minio/branch/master/graph/badge.svg)](https://codecov.io/gh/appleboy/drone-minio)
[![Go Report Card](https://goreportcard.com/badge/github.com/appleboy/drone-minio)](https://goreportcard.com/report/github.com/appleboy/drone-minio)
[![Docker Pulls](https://img.shields.io/docker/pulls/appleboy/drone-minio.svg)](https://hub.docker.com/r/appleboy/drone-minio/)
[![](https://images.microbadger.com/badges/image/appleboy/drone-minio.svg)](https://microbadger.com/images/appleboy/drone-minio "Get your own image badge on microbadger.com")
[![Build status](https://ci.appveyor.com/api/projects/status/pmkfbnwtlf1fm45l/branch/master?svg=true)](https://ci.appveyor.com/project/appleboy/drone-minio/branch/master)

Drone plugin to upload or remove filesystems and object storage.

## Build

Build the binary with the following commands:

```
go build
```

## Docker

Build the Docker image with the following commands:

```
GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -a -tags netgo -o release/linux/amd64/drone-minio
docker build --rm -t appleboy/drone-minio .
```
