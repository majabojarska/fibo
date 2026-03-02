# Fibo

[![Lint](https://github.com/majabojarska/fibo/actions/workflows/lint.yml/badge.svg)](https://github.com/majabojarska/fibo/actions/workflows/lint.yml)
[![Test](https://github.com/majabojarska/fibo/actions/workflows/test.yaml/badge.svg)](https://github.com/majabojarska/fibo/actions/workflows/test.yaml)
[![Build](https://github.com/majabojarska/fibo/actions/workflows/build.yaml/badge.svg)](https://github.com/majabojarska/fibo/actions/workflows/build.yaml)
[![Release Drafter](https://github.com/majabojarska/fibo/actions/workflows/release-drafter.yaml/badge.svg)](https://github.com/majabojarska/fibo/actions/workflows/release-drafter.yaml)

[![Docker Image Version](https://img.shields.io/docker/v/majabojarska/fibo)](https://hub.docker.com/r/majabojarska/fibo/tags)
[![Docker Image Size](https://img.shields.io/docker/image-size/majabojarska/fibo)](https://hub.docker.com/r/majabojarska/fibo/tags)
[![Docker Pulls](https://img.shields.io/docker/pulls/majabojarska/fibo)](https://hub.docker.com/r/majabojarska/fibo/tags)

## About

_Fibo_ is a showcase project, implementing a streaming REST API, based on [server-sent events](https://developer.mozilla.org/en-US/docs/Web/API/Server-sent_events/Using_server-sent_events).

## Live demo (hosted)

This service is currently hosted at [fibo.cloud.majabojarska.dev](https://fibo.cloud.majabojarska.dev/) (Swagger docs). Please do note that Swagger will buffer the entire stream before displaying it in the UI.

To query (stream) a Fibonacci sequence with `curl`:

```sh
curl --silent --verbose --no-buffer --header "Accept: text/event-stream" https://fibo.cloud.majabojarska.dev/api/v1/fibonacci/1000/stream
```

## Quick Start (local)

```sh
docker compose up
```

The web server will bind to `http://localhost:8080/`.

To query (stream) a Fibonacci sequence:

```sh
curl --silent --verbose --no-buffer --header "Accept: text/event-stream" localhost:8080/api/v1/fibonacci/10/stream
```

Here's a Fibonacci stream example with `api.event_delay: 200ms`, to better demonstrate the principle of operation.

[![asciicast](https://asciinema.org/a/wzkuKRlfRgfGKEvW.svg)](https://asciinema.org/a/wzkuKRlfRgfGKEvW)

Swagger API docs can be accessed at http://localhost:8080/swagger/index.html

![Swagger preview](./static/img/swagger.webp)

Prometheus-style metrics are exposed at `/metrics`. This includes Go and Gin (API) metrics.

```plain
curl localhost:9091/metrics
# HELP gin_request_duration_seconds The HTTP request latencies in seconds.
# TYPE gin_request_duration_seconds histogram
gin_request_duration_seconds_bucket{code="200",method="GET",url="/api/v1/fibonacci/100/stream",le="0.005"} 5
gin_request_duration_seconds_bucket{code="200",method="GET",url="/api/v1/fibonacci/100/stream",le="0.01"} 5
gin_request_duration_seconds_bucket{code="200",method="GET",url="/api/v1/fibonacci/100/stream",le="0.025"} 5
gin_request_duration_seconds_bucket{code="200",method="GET",url="/api/v1/fibonacci/100/stream",le="0.05"} 5
gin_request_duration_seconds_bucket{code="200",method="GET",url="/api/v1/fibonacci/100/stream",le="0.1"} 5
gin_request_duration_seconds_bucket{code="200",method="GET",url="/api/v1/fibonacci/100/stream",le="0.25"} 5
gin_request_duration_seconds_bucket{code="200",method="GET",url="/api/v1/fibonacci/100/stream",le="0.5"} 5
gin_request_duration_seconds_bucket{code="200",method="GET",url="/api/v1/fibonacci/100/stream",le="1"} 5
gin_request_duration_seconds_bucket{code="200",method="GET",url="/api/v1/fibonacci/100/stream",le="2.5"} 5
gin_request_duration_seconds_bucket{code="200",method="GET",url="/api/v1/fibonacci/100/stream",le="5"} 5
gin_request_duration_seconds_bucket{code="200",method="GET",url="/api/v1/fibonacci/100/stream",le="10"} 5
gin_request_duration_seconds_bucket{code="200",method="GET",url="/api/v1/fibonacci/100/stream",le="+Inf"} 5
gin_request_duration_seconds_sum{code="200",method="GET",url="/api/v1/fibonacci/100/stream"} 0.004046882
gin_request_duration_seconds_count{code="200",method="GET",url="/api/v1/fibonacci/100/stream"} 5
```

The Prometheus UI is available at `localhost:9090`.

![Prometheus UI preview](./static/img/prometheus.webp)

## Design & data format

I've made several intentional design choices for this API:

- Fibonacci sequence items get quite large, quickly. For reference, after the 93rd item, the value overflows an int64. That's why I've implemented sequence generation using [`big.Int`](https://pkg.go.dev/math/big#Int), which allows for arbitrarily large integers, at the expense of longer computation.
- Fibonacci sequence values are sent back as strings, not as numbers. That's because no JSON number representation can store (fit) sequence items larger than int64. That is, without corrupting the data. On the API consumer's side, the strings can be decoded back into an equivalent "big int" structure, retaining integrity end-to-end.
- This is a streaming API (except the healtcheck routes), based on [server-sent events](https://developer.mozilla.org/en-US/docs/Web/API/Server-sent_events/Using_server-sent_events). Sending the whole sequence in a single document would require the client to buffer the entire response body, before it can process the JSON document, causing it to quickly exhaust its memory. Beyond some sequence size, it becomes impossible for the client to even receive the entire message. That is, unless the client uses custom parsing and/or flushes to disk. The current implementation, simply put:
  - Each sequence item is sent as a separate event.
  - Each event is a complete JSON document, allowing for immediate processing.
  - Events are delimited with `\n\n`.

The event structure is as follows:

```json
{
  "id": 0,
  "event": "fibonacci",
  "data": {
    "ordinal": 1,
    "value": "0"
  }
}
```

- `id` tracks the number of events sent to the client (0-indexed integer)
- `event` denotes the event type.
- `data` is the message, i.e. what we actually want to transfer.
  - `ordinal` tracks the fibonacci sequence's ordinal (1-indexed integer).
  - `value` is the actual value, the big.Int serialized to a string.

You may have noticed that `data.ordinal` is always greater by 1 from `id`, and you'd be right. Perhaps I do waste some data in that regard, but it does help decouple the SSE metadata, from the actual data. If efficiency was a concern, [gRPC](https://grpc.io/) would be a wiser choice for better performance and efficency.

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
  allow_origins:
    - "http://localhost"
  event_delay: 0ms

docs:
  enabled: true

logging:
  level: "info"

metrics:
  enabled: true
  addr: ":9091"
  path: "/metrics"

debug:
  enabled: false
```

See the [Zap documentation](https://pkg.go.dev/go.uber.org/zap#AtomicLevel.UnmarshalText) for a reference of Zap log level string identifiers.

### Environment variables

| Name                     | Description                                                                                                                                                                                 | Type                                                       | Default                             |
| ------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------- | ----------------------------------- |
| `FIBO_API_ADDR`          | REST API bind address                                                                                                                                                                       | string                                                     | `":8080"`                           |
| `FIBO_API_ROOT_URL`      | URL through which the API will be externally available                                                                                                                                      | string                                                     | `"http://localhost:8080"`           |
| `FIBO_API_ALLOW_ORIGINS` | Populates the [`Access-Control-Allow-Origin`](https://developer.mozilla.org/en-US/docs/Web/HTTP/Reference/Headers/Access-Control-Allow-Origin) header. Comma-separated for multiple values. | []string                                                   | `"http://localhost,http://a.b.c.d"` |
| `FIBO_API_EVENT_DELAY`   | Time to wait between subsequent stream events                                                                                                                                               | string ([Duration](https://pkg.go.dev/time#ParseDuration)) | `"0s"`                              |
| `FIBO_DOCS_ENABLED`      | Enables the REST API docs server (Swagger)                                                                                                                                                  | bool                                                       | `true`                              |
| `FIBO_METRICS_ENABLED`   | Enables the Prometheus metrics server                                                                                                                                                       | bool                                                       | `true`                              |
| `FIBO_METRICS_ADDR`      | Metrics server bind address                                                                                                                                                                 | string                                                     | `":9091"`                           |
| `FIBO_METRICS_PATH`      | Metrics server base URL                                                                                                                                                                     | string                                                     | `"/metrics"`                        |
| `FIBO_LOGGING_LEVEL`     | Log level                                                                                                                                                                                   | string                                                     | `"info"`                            |
| `FIBO_DEBUG`             | Enables debug mode (Gin, Zap), starts pprof.                                                                                                                                                | bool                                                       | `false`                             |

## Development

See [DEVELOPMENT.md](./DEVELOPMENT.md)
