package app

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/mux"
	"github.com/kevinsudut/single-fizz-buzz/app/handler"
	"github.com/kevinsudut/single-fizz-buzz/pkg/lib/log"
)

func Init() error {
	server := &http.Server{
		Addr: ":8000",
		Handler: http.TimeoutHandler(
			handler.Init().RegisterHandlers(mux.NewRouter()),
			1*time.Second, // 1 second as timeout
			"",
		),
		ReadTimeout:  1 * time.Second, // 1 second as timeout
		WriteTimeout: 1 * time.Second, // 1 second as timeout
	}

	// Terminated gracefully using SIGTERM
	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt, syscall.SIGTERM)
		<-sigint
		if err := server.Shutdown(context.TODO()); err != nil {
			log.Errorln("Error shutting down server", err)
		}
	}()

	return server.ListenAndServe()
}
