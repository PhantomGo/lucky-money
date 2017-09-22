package main

import (
	"flag"
	"fmt"
	"net/http"
	"runtime"
)

func main() {
	flag.Parse()
	runtime.GOMAXPROCS(runtime.NumCPU())
	http.Handle("/", mux)
	err := http.ListenAndServe(":8088", nil)
	if err != nil {
		fmt.Println(err)
	}
	// block until a signal is received.
	InitSignal()
}
