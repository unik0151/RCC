package pck1

import (
	"fmt"
	_ "githup/unik0105/init-mod/pck2"
)

func init() {
	fmt.Println("pck1 has bean initialized -- ")
}

const param1 string = "pck1"

var param2Var = getPck1Name()

func getPck1Name() string {
	fmt.Println("getPck1Name ---- ")
	return "pck1"
}
