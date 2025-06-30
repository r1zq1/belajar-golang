package calculator

import "errors"

func Add(num1 int, num2 int) int {
	return num1 + num2
}

func Substract(num1, num2 int) int {
	return num1 - num2
}

func Divide(num1, num2 float64) (float64, error) {
	if num2 == 0.0 {
		return 0, errors.New("pembagian dengan nol tidak diizinkan")
	}
	return num1 / num2, nil
}

func AddSub(num1, num2 int) (add, sub int) {
	add = num1 + num2
	sub = num1 - num2
	return
}
