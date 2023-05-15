package logger

import (
	"fmt"
	"os"
	"time"

	"github.com/potatoschool/shapes/config"
)

var file *os.File

func Start() {
	ts := time.Now()

	if _, err := os.Stat(config.Conf.LogFolder); err != nil {
		os.Mkdir(config.Conf.LogFolder, 0755)
	}

	filepath := fmt.Sprintf("%s/%s.txt", config.Conf.LogFolder, genDate(&ts))

	target, err := os.OpenFile(filepath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}

	file = target
}

func End() error {
	return file.Close()
}
