FROM golang:1.26-alpine AS build

WORKDIR /app

# Cache Go module downloads
COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY ./ ./

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o /app/server .

FROM scratch

COPY --from=build /app/server /server

EXPOSE 8080
EXPOSE 9090

ENTRYPOINT ["/server"]
