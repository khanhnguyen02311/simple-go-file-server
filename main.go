package main

import (
	"FileServer/utils"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log"
	"net/http"
	"os"
	"strings"
)

//func failOnError(err error, msg string) {
//	if err != nil {
//		log.Panicf("%s: %s", msg, err)
//	}
//}

func pong(c echo.Context) error {
	return c.String(http.StatusOK, "PONG!")
}

func getFile(c echo.Context, storage *utils.Storage) error {
	fmt.Printf("Getting file %s\n", c.Param("filename"))
	return c.File("files/public/" + c.Param("filename"))
}

func uploadFile(c echo.Context, storage *utils.Storage) error {
	accessToken := ""
	if storage.UploadAuth == true {
		authorizationHeader := strings.Split(c.Request().Header.Get("Authorization"), " ")
		if authorizationHeader[0] != "Bearer" {
			return c.String(http.StatusBadRequest, "Invalid token")
		}
		accessToken = authorizationHeader[1]
	}
	//fmt.Printf("Token from header: %s\n", accessToken)
	file, err := c.FormFile("file")
	if err != nil {
		return c.String(http.StatusBadRequest, "Invalid file: "+err.Error())
	}
	savedFilename, err := storage.UploadFile(file, accessToken)
	if err != nil {
		return c.String(http.StatusBadRequest, "Upload error: "+err.Error())
	}
	return c.String(http.StatusOK, savedFilename)
}

func main() {
	err := os.MkdirAll("files/public", 0777)
	err = os.MkdirAll("files/static", 0777)
	if err != nil {
		fmt.Printf("Creating directory unsuccessful: %s\n", err.Error())
	}
	err = utils.ParseArgs()
	if err != nil {
		log.Fatal(err)
		return
	}
	server := echo.New()
	storage := &utils.Storage{}
	storage.Init()

	// Middleware
	server.Use(middleware.Logger())
	server.Use(middleware.Recover())

	// Routes
	server.GET("/fs/ping", pong)
	server.GET("/fs/get/:filename",
		func(c echo.Context) error { return getFile(c, storage) })
	server.POST("/fs/upload",
		func(c echo.Context) error { return uploadFile(c, storage) })

	// Static
	server.Static("/fs/static", "./files/static")
	server.Logger.Fatal(server.Start(":" + *utils.ConfigArgs.Port))
}
