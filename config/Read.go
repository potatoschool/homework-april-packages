package config

import (
	"encoding/json"
	"os"
)

type Configuration struct {
	Mode      string `json:"mode"`
	LogFolder string `json:"logFolder"`
}

var Conf Configuration

func Read(path string) {
	Conf = Configuration{
		Mode:      "PRODUCTION",
		LogFolder: "./_log",
	}

	file, err := os.ReadFile(path)

	if err == nil {
		json.Unmarshal(file, &Conf)
	}
}
