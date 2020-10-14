package main

import (
	"context"
	_ "github.com/joho/godotenv/autoload"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"goevent/database"
	"goevent/router"
)


func main() {


	router := router.SetupRouter()

	_ = database.Connect()

	srv := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down the server")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("shutdown:", err)
	}
	log.Println("Server OK")
}
