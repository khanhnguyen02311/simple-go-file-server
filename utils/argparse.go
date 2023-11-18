package utils

import (
	"errors"
	"fmt"
	"github.com/akamensky/argparse"
	"os"
)

type AppConfig struct {
	Port             *string
	Type             *string
	UploadAuth       *string
	DownloadAuth     *string
	AuthEndpoint     *string
	AllowedMIMETypes *string
	MaxFileSize      *int
}

var ConfigArgs *AppConfig

func ParseArgs() error {
	// Create new parser object
	parser := argparse.NewParser("Go Simple File Server", "Simple file server with upload and download capabilities, used for personal projects")
	// Create flags
	argPort := parser.String("p", "port", &argparse.Options{
		Required: false,
		Default:  "1323",
		Help:     "Port to run the server on"},
	)
	argType := parser.Selector("t", "type", []string{"local", "s3"}, &argparse.Options{
		Required: false,
		Default:  "local",
		Help:     "Server storage type"},
	)
	argUploadAuth := parser.Selector("u", "upload-auth", []string{"true", "false"}, &argparse.Options{
		Required: false,
		Default:  "false",
		Help:     "Request upload validation with Bearer token"},
	)
	argDownloadAuth := parser.Selector("d", "download-auth", []string{"true", "false"}, &argparse.Options{
		Required: false,
		Default:  "false",
		Help:     "Request download validation with Bearer token"},
	)
	argAuthEndpoint := parser.String("a", "auth-endpoint", &argparse.Options{
		Required: false,
		Default:  "",
		Help:     "Authentication endpoint to validate tokens (if needed)"},
	)
	argAllowedMIMEs := parser.String("l", "allowed-list", &argparse.Options{
		Required: false,
		Default:  "*",
		Help:     "Comma separated list of allowed MIME types. Example: 'image/png,image/jpeg,video/mp4'"},
	)
	argMaxFileSize := parser.Int("m", "max-file-size", &argparse.Options{
		Required: false,
		Default:  0,
		Help:     "Maximum file size in MB. Default to 0 (unlimited)"},
	)
	// Parse input
	err := parser.Parse(os.Args)
	if err != nil {
		// In case of error print error and print usage
		// This can also be done by passing -h or --help flags
		fmt.Print(parser.Usage(err))
		return err
	}
	if (*argUploadAuth == "true" || *argDownloadAuth == "true") && *argAuthEndpoint == "" {
		return errors.New("auth-endpoint must be provided when download-auth or upload-auth is enabled")
	}
	// Finally print the collected string
	fmt.Println(
		"Runtime configs: \n- Port:", *argPort,
		"\n- Type:", *argType,
		"\n- Upload Authentication:", *argUploadAuth,
		"\n- Download Authentication:", *argDownloadAuth,
		"\n- Authentication Endpoint:", *argAuthEndpoint,
		"\n- Allowed MIME types: [", *argAllowedMIMEs, "]",
		"\n- Max file size in MB (0 is unlimited):", *argMaxFileSize)

	ConfigArgs = &AppConfig{
		Port:             argPort,
		Type:             argType,
		UploadAuth:       argUploadAuth,
		DownloadAuth:     argDownloadAuth,
		AuthEndpoint:     argAuthEndpoint,
		AllowedMIMETypes: argAllowedMIMEs,
		MaxFileSize:      argMaxFileSize,
	}
	return nil
}
