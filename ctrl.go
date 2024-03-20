package main

import (
	"log/slog"
	"os"
	"os/signal"
	"pushup/single"
	"syscall"
)

func poll() {
	slog.Info("polling for events")
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-stop
}

func close() {
	single.Discord.Close()
}
