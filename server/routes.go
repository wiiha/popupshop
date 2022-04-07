package server

import (
	"github.com/gin-gonic/gin"
)

func (svc *FilesAPI) Run() error {
	// Init router
	r := gin.Default()

	/*
		Register routes.
		Each route needs a handler function.
		These are defined in the package file
		handlers.go.
	*/

	r.GET("/", svc.ListFilesInDir)

	r.GET("/dl/:fileName", svc.DownloadFile)

	addrAndPort := svc.Addr + ":" + svc.Port
	err := r.Run(addrAndPort)
	return err
}
