# go-simple-file-server

A simple file server with upload and download capabilities, used for personal projects. Written in Go.

### Start server:

```go mod tidy```

```go run main.go --flags-if-needed--```

### Flags:

| Flag                       | Attribute                                                                                          |
|----------------------------|----------------------------------------------------------------------------------------------------|
| ```-h / --help```          | CLI help                                                                                           |
| ```-p / --port```          | Port to run the server on (default: 1323)                                                          |
| ```-t / --type```          | Server storage type. Set it to "local" or "s3". Currently only local is supported (default: local) |
| ```-u / --upload-auth```   | Validate upload requests with Bearer token (default: false)                                        |
| ```-d / --download-auth``` | Validate download requests with Bearer token (default: false)                                      |
| ```-a / --auth-endpoint``` | Authentication endpoint to validate tokens (if needed) (default: "")                               |
| ```-l / --allowed-list```  | Comma separated list of allowed MIME types. Example: 'image/png,image/jpeg,video/mp4' (default: *) |
| ```-m / --max-file-size``` | Max file size in MB (default: 0 as unlimited)                                                      |

### Endpoints:

- ```GET /ping``` - Ping pong with server
- ```GET /get/{filename}``` - Get file by filename
- ```POST /upload``` - Upload file

### Run with Docker:

```docker build -t go-simple-file-server .```

```docker run -p 1323:1323 go-simple-file-server --name file-server```