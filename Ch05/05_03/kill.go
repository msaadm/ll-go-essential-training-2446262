package main

import (
	"fmt"
	"log"
)

func killServer(pidFile string) error {
	pid := 0

	// Simulate kill
	fmt.Printf("killing server with pid=%d\n", pid)
	return nil
}

func main() {
	if err := killServer("server.pid"); err != nil {
		log.Fatalf("error: %s", err)
	}
}
