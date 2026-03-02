FROM --platform=$BUILDPLATFORM golang:1.26-alpine AS build

# Docker build args
ARG BUILDPLATFORM
ARG TARGETPLATFORM

# Go build args
ARG TARGETOS
ARG TARGETARCH

WORKDIR /app

# Cache Go module downloads
COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY ./ ./

RUN CGO_ENABLED=0 GOOS=${TARGETOS} GOARCH=${TARGETARCH} go build -ldflags="-s -w" -o /app/api ./cmd/api

FROM scratch

COPY --from=build /app/api /api

EXPOSE 8080
EXPOSE 8081

ENTRYPOINT ["/api"]
