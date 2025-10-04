package subcommands

import (
	"fmt"

	"github.com/connormullett/dotman/util"
)

func List(args []string) {
	settings := util.ReadConfig()

	entries := GetFilesList(settings.Path)
	for _, entry := range entries {
		fmt.Println(entry)
	}
}

func GetFilesList(path string) []string {
	settings := util.ReadConfig()

	return settings.ManagedFiles
}
