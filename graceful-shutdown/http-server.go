package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func hello(w http.ResponseWriter, req *http.Request) {
	const loadTime = 10
	for i := 0; i < loadTime; i++ {
		fmt.Printf("Loading %v/%v ..\n", i+1, loadTime)
		time.Sleep(time.Second)
	}
	fmt.Fprintln(w, "<h1>", "Hello", req.URL.Path[1:], "</h1>")
}

func runServer(server *http.Server, stopTime time.Duration) error {
	done := make(chan error, 1)

	go func() {
		ch := make(chan os.Signal, 1)
		signal.Notify(ch, os.Interrupt, syscall.SIGTERM)
		<-ch

		fmt.Println("Server shutdown initiated ..")

		ctx := context.Background()
		var cancel context.CancelFunc

		if stopTime > 0 {
			ctx, cancel = context.WithTimeout(context.Background(), stopTime)
			defer cancel()
		}

		done <- server.Shutdown(ctx)

		fmt.Println("Completing shutdown ..")
	}()

	err := server.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		return err
	}

	return <-done
}

func main() {
	timeout := flag.Int("timeout", 4, "Number of seconds to kill since shutdown has initiated")
	flag.Parse()

	server := &http.Server{
		Addr:    ":8080",
		Handler: http.HandlerFunc(hello),
	}
	err := runServer(server, time.Second*time.Duration(*timeout))
	if err != nil {
		log.Fatal(err)
	}
}
