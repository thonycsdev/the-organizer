package directories

import (
	"fmt"
	"os"
	"path/filepath"
)

func GetApplicationCurrentDirectory() string {

	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	return filepath.Dir(ex)
}
func CreateDirectoriesByFoldersNames(names *[]string, location *string) {
	for _, name := range *names {
		filePath := fmt.Sprint(*location, "/", name[1:])
		os.Mkdir(filePath, 0755)
	}
}
func ConvertStringsToMaps(array *[]string) map[string]string {
	fileNamesMap := map[string]string{}
	for _, name := range *array {
		fileNamesMap[name[1:]] = name[1:]
	}
	return fileNamesMap
}
