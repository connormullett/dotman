package subcommands

import (
	"fmt"
	"log"
	"os"
)

func List(args []string) {
	entries := GetFilesList()
	for _, entry := range entries {
		fmt.Println(entry)
	}
}

func GetFilesList() []string {
	settings := ReadConfig()

	path := settings.Path

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
