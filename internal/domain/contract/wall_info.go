package contract

type WallInfo interface {
	WestWall() bool
	NorthWall() bool
	EastWall() bool
	SouthWall() bool
}
