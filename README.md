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
$ curl -v --no-buffer --header "Accept: text/event-stream" localhost:8080/api/v1/fibonacci/5/stream
* Host localhost:8080 was resolved.
* IPv6: ::1
* IPv4: 127.0.0.1
*   Trying [::1]:8080...
* Established connection to localhost (::1 port 8080) from ::1 port 43292
* using HTTP/1.x
> GET /api/v1/fibonacci/5/stream HTTP/1.1
> Host: localhost:8080
> User-Agent: curl/8.18.0
> Accept: text/event-stream
>
* Request completely sent off
< HTTP/1.1 200 OK
< Cache-Control: no-cache
< Connection: keep-alive
< Content-Type: text/event-stream
< X-Accel-Buffering: no
< Date: Mon, 02 Mar 2026 00:50:36 GMT
< Transfer-Encoding: chunked
<
{"id":0,"event":"fibonacci","data":{"ordinal":1,"value":"0"}}

{"id":1,"event":"fibonacci","data":{"ordinal":2,"value":"1"}}

{"id":2,"event":"fibonacci","data":{"ordinal":3,"value":"1"}}

{"id":3,"event":"fibonacci","data":{"ordinal":4,"value":"2"}}

{"id":4,"event":"fibonacci","data":{"ordinal":5,"value":"3"}}

* Connection #0 to host localhost:8080 left intact
```

Swagger API docs can be accessed at http://localhost:8080/swagger/index.html

![Swagger preview](./static/img/swagger.webp)

Prometheus-style metrics are exposed at `/metrics`. This includes Go and Gin (API) metrics.

```plain
$ curl -s localhost:9091/metrics
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

- Configuration is possible both through a config file, and environment variables.
- Defaults are available for every configuration item.
- The config file takes precedence over environment variables, whenever the same config item is defined through both.
- The config file location is non-configurable at the moment, and evaluates to the `$PWD` of the process.

### Config file

See [./fibo.yaml](./fibo.yaml) for reference.

Example:

```yaml
api:
  addr: ":8080"

docs:
  enabled: true

logging:
  level: "info"

metrics:
  enabled: true
  addr: ":9091"
  path: "/metrics"

debug:
  enabled: true
```

See the [Zap documentation](https://pkg.go.dev/go.uber.org/zap#AtomicLevel.UnmarshalText) for a reference of Zap log level string identifiers.

### Environment variables

| Name                 | Description                                  | Type   | Default      |
| -------------------- | -------------------------------------------- | ------ | ------------ |
| FIBO_API_ADDR        | REST API bind address                        | string | `":8080"`    |
| FIBO_DOCS_ENABLED    | Enables the REST API docs server (Swagger)   | bool   | `true`       |
| FIBO_METRICS_ENABLED | Enables the Prometheus metrics server        | bool   | `true`       |
| FIBO_METRICS_ADDR    | Metrics server bind address                  | string | `":9091"`    |
| FIBO_METRICS_PATH    | Metrics server base URL                      | string | `"/metrics"` |
| FIBO_LOGGING_LEVEL   | Log level                                    | string | "info"       |
| FIBO_DEBUG           | Enables debug mode (Gin, Zap), starts pprof. | bool   | `false`      |

## Development

See [DEVELOPMENT.md](./DEVELOPMENT.md)
