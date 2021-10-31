package shutdown

import (
	"log"
	"os"
	"os/signal"
	"syscall"
)

func Exit(cb func()) {
	sigs := make(chan os.Signal, 1)
	terminate := make(chan bool)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		sig := <-sigs
		log.Fatal("terminated due to: ", sig)
		close(terminate)
	}()

	<-terminate
	cb()
}
