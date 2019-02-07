package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	rpio "github.com/stianeikeland/go-rpio"
)

func main() {
	err := rpio.Open()
	if err != nil {
		log.Fatalf("Error occurred while trying to connect to rpio pins. Error=%s \n", err)
	}
	defer rpio.Close()
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go blinkForever(rpio.Pin(17))
	<-c
	os.Exit(0)
}

func blinkForever(led rpio.Pin) {
	led.Output()
	for {
		led.High()
		time.Sleep(time.Second)
		led.Low()
	}
}
