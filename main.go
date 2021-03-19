package main

import (
	"fmt"
	"image"
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"

	"github.com/egonelbre/expebiten/g"
	"github.com/egonelbre/expebiten/tilemap"
)

const (
	screenWidth  = 640
	screenHeight = 480
)

type World struct {
	Map *tilemap.Map

	Camera g.V2
	Zoom   float64
}

func NewWorld() *World {
	m := tilemap.NewMap(image.Pt(100, 100))
	for i := range m.Tiles.Pix {
		m.Tiles.Pix[i] = m.Sprites.Random()
	}

	return &World{
		Map:    m,
		Camera: g.V2{X: 50, Y: 50},
		Zoom:   1,
	}
}

func (world *World) Update() error {
	dt := 1.0 / ebiten.DefaultTPS // TODO: use proper timing
	dt *= 10

	if ebiten.IsKeyPressed(ebiten.KeyQ) {
		world.Zoom -= 0.05
	}
	if ebiten.IsKeyPressed(ebiten.KeyE) {
		world.Zoom += 0.05
	}

	world.Zoom = g.Clamp(world.Zoom, 0.1, 4)

	if ebiten.IsKeyPressed(ebiten.KeyW) {
		world.Camera.Y -= dt
	}
	if ebiten.IsKeyPressed(ebiten.KeyS) {
		world.Camera.Y += dt
	}
	if ebiten.IsKeyPressed(ebiten.KeyA) {
		world.Camera.X -= dt
	}
	if ebiten.IsKeyPressed(ebiten.KeyD) {
		world.Camera.X += dt
	}
	return nil
}

func (world *World) Draw(screen *ebiten.Image) {
	world.Map.Render(screen, world.Camera, world.Zoom)
	ebitenutil.DebugPrint(screen, fmt.Sprintf("TPS: %0.2f", ebiten.CurrentTPS()))
}

func (world *World) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	world := NewWorld()
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("TileMap")
	if err := ebiten.RunGame(world); err != nil {
		log.Fatal(err)
	}
}
