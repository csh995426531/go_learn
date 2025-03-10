package test

import "errors"

func Division(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("分母不能为0")
	}
	return a / b, nil
}
