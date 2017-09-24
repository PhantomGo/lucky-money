package main

import (
	"lucky-money/service"
	"time"

	log "github.com/thinkboy/log4go"
)

var (
	Srv *service.Service
)

func InitService() (err error) {
	Srv = service.NewService()
	job := time.NewTicker(time.Minute * 1)
	go func() {
		for _ = range job.C {
			log.Debug("job start")
			Srv.ClearExpired()
			log.Debug("job finish")
		}
	}()
	return
}
