package main

import (
	_ "image/jpeg"
	_ "image/png"
	"log"
	"os"

	"github.com/rootVIII/skulls"
)

func main() {
	_, err := skulls.Play()
	if err != nil {
		logf, _ := os.OpenFile("SKULLS-ERROR.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		log.SetOutput(logf)
		log.Println("** An error occurred during startup **")
		log.Fatal(err)
	}
}
