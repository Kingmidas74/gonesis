package util

func ModLikePython(number, module int) int {
	res := number % module
	if (res < 0 && module > 0) || (res > 0 && module < 0) {
		return res + module
	}
	return res
}
