package logger

import (
	"fmt"
	"time"

	"github.com/potatoschool/shapes/config"
)

func Log(message string) {
	ts := time.Now()

	logmessage := fmt.Sprintf("[shapes-debug] %s: %s\n", genDate(&ts), message)

	file.Write([]byte(logmessage))

	if config.Conf.Mode == "DEBUG" {
		fmt.Print(logmessage)
	}
}
