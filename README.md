# go-simple-file-server

A simple file server written in Go.

Start server:
```go mod tidy```
```go run main.go `<flags>````

Flags:
```-h / --help``` - help
```-p / --port``` - Port to run the server on (default: 1323)
```-t  --type``` - Server storage type. Set it to "local" or "s3" (currently only local is supported)
```-u  --upload-auth``` - Validate upload requests with Bearer token (default: false)
```-d  --download-auth``` - Validate download requests with Bearer token (default: false)
```-a  --auth-endpoint``` - Authentication endpoint to validate tokens (if needed) (default: "")
```-l  --allowed-list``` - Comma separated list of allowed MIME types. Example: 'image/png,image/jpeg,video/mp4' (default: *)
```-m  --max-file-size``` - Max file size in MB (default: 0 as unlimited)
