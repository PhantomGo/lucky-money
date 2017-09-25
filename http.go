package main

import (
	"encoding/json"
	"io/ioutil"
	"net"
	"net/http"
	"strconv"
	"time"

	log "github.com/thinkboy/log4go"
)

func InitHTTP() (err error) {
	// http listen
	var network, addr string
	httpServeMux := http.NewServeMux()
	httpServeMux.HandleFunc("/luckmoney/envelops", Histories)
	httpServeMux.HandleFunc("/luckmoney/envelop/open", Open)
	httpServeMux.HandleFunc("/luckmoney/envelop", Fill)
	httpServeMux.HandleFunc("/luckmoney/account/balance", Banlance)
	addr = Conf.HTTPAddr
	network = "tcp"
	log.Info("start http listen:\"%s\"", Conf.HTTPAddr)
	go httpListen(httpServeMux, network, addr)
	return
}

func httpListen(mux *http.ServeMux, network, addr string) {
	httpServer := &http.Server{Handler: mux, ReadTimeout: Conf.HTTPReadTimeout, WriteTimeout: Conf.HTTPWriteTimeout}
	httpServer.SetKeepAlivesEnabled(true)
	l, err := net.Listen(network, addr)
	if err != nil {
		log.Error("net.Listen(\"%s\", \"%s\") error(%v)", network, addr, err)
		panic(err)
	}
	if err := httpServer.Serve(l); err != nil {
		log.Error("server.Serve() error(%v)", err)
		panic(err)
	}
}

// retWrite marshal the result and write to client(get).
func retWrite(w http.ResponseWriter, r *http.Request, res map[string]interface{}, start time.Time) {
	data, err := json.Marshal(res)
	if err != nil {
		log.Error("json.Marshal(\"%v\") error(%v)", res, err)
		return
	}
	dataStr := string(data)
	if _, err := w.Write([]byte(dataStr)); err != nil {
		log.Error("w.Write(\"%s\") error(%v)", dataStr, err)
	}
	log.Info("req: \"%s\", get: res:\"%s\", ip:\"%s\", time:\"%fs\"", r.URL.String(), dataStr, r.RemoteAddr, time.Now().Sub(start).Seconds())
}

// retPWrite marshal the result and write to client(post).
func retPWrite(w http.ResponseWriter, r *http.Request, res map[string]interface{}, body *string, start time.Time) {
	data, err := json.Marshal(res)
	if err != nil {
		log.Error("json.Marshal(\"%v\") error(%v)", res, err)
		return
	}
	dataStr := string(data)
	if _, err := w.Write([]byte(dataStr)); err != nil {
		log.Error("w.Write(\"%s\") error(%v)", dataStr, err)
	}
	log.Info("req: \"%s\", post: \"%s\", res:\"%s\", ip:\"%s\", time:\"%fs\"", r.URL.String(), *body, dataStr, r.RemoteAddr, time.Now().Sub(start).Seconds())
}

func Histories(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method Not Allowed", 405)
		return
	}
	var (
		param = r.URL.Query()
		res   = map[string]interface{}{}
		aid   int
		err   error
	)
	if aid, err = strconv.Atoi(param.Get("account")); err != nil {
		res["error"] = err.Error()
	}
	defer retWrite(w, r, res, time.Now())
	if ret, err := Srv.Histories(int64(aid)); err != nil {
		res["error"] = err.Error()
	} else {
		res["history"] = ret
	}

	return
}

func Open(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method Not Allowed", 405)
		return
	}
	var (
		param     = r.Form
		res       = map[string]interface{}{}
		aid       int
		code      string
		err       error
		body      string
		bodyBytes []byte
	)
	if aid, err = strconv.Atoi(param.Get("account")); err != nil {
		res["error"] = err.Error()
	}
	if code = param.Get("code"); len(code) != 8 {
		res["error"] = "code is illegal"
	}
	defer retPWrite(w, r, res, &body, time.Now())
	if bodyBytes, err = ioutil.ReadAll(r.Body); err != nil {
		log.Error("ioutil.ReadAll() failed (%s)", err)
		res["error"] = err.Error()
		return
	}
	body = string(bodyBytes)
	if ret, err := Srv.Open(int64(aid), code); err != nil {
		res["error"] = err.Error()
	} else {
		res["code"] = ret
	}

	return
}

func Fill(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method Not Allowed", 405)
		return
	}
	var (
		param     = r.Form
		res       = map[string]interface{}{}
		aid       int
		amount    int
		number    int
		err       error
		body      string
		bodyBytes []byte
	)
	if aid, err = strconv.Atoi(param.Get("account")); err != nil {
		res["error"] = err.Error()
	}
	if amount, err = strconv.Atoi(param.Get("amount")); err != nil {
		res["error"] = err.Error()
	}
	if number, err = strconv.Atoi(param.Get("number")); err != nil {
		res["error"] = err.Error()
	}
	if number > amount {
		res["error"] = err.Error()
	}
	defer retPWrite(w, r, res, &body, time.Now())
	if bodyBytes, err = ioutil.ReadAll(r.Body); err != nil {
		log.Error("ioutil.ReadAll() failed (%s)", err)
		res["error"] = err.Error()
		return
	}
	body = string(bodyBytes)
	if ret, err := Srv.Fill(int64(aid), int64(amount), number); err != nil {
		res["error"] = err.Error()
	} else {
		res["code"] = ret
	}

	return
}

func Banlance(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method Not Allowed", 405)
		return
	}
	var (
		param = r.URL.Query()
		res   = map[string]interface{}{}
		aid   int
		err   error
	)
	if aid, err = strconv.Atoi(param.Get("account")); err != nil {
		res["error"] = err.Error()
	}
	defer retWrite(w, r, res, time.Now())
	if ret, err := Srv.Balance(int64(aid)); err != nil {
		res["error"] = err.Error()
	} else {
		res["balance"] = ret
	}

	return
}

func AddAccount(w http.ResponseWriter, r *http.Request) {

}
