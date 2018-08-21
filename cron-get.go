package main

import (
	"fmt"
	"github.com/robfig/cron"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	c := cron.New()
	c.AddFunc(os.Getenv("SCHEDULE"), func() {
		log.Print("Fetching " + os.Getenv("URL"))
		resp, err := http.Get(os.Getenv("URL"))
		if err != nil {
			log.Print("Failed: ", err)
		}
		if resp.StatusCode >= 400 {
			log.Print("Failed: ", resp.Status)
		}
		io.Copy(os.Stdout, resp.Body)
	})
	c.Start()

	// Now Run until SIGTERM or SIGINT
	sigs := make(chan os.Signal, 1)
	done := make(chan bool, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-sigs
		done <- true
	}()

	<-done
	fmt.Println("\nexiting")
}
