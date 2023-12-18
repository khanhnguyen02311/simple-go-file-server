FROM golang:1.21 AS builder

RUN useradd -u 1001 -m nonroot

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o bin/server

FROM golang:1.21-alpine AS runner

WORKDIR /app

COPY --from=builder /etc/passwd /etc/passwd

COPY --from=builder /app/bin/server ./bin/server

COPY entrypoint.sh ./

RUN chmod +x ./entrypoint.sh

STOPSIGNAL SIGQUIT

USER 1001

ENTRYPOINT [ "/app/entrypoint.sh" ]