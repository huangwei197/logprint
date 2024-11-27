package logprint

import (
	"fmt"
	"time"
)

func Error(err error) {
	t := time.Now()
	fmt.Printf("[ERROR] %s: %s\n", t.Format("2006-01-02 15:04:05.000"), err.Error())
}
