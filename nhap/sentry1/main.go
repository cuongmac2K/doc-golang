package main

import (
	"errors"
	"fmt"
	"github.com/getsentry/sentry-go"
	"time"
)

func main() {

	_ = sentry.Init(sentry.ClientOptions{
		Debug: true,
		//Dsn:   "https://ea001c39b67a49b5b8df1f20df80d98e@sentry.paas.vn/316",
		Dsn: "http://20e99356869f15ce83a56cc44a5fe085@10.3.54.229:9000/4",
		//Dsn: "https://833df11a7bad43908723ae6fb36e4518:9601c2ed2d7f443ebf1f049ef2504f24@sentry.paas.vn/317",
		//Dsn: "https://sentry.paas.vn/api/hooks/release/builtin/317/2e9b295007bc1005fae3bb862054a59c64fc7ffa038350f3729a0ea7421d097b",
	})
	sentry.CaptureException(errors.New(fmt.Sprintf("1my error11")))
	time.Sleep(time.Second)
	// Since sentry emits events in the background we need to make sure
	// they are sent before we shut down
	sentry.Flush(time.Second * 5)
}
