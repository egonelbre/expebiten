package tilemap

import (
	"bytes"
	_ "embed"
	"image"
	_ "image/png"
	"log"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"

	"github.com/egonelbre/expebiten/g"
)

//go:embed sheet.png
var sheetData []byte

const (
	tileSize   = 16
	tileMargin = 1
)

type Sprites struct {
	*ebiten.Image
}

func NewSprites() *Sprites {
	img, _, err := image.Decode(bytes.NewReader(sheetData))
	if err != nil {
		log.Fatal(err)
	}
	return &Sprites{
		Image: ebiten.NewImageFromImage(img),
	}
}

func (s *Sprites) Size() g.V2 { return g.V2{X: tileSize, Y: tileSize} }

func (s *Sprites) Random() uint8 {
	w, h := s.Image.Size()
	w /= (tileSize + tileMargin)
	h /= (tileSize + tileMargin)

	x, y := rand.Intn(w), rand.Intn(h)
	return uint8((x << 4) | y)
}
func (s *Sprites) Sprite(p uint8) *ebiten.Image {
	x := (tileSize + tileMargin) * int((p>>4)&(1<<4-1))
	y := (tileSize + tileMargin) * int((p>>0)&(1<<4-1))
	return s.Image.SubImage(image.Rect(x, y, x+tileSize, y+tileSize)).(*ebiten.Image)
}

type Map struct {
	Sprites *Sprites
	Tiles   *image.Gray
}

func NewMap(size image.Point) *Map {
	return &Map{
		Sprites: NewSprites(),
		Tiles:   image.NewGray(image.Rectangle{Max: size}),
	}
}

func (m *Map) Render(screen *ebiten.Image, worldCenter g.V2, tileScale float64) {
	tileSize := m.Sprites.Size().Scale(tileScale)
	screenSize := g.V2FromInt(screen.Size())

	// min, max can probably be calculated more precisely ... rendering more tiles for now

	tileCount := tileCoord(screenSize, tileSize).Add(image.Point{X: 2, Y: 2})

	center := worldCenter.Point()
	min := center.Sub(tileCount.Div(2))
	max := center.Add(tileCount.Div(2))

	minTile := maxPoint(min, m.Tiles.Rect.Min)
	maxTile := minPoint(max, m.Tiles.Rect.Max)

	op := &ebiten.DrawImageOptions{}

	screenCenter := screenSize.Scale(0.5)
	for y := minTile.Y; y <= maxTile.Y; y++ {
		for x := minTile.X; x <= maxTile.X; x++ {
			topLeft := g.V2FromPoint(image.Point{X: x, Y: y}).Sub(worldCenter).Mul(tileSize).Add(screenCenter)

			op.GeoM = ebiten.GeoM{}
			op.GeoM.Scale(tileScale, tileScale)
			op.GeoM.Translate(topLeft.XY())

			tile := m.Tiles.GrayAt(x, y)
			screen.DrawImage(m.Sprites.Sprite(tile.Y), op)
		}
	}
}

func tileCoord(p, tileSize g.V2) image.Point {
	return image.Point{
		X: int(p.X / tileSize.X),
		Y: int(p.Y / tileSize.Y),
	}
}

func minPoint(a, b image.Point) image.Point {
	if a.X > b.X {
		a.X = b.X
	}
	if a.Y > b.Y {
		a.Y = b.Y
	}
	return a
}

func maxPoint(a, b image.Point) image.Point {
	if a.X < b.X {
		a.X = b.X
	}
	if a.Y < b.Y {
		a.Y = b.Y
	}
	return a
}
