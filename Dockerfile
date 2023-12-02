FROM golang:1.21 AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o bin/server

FROM golang:1.21-alpine AS runner

WORKDIR /app

COPY --from=builder /app/bin/server ./bin/server

COPY entrypoint.sh ./

RUN chmod +x ./entrypoint.sh

STOPSIGNAL SIGQUIT

ENTRYPOINT [ "/app/entrypoint.sh" ]