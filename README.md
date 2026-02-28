# fibo

## Getting Started

```sh
go run .
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
$ curl -s localhost:9090/metrics
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

## Development

See [DEVELOPMENT.md](./DEVELOPMENT.md)
