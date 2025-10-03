package cmd

import (
	"encoding/json"
	"os"
	"os/user"
	"path/filepath"

	"github.com/kirsle/configdir"
)

type Settings struct {
	Path string `json:"path"`
}

func readConfig() Settings {
	configPath := configdir.LocalConfig("dotman")
	err := configdir.MakePath(configPath)
	if err != nil {
		panic((err))
	}

	configFile := filepath.Join(configPath, "config.json")

	var settings Settings

	// Does the file not exist?
	if _, err = os.Stat(configFile); os.IsNotExist(err) {
		user, err := user.Current()
		if err != nil {
			panic(err)
		}

		homedir := user.HomeDir

		dotmanPath := filepath.Join(homedir, ".dotman")

		settings = Settings{Path: dotmanPath}

		// Create the new config file.
		fh, err := os.Create(configFile)
		if err != nil {
			panic(err)
		}
		defer fh.Close()

		encoder := json.NewEncoder(fh)
		encoder.Encode(settings)
	} else {
		// Load the existing file.
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
