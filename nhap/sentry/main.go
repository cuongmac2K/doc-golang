package main

import (
	"errors"
	"fmt"
	"github.com/getsentry/sentry-go"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	{
		for i := 0; i < 400; i++ {
			wg.Add(1)
			defer wg.Done()
			go func() {
				for i := 0; i < 100000; i++ {
					_ = sentry.Init(sentry.ClientOptions{
						Debug: true,
						Dsn:   "http://8c2963bae57a836866e060fcf1d1ed63@10.3.54.253/6",
					})
					sentry.CaptureException(errors.New(fmt.Sprintf("my error11 %s", i)))
					time.Sleep(time.Second)
				}
			}()
		}
	}
	wg.Wait()
	// Since sentry emits events in the background we need to make sure
	// they are sent before we shut down
	sentry.Flush(time.Second * 5)
}
