package base

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
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
