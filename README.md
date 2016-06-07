gomonitor
=========


gomonitor is a simple monitoring tool that allows you to periodically record certain system information.

It does so by calling collector instances and passing their output to a logger.

All current collector implementations work with linux only. But feel free to build your own collectors that work for you.

Currently implemented collectors:

* compositeCollector: a collector that wraps other collectors to call them all

* cpuCollector: records the average CPU load

* memCollector: records the currently free memory

* procCollector: records the number of processes started

* demoCollector: used when run with --demo option. Produces values resembling a sin curve
