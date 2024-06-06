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
		for i := 0; i < 20; i++ {
			wg.Add(1)
			defer wg.Done()
			_ = sentry.Init(sentry.ClientOptions{
				Debug: true,
				Dsn:   "http://2c95a1e4c38e6b12128bcb2515be95bc@10.3.55.219:9000/2",
			})
			go func() {
				for i := 0; i < 10000; i++ {
					sentry.CaptureException(errors.New(fmt.Sprintf("2my error11 %s", i)))
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
