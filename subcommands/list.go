package subcommands

import (
	"fmt"
	"log"
	"os"
)

func List(args []string) {
	settings := ReadConfig()

	entries := GetFilesList(settings.Path)
	for _, entry := range entries {
		fmt.Println(entry)
	}
}

func GetFilesList(path string) []string {
	entries, err := os.ReadDir(path)
	if err != nil {
		log.Fatalf("Failed to read directory: %v", err)
	}

	var files []string
	for _, entry := range entries {
		if !entry.IsDir() {
			files = append(files, entry.Name())
		}
	}
	return files
}
