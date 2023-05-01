package entity

type Coords struct {
	X int
	Y int
}

func (this *Coords) GetX() int {
	return this.X
}

func (this *Coords) SetX(x int) {
	this.X = x
}

func (this *Coords) GetY() int {
	return this.Y
}

func (this *Coords) SetY(y int) {
	this.Y = y
}
