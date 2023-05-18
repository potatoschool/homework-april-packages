package logger

import (
	"fmt"
	"time"
)

func genDate(ts *time.Time) string {
	return fmt.Sprintf("%d-%d-%d %d:%d:%d",
		ts.Year(),
		ts.Month(),
		ts.Day(),
		ts.Hour(),
		ts.Minute(),
		ts.Second(),
	)
}
