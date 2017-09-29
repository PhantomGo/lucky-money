package main

import "time"

var (
	Conf *Config
)

type Config struct {
	HTTPAddr         string
	HTTPReadTimeout  time.Duration
	HTTPWriteTimeout time.Duration
	Password         string
	Log              string
}

func InitConfig() {
	Conf = &Config{
		HTTPAddr:         "127.0.0.1:9092",
		HTTPReadTimeout:  60 * 1000,
		HTTPWriteTimeout: 60 * 1000,
		Password:         "123456",
		Log:              "./luckymoney-log.xml",
	}
}
