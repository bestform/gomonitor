package main

import (
	"fmt"
	"io"
	"time"
)

type resultFormatter func(collectorResult) string

func defaultResultFormatter(c collectorResult) string {
	return fmt.Sprintf("%v: %v\n", c.Title, c.Content)
}

// stdLogger is a simple logger writing to an io.Writer instance
// It will add the date and time to the log output
type ioWriterLoggerWithTime struct {
	Writer    io.Writer
	Formatter resultFormatter
}

// Log writes the string to the io.Writer and adds date and time
func (l *ioWriterLoggerWithTime) Log(r collectorResult) error {
	_, err := l.Writer.Write([]byte(fmt.Sprintf("%v:\n%v\n", time.Now().Format(time.UnixDate), l.Formatter(r))))
	return err
}
