package main

import (
	"fmt"
	"math"

	linuxproc "github.com/c9s/goprocinfo/linux"
)

// compositeCollector groups collectors and runs them one by one
// The result of a Collect call is a combination of all outputs of its children
// Should one child yield an error, compositeCollector will abort and return the
// error
type compositeCollector struct {
	Collectors []collector
}

type collectorResult struct {
	Title, Content string
}

// Collect will return a combination off all outputs
// It will abort and return an error, should one child fail
func (c *compositeCollector) Collect() ([]collectorResult, error) {
	var collectedResults []collectorResult
	for _, child := range c.Collectors {
		result, err := child.Collect()
		if err != nil {
			return []collectorResult{collectorResult{}}, err
		}
		collectedResults = append(collectedResults, result...)
	}

	return collectedResults, nil
}

// Add lets you add a new collector to the collection
func (c *compositeCollector) Add(nc collector) {
	c.Collectors = append(c.Collectors, nc)
}

// CPU collector
type cpuCollector struct{}

func (c *cpuCollector) Collect() ([]collectorResult, error) {
	loadAvg, err := linuxproc.ReadLoadAvg("/proc/loadavg")
	if err != nil {
		return []collectorResult{collectorResult{}}, err
	}

	return []collectorResult{{"Load Average", fmt.Sprintf("%v", loadAvg.Last5Min)}}, nil
}

// Memory collector
type memCollector struct{}

func (c *memCollector) Collect() ([]collectorResult, error) {
	mem, err := linuxproc.ReadMemInfo("/proc/meminfo")
	if err != nil {
		return []collectorResult{collectorResult{}}, err
	}

	return []collectorResult{{"Free Memory", fmt.Sprintf("%v", mem.MemFree)}}, nil
}

// process collector
type procCollector struct{}

func (c *procCollector) Collect() ([]collectorResult, error) {
	stat, err := linuxproc.ReadStat("/proc/stat")
	if err != nil {
		return []collectorResult{collectorResult{}}, err
	}

	return []collectorResult{{"Processes running", fmt.Sprintf("%v", stat.Processes)}}, nil
}

// demo collector
type demoCollector struct {
	counter float64
}

func (c *demoCollector) Collect() ([]collectorResult, error) {
	r := math.Sin(c.counter)
	c.counter += 0.3
	return []collectorResult{{"Demo", fmt.Sprintf("%f", r)}}, nil
}

// Memory Available Collector
type memAvailableCollector struct{}

func (c *memAvailableCollector) Collect() ([]collectorResult, error) {
	buf, err := linuxproc.ReadMemInfo("/proc/meminfo")
	if err != nil {
		return []collectorResult{{}}, err
	}

	return []collectorResult{{"Mem Available", fmt.Sprintf("%v", buf.MemAvailable)}}, nil
}

// Memory Buffer Collector
type memBufferCollector struct{}

func (c *memBufferCollector) Collect() ([]collectorResult, error) {
	buf, err := linuxproc.ReadMemInfo("/proc/meminfo")
	if err != nil {
		return []collectorResult{{}}, err
	}

	return []collectorResult{{"Buffers", fmt.Sprintf("%v", buf.Buffers)}}, nil
}

// Memory Cached Collector
type memCachedCollector struct{}

func (c *memCachedCollector) Collect() ([]collectorResult, error) {
	buf, err := linuxproc.ReadMemInfo("/proc/meminfo")
	if err != nil {
		return []collectorResult{{}}, err
	}

	return []collectorResult{{"Memory Cached", fmt.Sprintf("%v", buf.Cached)}}, nil
}
