package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/pulse227/server-recruit-challenge-sample/api"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	r := api.NewRouter()

	server := &http.Server{
		Addr:    ":8888",
		Handler: r,
	}
	go func() {
		<-ctx.Done()
		fmt.Println("\nReceived Interupt")
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer func() {
			fmt.Println("Defer Gracefull")
			cancel()
		}()
		server.Shutdown(ctx)
		fmt.Println("After Shutdown")
	}()
	log.Println("server start running at :8888")
	log.Fatal(server.ListenAndServe())
}
