package errors

import "errors"

var MAZE_SIZE_INCORRECT = errors.New("can't create maze")
var MAZE_GENERATOR_UNDEFINED = errors.New("maze generator is undefined")
