package main

import (
	"fmt"
	_ "githup/unik0105/init-mod/pck1"
)

const mainVar1 string = "main"

var var2 string = var2Func()

var var3 = func() string {
	return "sss"
}

func main() {

	fmt.Println("main exe success !")
}

func var2Func() string {
	fmt.Println("var2Func ----- ")
	return "main2"
}
