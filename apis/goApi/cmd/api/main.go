package main

import (
	"fmt"

	"github.com/ksossa/NumericalMethods/apis/goApi/internal/methods"
)

func main() {
	err := methods.Bisection("x**3", 3, 6, 0.0001, 10)
	if err != nil {
		fmt.Println(err)
	}
}
