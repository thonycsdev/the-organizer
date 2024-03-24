package main

import (
	"fmt"

	"github.com/thonycsdev/the-organizer/directories"
	"github.com/thonycsdev/the-organizer/files"
)

func main() {
	//userInformedLocation := directories.GetApplicationCurrentDirectory()
	userInformedLocation := "/home/thonycsdev/Downloads"
	fmt.Println("Go Organizer")
	fmt.Printf("-> %s <-", userInformedLocation)
	question := fmt.Sprint(` is this the right location?: 
                            (y)Yes, n(No)`)
	fmt.Print(question)
	userResponse := ""
	fmt.Scan(&userResponse)

	switch userResponse {
	case "y":
		fmt.Println("Processing...")
	default:
		return

	}
	filesFromDir := files.GetFilesFromDirectory(&userInformedLocation)
	foldersNames := files.CreateFolderNamesBasedOnFilesExtensions(&filesFromDir)
	directories.CreateDirectoriesByFoldersNames(&foldersNames, &userInformedLocation)
	fileNamesMap := directories.ConvertStringsToMaps(&foldersNames)
	files.MoveFilesToDestination(&filesFromDir, fileNamesMap, &userInformedLocation)
}
