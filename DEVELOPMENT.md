# Developing Fibo

This guide assumes you're using a Linux system.

## Quick start

```sh
go run .
```

The web server will bind to `http://localhost:8080/`.

To query a Fibonacci sequence:

```sh
curl -s --no-buffer localhost:8080/api/v1/fibonacci/10
# Outputs:
# [0,1,1,2,3,5,8,13,21,34]
```

## Building

To be defined

## Development

Make sure you have the following tools installed:

- [Golang](https://go.dev/dl/)
- [golangci-lint](https://golangci-lint.run/)
- [pre-commit](https://pre-commit.com/)
- [Docker](https://www.docker.com/)
- [swaggo/swag](https://github.com/swaggo/swag/)

To setup [pre-commit](https://pre-commit.com/), run:

```sh
# Repository root
pre-commit install
```

(Re)build the Swagger documentation with:

```sh
swag init
```

Start the API in live-reload mode:

```sh
air
```
