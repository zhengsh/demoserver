package redisoper

import "fmt"

func SetDemoCache() {
	set("abc", "124")
	value := getStringValue("abc")
	fmt.Println(value)
}
