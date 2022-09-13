package methods

import (
	"errors"
	"fmt"

	"github.com/Knetic/govaluate"
)

func Bisection(f string, l, r, tol, n float64) error {
	exp, err := govaluate.NewEvaluableExpression(f)
	if err != nil {
		return errors.New("Expression mal formed")
	}
	params := make(map[string]interface{}, 8)
	params["x"] = l
	res, err := exp.Evaluate(params)
	if err != nil {
		return errors.New("Expression mal formed")
	}
	fmt.Println(res)
	return nil
}
