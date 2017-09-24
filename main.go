package main

import (
	"flag"
	"runtime"

	log "github.com/thinkboy/log4go"
)

func main() {
	flag.Parse()
	runtime.GOMAXPROCS(runtime.NumCPU())
	InitConfig()
	log.LoadConfiguration(Conf.Log)
	if err := InitService(); err != nil {
		panic(err)
	}
	if err := InitHTTP(); err != nil {
		panic(err)
	}
	// block until a signal is received.
	InitSignal()
}
