package main

import (
	"io/ioutil"
	"strings"
	"os"
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

func ArrangeInDateDirs(files []FileInfo, root string, cleanupDirName string) {
	for _, file := range files {
		dateDir := strings.Join([]string { root, cleanupDirName, file.LastModified }, "/")

		if _, err := os.Stat(dateDir); os.IsNotExist(err) {
			os.MkdirAll(dateDir, 0755)
		}

		err := os.Rename(strings.Join([]string {root, file.Name}, "/"), strings.Join([]string {dateDir, file.Name}, "/"))

		if err != nil {
			panic(err)
		}
	}
}

func main() {
	if len(os.Args) == 1 {
		fmt.Println("Must provide root directory.")
		return
	}

	root := os.Args[1]
	cleanupDir := "cleanup"

	if len(os.Args) > 2 && os.Args[2] != "" {
		cleanupDir = os.Args[2]
	}

	files, err := RetrieveFiles(root)
	
	if err != nil {
			panic(err)
	}

	ArrangeInDateDirs(files, root, cleanupDir)
}