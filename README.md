# fibo

## Development

See [DEVELOPMENT.md](./DEVELOPMENT.md)

## To-do

Roughly in order of execution:

- Build a REST API around the fibo sequence iterator.
  - Use [server-sent events](https://developer.mozilla.org/en-US/docs/Web/API/Server-sent_events/Using_server-sent_events) for subsequent fibo item streaming.
- Containerize
  - Dockerfile
  - Docker compose
  - Build & run docs in README.md.
- Implement docker image build & release workflow.
  - Triggered on release.
  - Dockerhub
  - (maybe?) Attestations
  - Re-use workflows from [bitwarden-cli-docker/.github/workflows/release.yaml](https://github.com/majabojarska/bitwarden-cli-docker/blob/main/.github/workflows/release.yaml).
- Document the API through an automatically generated OpenAPI specification, in a code-first, docs-second approach. [swaggo/swag](https://github.com/swaggo/swag) will likely be used to achieve this.
- O11y
  - Zap for logging.
  - The REST API will be instrumented with [prometheus/client_golang](https://github.com/prometheus/client_golang).
  - Traces if I have the time to do that.
  - LGTM stack or similar via docker compose.
- Helm chart.
  - Feature parity with docker compose.
