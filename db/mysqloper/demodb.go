package mysqloper

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type Demo struct {
	id         int64  `db:"id"`
	name       string `db:"name"`
	status     int    `db:"status"`
	createTime string `db:"create_time"`
}

func check(err error) { //因为要多次检查错误，所以干脆自己建立一个函数。
	if err != nil {
		fmt.Println(err)
	}

}

func SelectAll() {
	db, err := sql.Open("mysql", "root:zhh359#@tcp(127.0.0.1:3306)/easy-code")
	check(err)

	rows, err := db.Query("SELECT * FROM test")
	for rows.Next() {
		var s Demo
		err = rows.Scan(&s.id, &s.name, &s.status, &s.createTime)
		check(err)
		fmt.Println(s)
	}
	rows.Close()
}
