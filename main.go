package main

import (
	"os"
	"os/signal"
	"ow-udp-tracker/schemes"
	"syscall"
	"time"
)

func main() {
	worker := NewWorker(handle)
	worker.Start()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	<-c
}

func handle(player schemes.Player, raw []byte) {
	println("Got packet at:", time.Now().Format("2006-01-02 15:04:05"))
	data, err := player.MarshalIndent()
	if err != nil {
		println(err)
		return
	}

	println(string(data))
}
