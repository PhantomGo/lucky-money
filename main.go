package main

import (
	"flag"
	"runtime"
)

func main() {
	flag.Parse()
	runtime.GOMAXPROCS(runtime.NumCPU())
	InitConfig()
	if err := InitService(); err != nil {
		panic(err)
	}
	if err := InitHTTP(); err != nil {
		panic(err)
	}
	// block until a signal is received.
	InitSignal()
}
