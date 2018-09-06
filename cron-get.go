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
	"time"
)

func main() {
	var c *cron.Cron
	if os.Getenv("TZ") != "" {
		tz, err := time.LoadLocation(os.Getenv("TZ"))
		if err != nil {
			log.Fatal(err)
		}
		c = cron.NewWithLocation(tz)
	} else {
		c = cron.New()
	}
	err := c.AddFunc(os.Getenv("SCHEDULE"), func() {
		log.Print("Fetching " + os.Getenv("URL"))
		get, err := http.NewRequest("GET", os.Getenv("URL"), nil)
		if os.Getenv("USERNAME")+os.Getenv("PASSWORD") != "" {
			get.SetBasicAuth(os.Getenv("USER"), os.Getenv("PASS"))
		}
		if err != nil {
			log.Print("Failed: ", err)
		}
		resp, err := http.Client{}.Do(get)
		if err != nil {
			log.Print("Failed: ", err)
		}
		if resp.StatusCode >= 400 {
			log.Print("Failed: ", resp.Status)
		}
		io.Copy(os.Stdout, resp.Body)
	})
	if err != nil {
		log.Fatal(err)
	}
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
