package main

import (
	"flag"
	"log"
	"os"
	"time"
)

func main() {
	// parse command line arguments
	delayPtr := flag.Int("interval", 5, "Run checks every N seconds")
	flag.Parse()

	// construct the compositeCollector and add all needed collectors
	var collector compositeCollector
	collector.Add(cpuCollector{})
	collector.Add(memCollector{})
	collector.Add(procCollector{})

	// create a logger
	logger := ioWriterLoggerWithTime{os.Stdout, defaultResultFormatter}

	// run an endless loop periodically calling the collectors and sending their
	// output to the logger
	for {
		_, err := collector.Collect()
		if err != nil {
			log.Panic(err)
		}
		for _, r := range collector.CollectedResults {
			logger.Log(r)
		}

		time.Sleep(time.Duration(*delayPtr) * time.Second)
	}
}
