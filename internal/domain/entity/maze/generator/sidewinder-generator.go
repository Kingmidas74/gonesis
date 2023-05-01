package generator

import "github.com/kingmidas74/gonesis-engine/internal/util"

type SidewinderGenerator struct {
}

func (g SidewinderGenerator) Generate(width, height int, matrix []bool) []bool {

	for y := 0; y < height; y = y + 2 {
		for x := 0; x < width; x = x + 2 {

			if y == 0 {
				if x+1 < width {
					matrix[y*width+x+1] = true
				}
				continue
			}

			direction := util.RandomIntBetween(0, 1)
			if direction == 0 {
				matrix[(y-1)*width+x] = true
				continue
			}

			if x+1 >= width {
				continue
			}

			matrix[y*width+x+1] = true
		}
	}

	return matrix
}
