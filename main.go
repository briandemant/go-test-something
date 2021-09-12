package main

import "errors"

func Add2(input int) (result int) {
	result = input + 2
	return
}

func Add4(input int) (result int) {
	return input + 4
}

func Add(x int, y int) (result int) {
	return x + y
}

var ErrFail = errors.New("fail!")

func BetterError(ok bool) (result string, err error) {
	if ok {
		return "ok", nil
	} else {
		return "fail", ErrFail
	}
}
