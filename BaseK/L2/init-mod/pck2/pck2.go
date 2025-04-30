package pck2

import (
	"fmt"
)

func init() {
	fmt.Println("pck2 has bean initialized -- ")
}

const param2 string = "pck2"

var param2Var = getPck2Name()

func getPck2Name() string {

	return "pck2"
}
