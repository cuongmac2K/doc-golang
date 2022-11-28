package main

import "time"

type TelegramLog struct {
	URL          string
	Error        interface{}
	Header       interface{}
	Res          interface{}
	Function     interface{}
	Method       string
	Params       interface{}
	Title        string
	ErrorMessage string
	ResponseTime time.Duration
	Response     string
	Body         string
}

var Env = "dev"
var TELEGRAM_CONFIG = map[string]string{
	"BOT_LOG_REQUEST": "1619775560:AAGV2G4lGvFBGKBC389_XbmRLnLQgxMICgs",
	"DOMAIN":          "https://api.telegram.org/",
}
