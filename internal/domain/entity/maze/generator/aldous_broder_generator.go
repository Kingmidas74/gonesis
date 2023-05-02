package generator

import (
	"github.com/kingmidas74/gonesis-engine/internal/domain/errors"
	"github.com/kingmidas74/gonesis-engine/internal/util"
)

type AldousBroderGenerator struct {
	gridGenerator GridGenerator
}

func (g AldousBroderGenerator) Generate(width, height int) (maze []bool, err error) {
	if width <= 0 || height <= 0 {
		return make([]bool, 0), errors.MAZE_SIZE_INCORRECT
	}

	maze, err = g.gridGenerator.Generate(width, height)
	if err != nil {
		return nil, err
	}

	for i := range maze {
		maze[i] = false
	}

	actors := make([]int, 0)

	actorsCount := 7

	for i := 0; i < actorsCount; i++ {
		actors = append(actors, 0)
		actors = append(actors, 0)
	}

	directions := make(map[int][4]int)
	directions[0] = [4]int{0, -2, 0, 1}
	directions[1] = [4]int{2, 0, -1, 0}
	directions[2] = [4]int{0, 2, 0, -1}
	directions[3] = [4]int{-2, 0, 1, 0}

	isFinished := false

	for !isFinished {
		for i := 0; i < actorsCount; i++ {

			direction := directions[util.RandomIntBetween(0, 3)]

			if actors[i*2]+direction[0] >= 0 && actors[i*2]+direction[0] < width && actors[i*2+1]+direction[1] >= 0 && actors[i*2+1]+direction[1] < height {

				actors[i*2] += direction[0]
				actors[i*2+1] += direction[1]

				if maze[actors[i*2+1]*width+actors[i*2]] == false {
					maze[actors[i*2+1]*width+actors[i*2]] = true
					maze[(actors[i*2+1]+direction[3])*width+actors[i*2]+direction[2]] = true
				}
			}
		}

		isFinished = true
		for y := 0; y < height; y = y + 2 {
			for x := 0; x < width; x = x + 2 {
				isFinished = isFinished && maze[y*width+x]
			}
		}
	}
	maze[0] = true

	return maze, nil
}
