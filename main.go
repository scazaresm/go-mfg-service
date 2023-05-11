package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/scazaresm/go-file-storage-service/client"
)

var fileStorageClient = client.NewFileStorageClient(&http.Client{}, os.Getenv("FILE_STORAGE_SERVICE_URL"))

func main() {
	router := gin.Default()
	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8085"
	}

	router.POST("v1/file", UploadFile)

	router.Run(fmt.Sprintf(":%s", port))
}

func UploadFile(c *gin.Context) {
	
}
