package skullsebitenbind

import (
	"github.com/hajimehoshi/ebiten/v2/mobile"
	"github.com/rootVIII/skulls"
)

func init() {

	game, err := skulls.Load()
	if err != nil {
		panic(err)
	}
	mobile.SetGame(game)
}

// Dummy forces gomobile to compile this package.
func Dummy() {}
