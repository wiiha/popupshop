/*
This package will handle interaction
with the file system.
*/
package files

import (
	"fmt"
	"io/fs"
	"io/ioutil"
	"os"
	"path/filepath"
)

/*
FileService is a struct that holds the root directory
for which this service will operate.
*/
type FileService struct {
	Root string
}

/*
NewFileService makes one
*/
func NewFileService(root string) *FileService {
	return &FileService{
		Root: root,
	}
}

/*
IOReadDir lists the contents of a directory.
It does not list the contents of subdirectories.
*/
func (svc *FileService) IOReadDir() ([]fs.FileInfo, error) {
	var files []fs.FileInfo
	fileInfo, err := ioutil.ReadDir(svc.Root)
	if err != nil {
		return nil, fmt.Errorf("when reading dir: %v", err)
	}
	for _, file := range fileInfo {
		files = append(files, file)
	}
	return files, nil
}

/*
FilesInDir lists only files in a directory.
It does not list files in subdirectories.
*/
func (svc *FileService) FilesInDir() ([]string, error) {
	files, err := svc.IOReadDir()
	if err != nil {
		return nil, fmt.Errorf("IOReadDir: %v", err)
	}

	fileNames := []string{}
	for i := range files {
		if !files[i].IsDir() {
			fileNames = append(fileNames, files[i].Name())
		}
	}

	return fileNames, nil
}

/*
Exists checks if a _filename_ exists
in the root directory.
*/
func (svc *FileService) Exists(filename string) bool {
	path := filepath.Join(svc.Root, filename)
	_, err := os.Stat(path)
	if err != nil {
		return false
	}
	return true
}
