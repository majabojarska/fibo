# fibo

[![Lint](https://github.com/majabojarska/fibo/actions/workflows/golangci-lint.yml/badge.svg)](https://github.com/majabojarska/fibo/actions/workflows/golangci-lint.yml)
[![Test](https://github.com/majabojarska/fibo/actions/workflows/test.yaml/badge.svg)](https://github.com/majabojarska/fibo/actions/workflows/test.yaml)
[![Build](https://github.com/majabojarska/fibo/actions/workflows/build.yaml/badge.svg)](https://github.com/majabojarska/fibo/actions/workflows/build.yaml)
[![Release Drafter](https://github.com/majabojarska/fibo/actions/workflows/release-drafter.yaml/badge.svg)](https://github.com/majabojarska/fibo/actions/workflows/release-drafter.yaml)

[![Docker Image Version](https://img.shields.io/docker/v/majabojarska/fibo)](https://hub.docker.com/r/majabojarska/fibo/tags)
[![Docker Image Size](https://img.shields.io/docker/image-size/majabojarska/fibo)](https://hub.docker.com/r/majabojarska/fibo/tags)
[![Docker Pulls](https://img.shields.io/docker/pulls/majabojarska/fibo)](https://hub.docker.com/r/majabojarska/fibo/tags)

## Quick Start

```sh
docker compose up
```

The web server will bind to `http://localhost:8080/`.

To query a Fibonacci sequence:

```sh
$ curl -s --no-buffer localhost:8080/api/v1/fibonacci/10
# Outputs:
# [0,1,1,2,3,5,8,13,21,34]
```

Swagger API docs can be accessed at http://localhost:8080/swagger/index.html

![Swagger preview](./static/img/swagger.webp)

Prometheus-style metrics are exposed at `/metrics`. This includes Go and Gin (API) metrics.

```plain
$ curl -s localhost:8081/metrics
# HELP gin_request_duration_seconds The HTTP request latencies in seconds.
# TYPE gin_request_duration_seconds histogram
gin_request_duration_seconds_bucket{code="200",method="GET",url="/api/v1/fibonacci/:count",le="0.005"} 169
gin_request_duration_seconds_bucket{code="200",method="GET",url="/api/v1/fibonacci/:count",le="0.01"} 169
gin_request_duration_seconds_bucket{code="200",method="GET",url="/api/v1/fibonacci/:count",le="0.025"} 169
gin_request_duration_seconds_bucket{code="200",method="GET",url="/api/v1/fibonacci/:count",le="0.05"} 169
gin_request_duration_seconds_bucket{code="200",method="GET",url="/api/v1/fibonacci/:count",le="0.1"} 169
gin_request_duration_seconds_bucket{code="200",method="GET",url="/api/v1/fibonacci/:count",le="0.25"} 169
gin_request_duration_seconds_bucket{code="200",method="GET",url="/api/v1/fibonacci/:count",le="0.5"} 169
gin_request_duration_seconds_bucket{code="200",method="GET",url="/api/v1/fibonacci/:count",le="1"} 169
gin_request_duration_seconds_bucket{code="200",method="GET",url="/api/v1/fibonacci/:count",le="2.5"} 169
gin_request_duration_seconds_bucket{code="200",method="GET",url="/api/v1/fibonacci/:count",le="5"} 169
gin_request_duration_seconds_bucket{code="200",method="GET",url="/api/v1/fibonacci/:count",le="10"} 169
gin_request_duration_seconds_bucket{code="200",method="GET",url="/api/v1/fibonacci/:count",le="+Inf"} 169
gin_request_duration_seconds_sum{code="200",method="GET",url="/api/v1/fibonacci/:count"} 0.013545205999999997
gin_request_duration_seconds_count{code="200",method="GET",url="/api/v1/fibonacci/:count"} 169
```

The Prometheus UI is available at `localhost:9090`.

![Prometheus UI preview](./static/img/prometheus.webp)

## Configuration

This API uses [Viper](https://github.com/spf13/viper) for configuration management.

At the moment, configuration is possible through environment variables:

| Name                 | Description                                  | Type   | Default      |
| -------------------- | -------------------------------------------- | ------ | ------------ |
| FIBO_API_ADDR        | REST API bind address                        | string | `":8080"`    |
| FIBO_DOCS_ENABLED    | Enables the REST API docs server (Swagger)   | bool   | `true`       |
| FIBO_METRICS_ENABLED | Enables the Prometheus metrics server        | bool   | `true`       |
| FIBO_METRICS_ADDR    | Metrics server bind address                  | string | `":8081"`    |
| FIBO_METRICS_PATH    | Metrics server base URL                      | string | `"/metrics"` |
| FIBO_DEBUG           | Enables debug mode (Gin, Zap), starts pprof. | bool   | `false`      |

## Development

See [DEVELOPMENT.md](./DEVELOPMENT.md)
