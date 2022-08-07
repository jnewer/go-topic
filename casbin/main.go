package main

import (
	"fmt"
	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v2"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	a, _ := gormadapter.NewAdapter("mysql", "root:@tcp(127.0.0.1:3306)/test_db", true)
	e, _ := casbin.NewEnforcer("rbac_model.conf", a)

	sub := "ghz"
	obj := "data1"
	act := "read"

	ok, _ := e.Enforce(sub, obj, act)

	if ok {
		fmt.Println("yes")
	} else {
		fmt.Println("no")
	}
}
