package errors

import "errors"

var ErrMazeSizeIncorrect = errors.New("can't create maze")
var ErrFreeRequirementIncorrect = errors.New("size of maze is too small")
var ErrMazeTypeNotSupported = errors.New("maze generator is undefined")
var ErrTopologyTypeNotSupported = errors.New("topology is undefined")
var ErrReproductionSystemTypeNotSupported = errors.New("reproduction system is undefined")
