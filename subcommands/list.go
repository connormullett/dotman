package subcommands

import (
	"fmt"
	"os"

	"github.com/connormullett/dotman/util"
)

func List(args []string) {
	settings := util.ReadConfig()

	entries := GetManagedFilesList(settings.Path)
	for _, entry := range entries {
		fmt.Println(entry)
	}
}

func GetManagedFilesList(path string) []string {
	settings := util.ReadConfig()

	return settings.ManagedFiles
}

func GetFilesList(path string) []string {
	var files []string
	dir, err := os.ReadDir(path)
	if err != nil {
		fmt.Println("Error reading directory:", err)
		return files
	}

	for _, entry := range dir {
		if entry.IsDir() {
			continue
		}
		files = append(files, entry.Name())
	}

	return files
}
