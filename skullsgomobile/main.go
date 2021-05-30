package main

import (
	_ "image/jpeg"
	_ "image/png"
	"log"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/rootVIII/skulls"
)

func exitErr(err error) {
	logf, _ := os.OpenFile("SKULLS-ERROR.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	log.SetOutput(logf)
	log.Println("** An error occurred during startup **")
	log.Fatal(err)
}

func main() {
	game, err := skulls.Load()
	if err != nil {
		exitErr(err)
	}
	if err := ebiten.RunGame(game); err != nil {
		exitErr(err)
	}
}
