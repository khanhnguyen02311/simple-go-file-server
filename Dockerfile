FROM golang:1.21 AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o bin/server

FROM golang:1.21-alpine AS runner

WORKDIR /app

COPY --from=builder /app/bin/server ./bin/server

STOPSIGNAL SIGQUIT

CMD ./bin/server \
    --port 1323 \
    --type local \
    --upload-auth false \
    --download-auth false \
    --allowed-list image/png,image/jpeg,image/jpg,image/gif,image/webp \
    --max-file-size 10