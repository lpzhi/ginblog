package models

import "fmt"


type dbshell struct {
	Model
}

func CreateRoleTotalDatabase() {

	//sql := "create database role_total charset utf8"
	////sql := "select * from blog_tag"
	//rs := db.Exec(sql)
	sq := "use role_total"
	rs := db.Exec(sq)
	fmt.Println(rs)
	//strTable := "Create table t_agent(" +
	//	"`pfrom_id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '代理商id'," +
	//	"  PRIMARY KEY (`pfrom_id`)"+
	//	")"
	//rs := db.Exec(strTable)
	//fmt.Println(rs)
}
