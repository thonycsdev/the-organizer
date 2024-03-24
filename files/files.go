package files

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

func GetFilesFromDirectory(directory *string) []fs.DirEntry {
	files, _ := os.ReadDir(*directory)
	files = getOnlyFilesNotDirectories(&files)
	err := checkIfThereIsAnyFileToBeOrganized(&files)
	if err != nil {
	panic(err)	
	}
	return files
}

func getOnlyFilesNotDirectories(files *[]fs.DirEntry) []fs.DirEntry {
	var result []os.DirEntry
	for _, file := range *files {
		if strings.Contains(file.Name(), ".") {
			result = append(result, file)
		}
	}
	return result
}

func checkIfThereIsAnyFileToBeOrganized(files *[]fs.DirEntry) error {
	if len(*files) == 0 {
		response := fmt.Sprint("Cannot find files to move. There is any file in this directory?")
		err := errors.New(response)
		return err
	}
	return nil
}

func CreateFolderNamesBasedOnFilesExtensions(files *[]fs.DirEntry) []string {
	var foldersNames []string
	for _, file := range *files {
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
	return foldersNames
}

func RemoveDuplicate(array []string) []string {
	m := make(map[string]string)
	for _, x := range array {
		m[x] = x
	}
	var ClearedArr []string
	for x := range m {
		ClearedArr = append(ClearedArr, x)
	}
	return ClearedArr
}

func checkFileExtension(fileName *string) string {
	result := filepath.Ext(*fileName)
	return result
}
func MoveFilesToDestination(files *[]os.DirEntry, fileNamesMap map[string]string, userCurrentLocation *string) {
	for _, file := range *files {
		name := file.Name()
		s := checkFileExtension(&name)
		goToFolder, exists := fileNamesMap[s[1:]]
		if exists {
			source := fmt.Sprint(*userCurrentLocation, "/", name)
			destination := fmt.Sprint(*userCurrentLocation, "/", goToFolder, "/", name)
			moveFile(&source, &destination)
		}
	}
}
func moveFile(sourcePath, destPath *string) error {
	file, err := os.Open(*sourcePath)
	if err != nil {
		return err
	}
    defer file.Close()

	read, _ := os.ReadFile(file.Name())
	erroCreate := os.WriteFile(*destPath, read, 0755)
	if erroCreate != nil {
		return erroCreate
	}

	err = os.Remove(*sourcePath)
	if err != nil {
		return err
	}
	fmt.Println("De: ", *sourcePath)
	fmt.Println("Para: ", *destPath)
	return nil
}
