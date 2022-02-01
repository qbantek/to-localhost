# Builder
FROM golang:1.17-alpine3.14 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 go build -ldflags "-s -w" -o bin/to-localhost .

# Runner
FROM alpine:latest

ARG GIN_MODE=release
ENV GIN_MODE=$GIN_MODE
ARG PORT=5000
ENV PORT=$PORT

EXPOSE $PORT

WORKDIR /app
COPY --from=builder /app/bin/to-localhost bin/to-localhost
COPY --from=builder /app/templates/*.tmpl.html templates/
COPY --from=builder /app/static/main.min.css /app/static/favicon.ico static/

RUN chmod +x /app/bin/to-localhost

CMD ["/app/bin/to-localhost"]
