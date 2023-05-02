package generator

import "testing"

func TestGridGenerator_Generate(t *testing.T) {
	sut := GridGenerator{}

	maze, err := sut.Generate(10, 5)
	if err != nil {
		t.Errorf("unexpected error")
	}

	if len(maze) != 10*5 {
		t.Errorf("size doesn't match")
	}

	for y := 0; y < 5; y++ {
		for x := 0; x < 10; x++ {
			if y%2 == 0 {
				if x%2 == 0 {
					if maze[y*10+x] == false {
						t.Errorf("wrong maze")
					}
				} else {
					if maze[y*10+x] == true {
						t.Errorf("wrong maze")
					}
				}
			} else {
				if maze[y*10+x] == true {
					t.Errorf("wrong maze")
				}
			}
		}
	}
}
