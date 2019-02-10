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

## Setup for Running on boot

Note: This probably depends on the OS etc, but for Raspberry Pi this is what's worked.

Make the `website` file executable

```bash
chmod +x website
```

Copy the `website` file into `/etc/init.d`

```bash
sudo cp website /etc/init.d
```

Update system services

```bash
sudo update-rc.d website defaults
```

Then can use with

```shell
sudo service website {start|stop|status}
```

