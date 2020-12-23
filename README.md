# Simple GO HTTP server

This is a simple example showing how-to create a simple GO web app and then package it inside an "OpenShift ready" container;

## project setup

```
go mod init github.com/tommaso-borgato/go-server
```

## build

```
go build -o app
```

## run

```
./app
```

## test

```
curl localhost:8080
```

## containerize

```
go build -o app
sudo podman login registry.redhat.io
sudo podman pull registry.redhat.io/ubi8/ubi-minimal
sudo podman build --tag go-server -f ./Dockerfile
sudo podman run -p 8080:80 --name go-server --rm go-server
curl localhost:8080
```

### Verify the size of your image

```
sudo buildah images
```
