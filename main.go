package main

import (
	"fmt"
	"unsafe"

	"demoserver/commlib"
	"demoserver/handleservice"
)

func main() {
	//test.Funtest()
	handleservice.SelectAll()
	var str = "this is a start func"
	fmt.Println(str, " ==> size = ", unsafe.Sizeof(str))
	var sum = commlib.Addnumber(1, 2)
	fmt.Println("a+b= ", sum)
}
