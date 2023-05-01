package maze

import (
	"math/rand"
	"testing"
	"time"

	"github.com/kingmidas74/gonesis-engine/internal/domain/entity/maze/generator"
)

func TestMazeService_Generate(t *testing.T) {
	rand.Seed(time.Now().UnixNano())

	s := NewMazeService[generator.AldousBroderGenerator]()
	res, err := s.Generate(5, 5)
	if err != nil {
		t.Error(err)
	}
	if res.GetWidth() != 5 || res.GetHeight() != 5 {
		t.Error("wrong size")
	}
}
