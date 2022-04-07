package server

import (
	"fmt"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func (svc *FilesAPI) ListFilesInDir(c *gin.Context) {
	files, err := svc.FileService.FilesInDir()
	if err != nil {
		c.String(http.StatusInternalServerError, "Error: %v", err)
		return
	}

	htmlString := "<h1>Files:</h1><ul>"
	for i := range files {
		f := files[i]
		htmlString += fmt.Sprintf("<li><a href='/dl/%s' target='_blank' rel='noreferrer'>%s</a></li>", f, f)
	}
	htmlString += "</ul>"
	htmlData := []byte(htmlString)
	c.Data(http.StatusOK, "text/html; charset=utf-8", htmlData)
}

func (svc *FilesAPI) DownloadFile(c *gin.Context) {
	fileName := c.Param("fileName")
	if fileName == "" {
		c.String(http.StatusBadRequest, "Error: no file name specified")
		return
	}
	if !svc.FileService.Exists(fileName) {
		c.String(http.StatusNotFound, "Error: file not found")
		return
	}
	path := filepath.Join(svc.FileService.Root, fileName)

	c.Header("Content-Description", "File Transfer")
	c.Header("Content-Transfer-Encoding", "binary")
	c.Header("Content-Disposition", fmt.Sprintf("attachment; filename=%s", fileName))
	c.Header("Content-Type", "application/octet-stream")
	// c.Header("Content-Length", fmt.Sprintf("%d", svc.FileService.Size(path)))
	c.File(path)
}
