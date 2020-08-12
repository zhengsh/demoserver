package handleservice

import (
	"demoserver/db/mysqloper"
	"demoserver/db/redisoper"
	"fmt"
)

func SelectAll() {
	mysqloper.SelectAll()
	redisoper.SetDemoCache()
	fmt.Println("select all exec")
}
