package errors

import "errors"

var ErrMazeSizeIncorrect = errors.New("can't create maze")
var ErrFreeRequirementIncorrect = errors.New("size of maze is too small")
var MAZE_GENERATOR_UNDEFINED = errors.New("maze generator is undefined")
