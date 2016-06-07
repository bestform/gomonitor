package main

import (
	"fmt"
	"io"
	"time"
)

// stdLogger is a simple logger writing to an io.Writer instance
// It will add the date and time to the log output
type stdOutLogger struct {
	Writer io.Writer
}

// Log writes the string to the io.Writer and adds date and time
func (l *stdOutLogger) Log(s string) error {
	_, err := l.Writer.Write([]byte(fmt.Sprintf("%v:\n%v\n", time.Now().Format(time.UnixDate), s)))
	return err
}
