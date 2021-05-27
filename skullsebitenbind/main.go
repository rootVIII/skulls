package skullsebitenbind

import (
	"github.com/hajimehoshi/ebiten/v2/mobile"
	"github.com/rootVIII/skulls"
)

func init() {

	sp, err := skulls.Play()
	if err != nil {
		panic(err)
	}
	mobile.SetGame(sp)
}

// Dummy forces gomobile to compile this package.
func Dummy() {}
