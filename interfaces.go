package main

// a collector can collect info for a specific area
type collector interface {
	Collect() (string, error)
}

// a logger will write collected strings
type logger interface {
	Log(string) error
}
