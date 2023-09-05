package graceful

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func PerformGracefulShutdown(server *http.Server, httpErrChan chan error, httpShutdownChan chan struct{}) {
	// Context for graceful shutdown with a timeout of 5 seconds
	shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer shutdownCancel()

	// Channel for processing stop signals
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	// Block select that waits for either a signal from the operating system,
	// or an error from the server, and terminates the application accordingly
	select {
	case <-sigChan:
		fmt.Println("\nGraceful shutdown initiated...")
		if err := server.Shutdown(shutdownCtx); err != nil {
			fmt.Printf("HTTP server shutdown error: %v\n", err)
		} else {
			fmt.Println("HTTP server shutdown successfully")
		}
		close(httpShutdownChan)
	case err := <-httpErrChan:
		fmt.Printf("HTTP server error: %v\n", err)
	}

	fmt.Println("Application stopped")
}
