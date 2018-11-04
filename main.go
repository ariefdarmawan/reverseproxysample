package main

import (
	"os"
	"os/signal"
)

func main() {
	go helloApp()
	go welcomeApp()
	go proxyApp()

	cstop := make(chan os.Signal, 1)
	signal.Notify(cstop, os.Interrupt)
	<-cstop
}
