# Website
My personal website

For sharing projects I've worked on and neat Canvas animations.

## Tech

Frontend: [Vue](https://vuejs.org/)

Backend: [Go](https://golang.org/)

Machine: [Raspberry Pi 3 Model B](https://www.raspberrypi.org/products/raspberry-pi-3-model-b/)

## Compiling

For Raspberry Pi:

```shell
env GOOS=linux GOARCH=arm GOARM=5 go build -o server.exe server/*
```

