package main

//package generator
//
//import (
//	_ "github.com/go-sql-driver/mysql"
//	"github.com/jmoiron/sqlx"
//	"log"
//	"time"
//)
//
//var Mysql *sqlx.DB
//
//func SqlInit() {
//	db, err := sqlx.Open("mysql", "root:.Fangke123@tcp(39.108.145.221:3306)/information_schema?charset=utf8")
//	if db == nil {
//		return
//	}
//	if err != nil {
//		log.Println("mysql connect error", err.Error())
//	} else {
//		log.Println("mysql connect success")
//	}
//	db.SetMaxIdleConns(10)
//	db.SetMaxOpenConns(10)
//	db.SetConnMaxLifetime(60 * time.Second)
//	Mysql = db
//
//}
