package main

import (
	"io/ioutil"
	"strings"
	"fmt"
)

type FileInfo struct {
	Name string
	LastModified string
}

func RetrieveFiles(root string) ([]FileInfo, error) {
	var files []FileInfo
	
	fileInfo, err := ioutil.ReadDir(root)
	if err != nil {
		return files, err
	}

	for _, file := range fileInfo {
		if !strings.HasPrefix(file.Name(), ".") && !strings.HasPrefix(file.Name(), "cleanup") {
			files = append(files, FileInfo{ file.Name(), file.ModTime().Format("02-01-2006")})
		}
	}
	return files, nil
}

func main() {
	root := "/Users/lukasz/Desktop"
	files, err := RetrieveFiles(root)
	
	if err != nil {
			panic(err)
	}
	fmt.Print(files)
}