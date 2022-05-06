package main

import (
	"fmt"
	"github.com/harlesbayu/kuncie/internal/interface/app/http"
	"github.com/harlesbayu/kuncie/internal/interface/container"
	"github.com/harlesbayu/kuncie/internal/shared/config"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	cont := container.NewContainer(config.NewConfig("./resources"))

	fmt.Println("--------------------------------------------")
	fmt.Printf("HTTP API: %s (%s) \n", cont.Config.Name, cont.Config.Version)
	fmt.Println("--------------------------------------------")

	initGracefulShutdown()
	err := http.StartHTTPServer(cont)
	if err != nil {
		panic(err)
	}
}

func initGracefulShutdown() {
	s := make(chan os.Signal, 1)
	signal.Notify(s, os.Interrupt)
	signal.Notify(s, syscall.SIGTERM)
	go func() {
		<-s
		fmt.Println("Shutting down gracefully.")
		// clean up here
		os.Exit(1)
	}()
}
