package main

import (
	"fmt"
	"image/color"
	"log"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const (
	scale = 2
	w     = 1024
	h     = 768
	halfW = w / 2
	halfH = h / 2
	halfS = scale / 2
)

var (
	img  *ebiten.Image
	game Game
)

type Thing struct {
	X float64
	Y float64
}

type Game struct {
	time float64
}

func Init() {
	game = Game{0.0}
	img = ebiten.NewImage(scale, scale)
	img.Fill(color.RGBA{0xff, 0xff, 0xff, 0xff})
}

func (t *Thing) draw(screen *ebiten.Image, theta float64) {
	op := &ebiten.DrawImageOptions{}
	// Translate the thing position (y axis is negative to invert the coordinates)
	op.GeoM.Translate(t.X, -t.Y)
	op.GeoM.Rotate(theta)
	op.GeoM.Translate(halfW - halfS, halfH - halfS)
	screen.DrawImage(img, op)
}

func (g *Game) Update() error {
	g.time += .02
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	theta := math.Sin(g.time*60.0*0.001) * math.Pi * 2.0 * 4.0
	for y := 0; y < h; y++ {
		y := float64(y)
		for i := 0; i < 10; i++ {
			i := float64(i)
			rot := i + 1.0 + y * 10.0
			thing := Thing{math.Sin(y*0.1+theta+i*2.0) * 100., y}
			thing.draw(screen, theta*0.001*rot)
		}
	}
	msg := fmt.Sprintf(`TPS: %0.2f FPS: %0.2f`, ebiten.CurrentTPS(), ebiten.CurrentFPS())
	ebitenutil.DebugPrint(screen, msg)

}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return w, h
}

func main() {
	ebiten.SetWindowSize(w, h)
	ebiten.SetWindowTitle("Spiral")
	Init()
	if err := ebiten.RunGame(&game); err != nil {
		log.Fatal(err)
	}
}
