## To-do

Roughly in order of execution:

- Use strings for fibo sequence. JSON Number breaks integrity.
- Use JSON-error style responses ("message", "error", "data")
- Set up Viper for config mgmt
- Fuzzing tests? https://go.dev/doc/tutorial/fuzz
- Containerize
  - Dockerfile
  - Docker compose
  - Build & run docs in README.md.
- Implement docker image build & release workflow.
  - Triggered on release.
  - Dockerhub
  - (maybe?) Attestations
  - Re-use workflows from [bitwarden-cli-docker/.github/workflows/release.yaml](https://github.com/majabojarska/bitwarden-cli-docker/blob/main/.github/workflows/release.yaml).
- O11y
  - Traces if I have the time to do that.
    - https://github.dev/gin-gonic/examples/blob/ec3c1774716d51e4100ae4f995957048d2c86030/otel/main.go
  - LGTM stack or similar via docker compose.
- Helm chart.
  - Feature parity with docker compose.
- Add optional pprof

