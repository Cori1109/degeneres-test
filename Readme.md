## Degeneres Test

Description for Degeneres Test goes here.

### Installation

```sh
go get -u -v github.com/rms1000watt/degeneres-test
```

### Usage

```sh
go run main.go
```

### Deploy

```sh
GOOS=linux go build
docker build --rm --no-cache -t docker.io/rms1000watt/degeneres-test:v0.1.0 .
docker push docker.io/rms1000watt/degeneres-test:v0.1.0
docker run -itd -p 443:443 docker.io/rms1000watt/degeneres-test:v0.1.0 [ARGS...]
```


**Built using [Degeneres](https://www.github.com/rms1000watt/degeneres)**
