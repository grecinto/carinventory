package apihandler

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func listenAndServe(srv *http.Server) error {
	errc := make(chan error, 1)

	go func() {
		err := srv.ListenAndServe(); 
		if err != http.ErrServerClosed {
			errc<-err
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	
	select {
	case <-c:
		ctx, cancel := context.WithTimeout(context.Background(), time.Second * 15)
		defer cancel()
		srv.Shutdown(ctx)
	case err := <-errc:
		return err
	}

	return nil
}
