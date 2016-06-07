package main

import (
	"fmt"

	linuxproc "github.com/c9s/goprocinfo/linux"
)

// compositeCollector groups collectors and runs them one by one
// The result of a Collect call is a combination of all outputs of its children
// Should one child yield an error, compositeCollector will abort and return the
// error
type compositeCollector struct {
	Collectors []collector
}

// Collect will return a combination off all outputs
// It will abort and return an error, should one child fail
func (c compositeCollector) Collect() (string, error) {
	var completeResult string
	for _, child := range c.Collectors {
		result, err := child.Collect()
		if err != nil {
			return "", err
		}
		completeResult = completeResult + result + "\n"
	}

	return completeResult, nil
}

// Add lets you add a new collector to the collection
func (c *compositeCollector) Add(nc collector) {
	c.Collectors = append(c.Collectors, nc)
}

// CPU collector
type cpuCollector struct{}

func (c cpuCollector) Collect() (string, error) {
	loadAvg, err := linuxproc.ReadLoadAvg("/proc/loadavg")
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("Load Average last 5 minutes: %f", loadAvg.Last5Min), nil
}

// Memory collector
type memCollector struct{}

func (c memCollector) Collect() (string, error) {
	mem, err := linuxproc.ReadMemInfo("/proc/meminfo")
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("Free Memory: %v", mem.MemFree), nil
}

// process collector
type procCollector struct{}

func (c procCollector) Collect() (string, error) {
	stat, err := linuxproc.ReadStat("/proc/stat")
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("Processes running: %v", stat.Processes), nil
}
