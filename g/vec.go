package g

import "image"

type V2 struct {
	X, Y float64
}

func V2FromPoint(p image.Point) V2 {
	return V2{
		X: float64(p.X),
		Y: float64(p.Y),
	}
}

func V2FromInt(x, y int) V2 {
	return V2{X: float64(x), Y: float64(y)}
}

func (a V2) Point() image.Point {
	return image.Point{
		X: int(a.X),
		Y: int(a.Y),
	}
}

func (a V2) Add(b V2) V2 {
	return V2{
		X: a.X + b.X,
		Y: a.Y + b.Y,
	}
}

func (a V2) Sub(b V2) V2 {
	return V2{
		X: a.X - b.X,
		Y: a.Y - b.Y,
	}
}

func (a V2) Mul(s V2) V2 {
	return V2{
		X: a.X * s.X,
		Y: a.Y * s.Y,
	}
}

func (a V2) Scale(s float64) V2 {
	return V2{
		X: a.X * s,
		Y: a.Y * s,
	}
}

func (a V2) XY() (x, y float64) {
	return a.X, a.Y
}

func Clamp(p, min, max float64) float64 {
	if p < min {
		p = min
	}
	if p > max {
		p = max
	}
	return p
}
