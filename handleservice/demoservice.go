package handleservice

import (
	"../db/mysqloper"
	"../db/redisoper"
)

func SelectAll() {
	mysqloper.SelectAll()
	redisoper.SetDemoCache()
}
