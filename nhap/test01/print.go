package main

import (
	"doc-golang/nhap/logger"
	"errors"
	"github.com/getsentry/sentry-go"
	"net/http"
)

func main() {
	dns := "https://650d55e7a43c408ca7bbb0fd5a061d5c@sentry.paas.vn/323"
	if err := sentry.Init(sentry.ClientOptions{
		Dsn:         dns,
		Debug:       true,
		Environment: "test",
		BeforeSend: func(event *sentry.Event, hint *sentry.EventHint) *sentry.Event {
			if hint.Context != nil {
				if req, ok := hint.Context.Value(sentry.RequestContextKey).(*http.Request); ok {
					logger.Info(req)
				}
			}

			return event
		},
	}); err != nil {
		logger.Info("Sentry initialization failed: %v\n", err)
	}
	// Gọi hàm Debug với nội dung log "Hello, world!"
	sentry.CaptureException(errors.New("new Errror"))
	logger.Debug("Hello, world!")
	logger.Info("Hello, world!")
}
