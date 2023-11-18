FROM golang:1.19

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o bin/server .

CMD ./bin/server \
    --port 8000 \
    --type local \
    --upload-auth false \
    --download-auth false \
    --allowed-list image/png,image/jpeg,image/jpg,image/gif,image/webp \
    --max-size 10