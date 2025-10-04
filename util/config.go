package util

import (
	"encoding/json"
	"os"
	"os/user"
	"path/filepath"
	"slices"

	"github.com/kirsle/configdir"
)

func WriteConfig(settings Settings) {
	configFile := GetConfigFilePath()
	fh, err := os.Create(configFile)
	if err != nil {
		panic(err)
	}
	defer fh.Close()

	encoder := json.NewEncoder(fh)
	encoder.Encode(settings)
}

type Settings struct {
	Path         string   `json:"path"`
	ManagedFiles []string `json:"managedFiles"`
}

func GetConfigFilePath() string {
	configPath := configdir.LocalConfig("dotman")
	err := configdir.MakePath(configPath)
	if err != nil {
		panic((err))
	}
	configFile := filepath.Join(configPath, "config.json")
	return configFile
}

func AddManagedFile(file string) {
	settings := ReadConfig()
	if slices.Contains(settings.ManagedFiles, file) {
		return
	}
	settings.ManagedFiles = append(settings.ManagedFiles, file)

	WriteConfig(settings)
}

func ReadConfig() Settings {
	configFile := GetConfigFilePath()

	var settings Settings

	if _, err := os.Stat(configFile); os.IsNotExist(err) {
		// file doesnt exist
		user, err := user.Current()
		if err != nil {
			panic(err)
		}

		homedir := user.HomeDir

		dotmanPath := filepath.Join(homedir, ".dotman")

		settings = Settings{Path: dotmanPath, ManagedFiles: []string{}}

		WriteConfig(settings)
	} else {
		// file exists
		fh, err := os.Open(configFile)
		if err != nil {
			panic(err)
		}
		defer fh.Close()

		decoder := json.NewDecoder(fh)
		decoder.Decode(&settings)
	}

	return settings
}
