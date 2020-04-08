package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	router := http.NewServeMux() // here we could also go with third party packages to create a router
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World!"))
	})
	srv := &http.Server{
		Addr:    ":9620",
		Handler: router,
	}

	// it’s not necessary to wrap another go-routine with srv.ListenAndServe()
	// this method is blocking and in the go documentation it already described that will use separate go-routine for each incoming request.
	// The only reason I’m using go-routine to have another wrapper is because it’s more easier for me to handle the channel interactions and rest of the shutdown steps
	go func() {
		if err := srv.ListenAndServe(); err != http.ErrServerClosed {
			log.Fatalf("Server ListenAndServe Error: %v", err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	<-quit
	log.Print("Server Shutting Down")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer func() {
		// Close database, redis, truncate message queues, etc.
		cancel()
	}()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server Shutdown Error: %v", err)
	}
	log.Print("Server Shutdown Properly")
}

func useChannel() {
	done := make(chan os.Signal, 1)                                    // creates a done channel and it can only accept os.Signal type with 1 capacity
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM) // causes the package signal to relay incoming signals to done channel

	<-done // trying to receive output from done channel
	log.Print("Server Stopped")
}
