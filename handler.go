package main

import (
	"lucky-money/service"
	"net/http"
	"strings"
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

func AuthHandler(handler func(http.ResponseWriter, *http.Request)) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		_, ok := r.Header["Authorization"]
		if !ok {
			http.Error(w, "Need Authorization Header", http.StatusBadRequest)
			return
		}

		auth := strings.SplitN(r.Header["Authorization"][0], " ", 2)

		if len(auth) != 2 {
			http.Error(w, "Bad Syntax", http.StatusBadRequest)
			return
		}

		if !authorize(auth[0], auth[1]) {
			http.Error(w, "Authorization Failed", http.StatusUnauthorized)
			return
		}
		handler(w, r)
	}
}

func authorize(scheme, password string) bool {
	return Conf.Password == password
}
