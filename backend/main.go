package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/Hans-Kerman/GinBlogPrimer/backend/config"
	"github.com/Hans-Kerman/GinBlogPrimer/backend/router"
)

func main() {
	config.InitConfig()

	r := router.SetupRouter()
	port := config.AppConfig.App.Port

	if port == "" {
		port = ":8080"
	}

	srv := &http.Server{
		Addr:    port,
		Handler: r,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit //在这里阻塞住主goroutine不退出
	log.Println("Shutdown Server...")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	log.Println("Server exiting")
}
