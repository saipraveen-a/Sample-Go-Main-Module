package main

import (
	"fmt"

	"github.com/saipraveen-a/number-manipulation/calc"
	calcNew "github.com/saipraveen-a/number-manipulation/v2/calc"
)

func main() {
	result := calc.Add(1, 2)
	fmt.Println("calc.Add(1,2) =>", result)

	result = calc.Add(1, 2, 3, 4, 5)
	fmt.Println("calc.Add(1,2,3,4,5) =>", result)

	newResult, err := calcNew.Add()

	if err != nil {
		fmt.Println("Error: =>", err)
	} else {
		fmt.Println("calcNew.Add(1,2,3,4) =>", newResult)
	}
}
