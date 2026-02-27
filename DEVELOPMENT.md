# Developing Fibo

This guide assumes you're using a Linux system.

## Building

### Binary

```sh
go build .
# Outputs binary: ./fibo
```

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
