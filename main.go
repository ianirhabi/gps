package main

import (
	orm "github.com/alfatih/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"gps.com/gps/engine"
)

func init() {
	orm.RegisterDataBase("default", "mysql", "root:ian!@#$%^@tcp(localhost:3306)/gps_app?charset=utf8")
}

func main() {
	engine.Router()
}
