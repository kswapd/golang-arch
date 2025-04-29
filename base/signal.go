package base

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func StartMySignal() {

	// Create a channel to receive signals
	sigs := make(chan os.Signal, 5)

	// Notify the channel on SIGINT (Ctrl+C) and SIGTERM (kill command)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	// Block until a signal is received
	fmt.Println("Waiting for a signal...")
	sig := <-sigs
	fmt.Println("Received signal:", sig)

	// Perform cleanup or graceful shutdown here
	fmt.Println("Exiting...")

}

func WaitGoroutine() {

	done := make(chan struct{})
	//defer close(done)

	go func() {
		//time.Sleep(time.Second * 2)
		fmt.Println("Goroutine done")
		close(done)
		//done <- struct{}{}
	}()
	time.Sleep(time.Second * 5)
	<-done
	// Perform cleanup or graceful shutdown here
	fmt.Println("Exiting main routine...")

}
