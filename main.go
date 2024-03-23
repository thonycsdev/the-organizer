package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	userInformedLocation := "/home/thonycsdev/Downloads"
	fmt.Println("Go Organizer")
	files, _ := os.ReadDir(userInformedLocation)
	files = filterFilesFromDirectories(files)
	if len(files) == 0 {
		fmt.Println("Cannot find files to move. There is any file in this directory?")
		return
	}
	var foldersNames []string
	for _, file := range files {
		fileName := file.Name()
		r := checkFileExtension(&fileName)
		if r == "" {
			continue
		}
		if strings.Contains(fileName, ".") {
			foldersNames = append(foldersNames, r)
		}
	}
	foldersNames = RemoveDuplicate(foldersNames)

	CreateFolderByFolderNames(&foldersNames, &userInformedLocation)
	fileNamesMap := PassToMap(foldersNames)
    moveFilesToDestination(files, fileNamesMap, userInformedLocation )
}

func filterFilesFromDirectories(files []os.DirEntry) []os.DirEntry {

	var result []os.DirEntry
	for _, file := range files {
		if strings.Contains(file.Name(), ".") {
			result = append(result, file)
		}
	}
	return result
}
func moveFilesToDestination(files []os.DirEntry, fileNamesMap map[string]string, userCurrentLocation string) {
	for _, file := range files {
		name := file.Name()
		s := checkFileExtension(&name)
		goToFolder, exists := fileNamesMap[s[1:]]
		if exists {
			source := fmt.Sprint(userCurrentLocation, "/", name)
			destination := fmt.Sprint(userCurrentLocation, "/", goToFolder, "/", name)
			moveFile(source, destination)
		}
	}
}
func moveFile(sourcePath, destPath string) error {
	_, err := os.Open(sourcePath)
	if err != nil {
		return err
	}

	_, erroCreate := os.Create(destPath)
	if erroCreate != nil {
		return erroCreate
	}

	err = os.Remove(sourcePath)
	if err != nil {
		return err
	}
	logString := fmt.Sprint("De: ", sourcePath, " Para ", sourcePath)
	log.Print(logString)
	return nil
}
func MoveFile(source, destination string) (err error) {
	return nil
}

func CreateFolderByFolderNames(names *[]string, location *string) {
	for _, name := range *names {
		filePath := fmt.Sprint(*location, "/", name[1:])
		os.Mkdir(filePath, 0755)
	}
}

func RemoveDuplicate(array []string) []string {
	m := make(map[string]string)
	for _, x := range array {
		m[x] = x
	}
	var ClearedArr []string
	for x, _ := range m {
		ClearedArr = append(ClearedArr, x)
	}
	return ClearedArr
}

func PassToMap(array []string) map[string]string {
	fileNamesMap := map[string]string{}
	for _, name := range array {
		fileNamesMap[name[1:]] = name[1:]
	}
	return fileNamesMap
}

func checkFileExtension(fileName *string) string {
	result := filepath.Ext(*fileName)
	return result
}
func filterSlice[T any](s *[]T, testFunction func(T) bool) []T {
	result := []T{}
	for _, val := range *s {
		if testFunction(val) {
			result = append(result, val)
		}
	}
	return result
}
