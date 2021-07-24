/*
Package skulls
A simple strategy game about skulls
Copyright (C) 2021 rootVIII

colleyloyejames@gmail.com

This program is free software; you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation; either version 2 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License along
with this program; if not, write to the Free Software Foundation, Inc.,
51 Franklin Street, Fifth Floor, Boston, MA 02110-1301 USA.
*/
package skulls

import (
	"bytes"
	"fmt"
	"image"
	"image/color"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2/inpututil"

	// Required for Ebiten.
	_ "image/jpeg"
	_ "image/png"

	"github.com/goki/freetype/truetype"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/audio/mp3"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/rootVIII/skulls/assets"
	"golang.org/x/image/font"
)

const (
	screenW        = 640
	screenH        = 960
	frame          = 64
	rowMax         = 22
	colMax         = 15
	hotspotUpX     = 123.0
	hotspotUpY     = 788.0
	hotspotLeftX   = 39.0
	hotspotLeftY   = 862.0
	hotspotRightX  = 210.0
	hotspotRightY  = 865.0
	hotspotButtonX = 373.0
	hotspotButtonY = 818.0
)

// Clock is the global game clock.
var Clock float64

// Game controls overall gameplay.
type Game struct {
	background, explosion, intro         *ebiten.Image
	skulls                               map[string]*ebiten.Image
	planchette, onDeck                   []string
	skullColors                          []string
	skullCollector                       [][]string
	skullCoords                          [][][]int
	empties                              [][2]int
	searchHead                           [2]int
	beep, clear, track                   *audio.Player
	fontFace                             font.Face
	green                                color.Color
	isMovingL, isMovingR, isMovingU      bool
	havePlanchette, isPlaying            bool
	wonGame, lostGame                    bool
	rollCount, moveCount, explosionCount int
	jumpCount, jumpMax, loseCount        int
	score, level, best, matchMin         int
}

/*  - - - - - - - -  U P D A T E   M E T H O D S  - - - - - - - -  */

func (g *Game) rollBones() {

	g.planchette = append(g.planchette[len(g.planchette)-1:], g.planchette[:len(g.planchette)-1]...)

	row, col := g.searchHead[0], g.searchHead[1]
	for _, skull := range g.planchette {
		g.skullCollector[row][col] = skull
		row++
	}
}

func (g *Game) removeCurrentPos() {
	row, col := g.searchHead[0], g.searchHead[1]
	for range g.planchette {
		g.skullCollector[row][col] = ""
		row++
	}
}

func (g *Game) shiftPlanchette() {
	row, col := g.searchHead[0], g.searchHead[1]
	for _, skull := range g.planchette {
		g.skullCollector[row][col] = skull
		row++
	}
}

func (g *Game) matchSkulls() [][2]int {
	return append(g.checkCols(), g.checkRows()...)
}

func (g *Game) removeEmpties() {
	for _, coords := range g.empties {
		g.skullCollector[coords[0]][coords[1]] = ""
	}
}

func (g *Game) checkRows() [][2]int {

	var matchesIndex = make([][2]int, 0)
	var matchColor string
	var matchCount int

	for x := 0; x < len(g.skullCollector[0]); x++ {

		matchColor = ""
		matchCount = 0

		for y := 0; y < len(g.skullCollector); y++ {
			if g.skullCollector[y][x] == matchColor {
				matchCount++
			} else {
				matchCount = 0
			}
			var last = (y == len(g.skullCollector)-1) || (g.skullCollector[y+1][x] != g.skullCollector[y][x])
			if matchCount > 2 && last && len(matchColor) > 0 {
				for index := y - matchCount; index <= y; index++ {
					matchesIndex = append(matchesIndex, [2]int{index, x})
				}
			}
			matchColor = g.skullCollector[y][x]
		}
	}

	return matchesIndex
}

func (g *Game) checkCols() [][2]int {

	var matchesIndex = make([][2]int, 0)
	var matchColor string
	var matchCount int

	for ri, row := range g.skullCollector {
		matchColor = ""
		matchCount = 0
		for ci, col := range row {
			if col == matchColor {
				matchCount++
			} else {
				matchCount = 0
			}
			if matchCount > 2 && (ci == len(row)-1 || g.skullCollector[ri][ci+1] != col) && len(matchColor) > 1 {
				for index := ci - matchCount; index <= ci; index++ {
					matchesIndex = append(matchesIndex, [2]int{ri, index})
				}
			}
			matchColor = col
		}
	}
	return matchesIndex
}

func (g *Game) bubbleSortSkulls() {
	var length = len(g.skullCollector)

	for x := 0; x < len(g.skullCollector[0]); x++ {
		for {
			var madeChange = false
			for y := 0; y < length; y++ {
				for i := 0; i < length-y-1; i++ {

					if len(g.skullCollector[y][x]) < 1 && len(g.skullCollector[y+1][x]) > 1 {
						g.skullCollector[y][x], g.skullCollector[y+1][x] = g.skullCollector[y+1][x], g.skullCollector[y][x]
						madeChange = true
					}
				}
			}
			if !madeChange {
				break
			}
		}
	}

}

func (g *Game) isValidMove(direction byte) bool {

	row, col := g.searchHead[0], g.searchHead[1]
	switch direction {

	case 0x4C:
		if col < 1 || len(g.skullCollector[row][col-1]) > 0 {
			return false
		}
	case 0x52:
		if col > 13 || len(g.skullCollector[row][col+1]) > 0 {
			return false
		}
	case 0x55:
		if row < 1 || len(g.skullCollector[row-1][col]) > 0 {

			for {
				var combined = g.matchSkulls()
				if len(combined) > 0 {
					g.clear.Play()
					g.score += len(combined)
					g.best = g.score
					g.empties = combined
					g.explosionCount = 30
					g.removeEmpties()
					g.bubbleSortSkulls()
					g.clear.Rewind()
				} else {
					break
				}
			}

			g.reset()
			return false
		}
	}

	return true
}

func (g Game) inHotSpot(courseX, courseY float64) bool {

	touchX, touchY := ebiten.TouchPosition(0)

	if float64(touchX) < courseX || float64(touchX) > courseX+frame {
		return false
	}
	if float64(touchY) < courseY || float64(touchY) > courseY+frame {
		return false
	}
	return true

}

func (g Game) hotspotClickedLeft() bool {
	return g.inHotSpot(hotspotLeftX, hotspotLeftY)
}

func (g Game) hotspotClickedRight() bool {
	return g.inHotSpot(hotspotRightX, hotspotRightY)
}

func (g Game) hotspotClickedUp() bool {
	return g.inHotSpot(hotspotUpX, hotspotUpY)
}

func (g Game) hotspotClickedButton() bool {
	return g.inHotSpot(hotspotButtonX, hotspotButtonY)
}

func (g *Game) updatePlanchette() {

	if !g.lostGame && g.hotspotClickedLeft() {
		g.isMovingL = true
		g.isMovingR = false
		g.isMovingU = false
	}

	if !g.lostGame && g.hotspotClickedRight() {
		g.isMovingR = true
		g.isMovingL = false
		g.isMovingU = false
	}

	if g.hotspotClickedUp() {
		g.isMovingR = false
		g.isMovingL = false
		g.isMovingU = true
	}

	if g.isMovingL && g.moveCount < 1 && g.isValidMove('L') {

		g.removeCurrentPos()

		g.searchHead[1] -= 1
		g.shiftPlanchette()
	}

	if g.isMovingR && g.moveCount < 1 && g.isValidMove('R') {

		g.removeCurrentPos()

		g.searchHead[1] += 1
		g.shiftPlanchette()
	}

	if (g.isMovingU && g.moveCount < 1 || g.jumpCount == 0) && g.isValidMove('U') {

		g.removeCurrentPos()

		g.searchHead[0] -= 1
		g.shiftPlanchette()
	}

	if inpututil.IsTouchJustReleased(0) && (g.isMovingL || g.isMovingR || g.isMovingU) {
		g.isMovingL, g.isMovingR, g.isMovingU = false, false, false
	}

	if g.hotspotClickedButton() {
		if g.rollCount < 1 {
			g.beep.Play()
			g.rollBones()
			g.beep.Rewind()
		}
		g.rollCount++
	}

	if g.isMovingL || g.isMovingR || g.isMovingU {
		g.moveCount++
	}

	if g.rollCount > 9 {
		g.rollCount = 0
	}

	if g.moveCount > 7 {
		g.moveCount = 0
	}

	if g.jumpCount > g.jumpMax {
		g.jumpCount = 0
	} else {
		g.jumpCount++
	}

}

func (g *Game) spawn() {
	g.onDeck = nil
	maxLen := randNo(1, 5)

	for i := 0; i < maxLen; i++ {
		g.onDeck = append(g.onDeck, g.skullColors[randNo(0, 4)])
	}

}

func (g *Game) deepCopyPlanchette() {
	g.planchette = nil
	g.planchette = append(g.planchette, g.onDeck...)
}

func (g *Game) insertPlanchette() {
	row, col := 21, 8
	for i := len(g.planchette) - 1; i >= 0; i-- {
		if len(g.skullCollector[row][col]) > 0 {
			g.lostGame = true
			g.isPlaying = false
		}
		g.skullCollector[row][col] = g.planchette[i]
		row--
	}

	g.searchHead[0], g.searchHead[1] = row+1, col
}

func (g *Game) initSkullCollector() {
	g.skullCollector = make([][]string, rowMax)
	for row := 0; row < rowMax; row++ {
		cols := make([]string, colMax)
		for col := 0; col < colMax; col++ {
			cols[col] = ""
		}
		g.skullCollector[row] = cols
	}
}

func (g *Game) initSkullCoords() {
	g.skullCoords = make([][][]int, rowMax)

	for outer, y := 0, 60; y < (rowMax+1)*32; outer, y = outer+1, y+32 {
		g.skullCoords[outer] = make([][]int, colMax)
		for inner, x := 0, 22; x < colMax*32; inner, x = inner+1, x+32 {
			g.skullCoords[outer][inner] = make([]int, 2)
			g.skullCoords[outer][inner] = []int{x, y}
		}
	}
}

func (g *Game) initPlanchettes() {
	g.planchette = make([]string, 0)
	g.onDeck = make([]string, 0)
}

func (g *Game) reset() {
	g.havePlanchette = false
	g.isMovingL = false
	g.isMovingR = false
	g.isMovingU = false
	g.rollCount = 1
	g.moveCount = 0
	g.jumpCount = 1

	switch score := g.score; {

	case score > 99999:
		g.wonGame = true
	case score > 250:
		g.level = 7
		g.jumpMax = 14
	case score > 100:
		g.level = 6
		g.jumpMax = 16
	case score > 50:
		g.level = 5
		g.jumpMax = 18
	case score > 40:
		g.level = 4
		g.jumpMax = 20
	case score > 30:
		g.level = 3
		g.jumpMax = 25
	case score > 20:
		g.level = 2
		g.jumpMax = 30
	case score > 10:
		g.level = 1
		g.jumpMax = 35
	}

}

func (g *Game) resetGame() {
	g.score = 0
	g.reset()
	g.loseCount = 300
	g.lostGame = false
	g.isPlaying = false
	g.wonGame = false
}

func (g *Game) checkTrackPlaying() {
	if !g.track.IsPlaying() {
		g.track.Rewind()
		g.track.Play()
	}
}

/*  - - - - - - - -  D R A W   M E T H O D S  - - - - - - - -  */

func (g Game) drawBackground(screen *ebiten.Image) {
	opts := &ebiten.DrawImageOptions{}
	opts.GeoM.Translate(0, 0)
	screen.DrawImage(g.background, opts)
}

func (g Game) drawAllGameText(screen *ebiten.Image) {
	text.Draw(screen, fmt.Sprintf("%05d", g.score), g.fontFace, 528, 610, g.green)
	text.Draw(screen, fmt.Sprintf("%05d", g.best), g.fontFace, 528, 788, g.green)
	text.Draw(screen, fmt.Sprintf("%02d", g.level), g.fontFace, 556, 926, g.green)
}

func (g Game) drawOnDeck(screen *ebiten.Image) {
	var opts *ebiten.DrawImageOptions

	var odY float64

	switch len(g.onDeck) {
	case 4:
		odY = 186.00
	case 3:
		odY = 209.00
	case 2:
		odY = 222.00
	case 1:
		odY = 244.00
	}

	for _, skull := range g.onDeck {
		opts = &ebiten.DrawImageOptions{}
		opts.GeoM.Translate(561.00, odY)
		screen.DrawImage(g.skulls[skull], opts)
		odY += 32.00
	}
}

func (g Game) drawSkullCollector(screen *ebiten.Image) {

	var opts *ebiten.DrawImageOptions

	for i, row := range g.skullCollector {
		for j, col := range row {
			if len(col) > 0 {
				opts = &ebiten.DrawImageOptions{}
				opts.GeoM.Translate(float64(g.skullCoords[i][j][0]), float64(g.skullCoords[i][j][1]))
				screen.DrawImage(g.skulls[col], opts)
			}
		}
	}
}

func (g *Game) drawExplosions(screen *ebiten.Image) {
	var opts *ebiten.DrawImageOptions
	var width = frame / 2
	if g.explosionCount > 0 {
		for _, coords := range g.empties {
			opts = &ebiten.DrawImageOptions{}
			opts.GeoM.Translate(float64(g.skullCoords[coords[0]][coords[1]][0]), float64(g.skullCoords[coords[0]][coords[1]][1]))
			i := int(Clock/5) % 8
			sX, sY := i*width, 0
			screen.DrawImage(g.explosion.SubImage(image.Rect(sX, sY, sX+width, sY+width)).(*ebiten.Image), opts)
			g.explosionCount--
		}
	}
}

func (g Game) drawIntro(screen *ebiten.Image) {
	opts := &ebiten.DrawImageOptions{}
	opts.GeoM.Translate(0, 0)
	screen.DrawImage(g.intro, opts)

	if int(Clock)%50 < 40 {
		text.Draw(screen, "TOUCH TO BEGIN", g.fontFace, 200, 890, color.White)
	}
}

func (g *Game) drawLostGame(screen *ebiten.Image) {
	if g.loseCount > 0 {
		text.Draw(screen, "Game Over", g.fontFace, 200, 400, color.White)
		g.loseCount--
	}
}

func (g Game) drawWonGame(screen *ebiten.Image) {
	text.Draw(screen, " You Win ", g.fontFace, 200, 400, color.White)
}

/*  - - - - - - - -  E B I T E N   M E T H O D S  - - - - - - - -  */

// Update proceeds the game state every tick (1/60 [s] by default).
func (g *Game) Update() error {
	Clock++

	if g.isPlaying {
		if !g.havePlanchette {
			g.deepCopyPlanchette()
			g.insertPlanchette()
			g.spawn()
			g.havePlanchette = true
		}
		g.updatePlanchette()
	} else if g.lostGame && g.loseCount < 1 {
		g.resetGame()
		g.initSkullCollector()
		g.initPlanchettes()
		g.spawn()
	} else if g.wonGame {
		g.isPlaying = false
	} else {
		if inpututil.IsTouchJustReleased(0) {
			g.isPlaying = true
			Clock = 0
		}
	}

	return nil
}

// Layout takes the outside/window size and returns the (logical) screen size.
func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenW, screenH
}

// Draw the screen  every frame (typically 1/60[s] for 60Hz display).
func (g *Game) Draw(screen *ebiten.Image) {
	g.checkTrackPlaying()
	if g.lostGame {
		g.drawBackground(screen)
		g.drawSkullCollector(screen)
		g.drawAllGameText(screen)
		g.drawLostGame(screen)
	} else if g.wonGame {
		g.drawBackground(screen)
		g.drawAllGameText(screen)
		g.drawWonGame(screen)
	} else if !g.isPlaying {
		g.drawIntro(screen)
	} else {
		g.drawBackground(screen)
		g.drawOnDeck(screen)
		g.drawSkullCollector(screen)
		g.drawExplosions(screen)
		g.drawAllGameText(screen)
	}

	ebitenutil.DebugPrint(screen, "")
}

/*  - - - - - - - -  U T I L I T Y  F U N C T I O N S  - - - - - - - -  */

func randNo(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min) + min
}

func readRawIMG(asset []byte) (*ebiten.Image, error) {
	rawIMG, _, err := image.Decode(bytes.NewReader(asset))
	if err != nil {
		return nil, err
	}
	newImage := ebiten.NewImageFromImage(rawIMG)
	return newImage, nil
}

func readAudio(context *audio.Context, asset []byte) (*audio.Player, error) {
	mp3Decoded, err := mp3.Decode(context, bytes.NewReader(asset))
	if err != nil {
		return nil, err
	}
	player, err := audio.NewPlayer(context, mp3Decoded)
	if err != nil {
		return nil, err
	}
	return player, nil
}

// Load is the entry point to the game.
func Load() (*Game, error) {

	audioContext := audio.NewContext(44100)

	radioLand, err := truetype.Parse(assets.RadioLandTTF)
	if err != nil {
		return nil, err
	}

	PNGs := [][]byte{
		assets.BackgroundPNG,
		assets.ExplosionPNG,
		assets.IntroPNG,
		assets.GreenskullPNG,
		assets.RedskullPNG,
		assets.PurpleskullPNG,
		assets.BlueskullPNG,
	}

	var images []*ebiten.Image

	for _, png := range PNGs {
		ebImage, err := readRawIMG(png)
		if err != nil {
			return nil, err
		}
		images = append(images, ebImage)
	}

	beepSound, err := readAudio(audioContext, assets.BeepMP3)
	if err != nil {
		return nil, err
	}

	clearSound, err := readAudio(audioContext, assets.ClearMP3)
	if err != nil {
		return nil, err
	}

	themeSong, err := readAudio(audioContext, assets.ThemeMP3)
	if err != nil {
		return nil, err
	}

	var game = &Game{
		background:  images[0],
		explosion:   images[1],
		intro:       images[2],
		beep:        beepSound,
		clear:       clearSound,
		track:       themeSong,
		fontFace:    truetype.NewFace(radioLand, &truetype.Options{Size: 28, DPI: 72, Hinting: font.HintingFull}),
		green:       color.RGBA{R: 0x7C, G: 0xFC, B: 0x00, A: 0xFF},
		skullColors: []string{"purple", "blue", "red", "green"},
		jumpCount:   1,
		jumpMax:     35,
		matchMin:    3,
		loseCount:   300,
	}

	game.skulls = map[string]*ebiten.Image{
		"green":  images[3],
		"red":    images[4],
		"purple": images[5],
		"blue":   images[6],
	}

	game.initSkullCollector()
	game.initSkullCoords()
	game.initPlanchettes()

	game.track.SetVolume(.70)
	game.clear.SetVolume(.50)
	game.beep.SetVolume(.50)

	game.spawn()

	ebiten.SetWindowSize(screenW, screenH)
	ebiten.SetWindowTitle("💀")

	return game, nil
}
