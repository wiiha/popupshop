package server

import "github.com/wiiha/popupshop/files"

type FilesAPI struct {
	Addr        string
	Port        string
	FileService *files.FileService
}

func NewFilesAPI(addr, port string, fileService *files.FileService) *FilesAPI {
	return &FilesAPI{
		Addr:        addr,
		Port:        port,
		FileService: fileService,
	}
}
