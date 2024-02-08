package main

import (
	"context"
	"fmt"
	"github.com/isso-719/counter-api/src/infra/config"
	"github.com/isso-719/counter-api/src/infra/router"
	"log"
	"net/http"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	httpConfig := config.LoadHTTPConfig()

	router := router.InitRouter(ctx)
	addr := httpConfig.Host + ":" + httpConfig.Port
	srv := &http.Server{
		Addr:           addr,
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	go func() {
		fmt.Println("Server started at " + addr)
		if err := srv.ListenAndServe(); err != nil {
			log.Fatal("Server failed to start: ", err)
		}
	}()
	<-ctx.Done()
	stop()
	log.Println("shutting down gracefully, press Ctrl+C again to force")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown: ", err)
	}

	log.Println("Server exiting")
}
