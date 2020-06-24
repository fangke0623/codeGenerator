package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"log"
	"os"
	"text/template"
	"time"
)

var mysql *sqlx.DB

func SqlInit() {
	db, err := sqlx.Open("mysql", "root:.Fangke123@tcp(39.108.145.221:3306)/information_schema?charset=utf8")
	if db == nil {
		return
	}
	if err != nil {
		log.Println("mysql connect error", err.Error())
	} else {
		log.Println("mysql connect success")
	}
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(10)
	db.SetConnMaxLifetime(60 * time.Second)
	mysql = db

}
func init() {
	SqlInit()
}

type Table struct {
	ColumnName string `db:"columnName"`
	Comment    string `db:"comment"`
	DataType   string `db:"dataType"`
	PriKey     string `db:"priKey"`
}
type Export struct {
	Table     []Table
	TableName string
}

func main() {
	var table []Table
	tableName := "discuss"
	queryString := "SELECT t.COLUMN_NAME as columnName, t.COLUMN_COMMENT as comment, t.DATA_TYPE as dataType , t.COLUMN_KEY as priKey FROM `COLUMNS` t WHERE t.TABLE_SCHEMA = 'chatroom' AND t.TABLE_NAME = 'd_discuss'"

	err := mysql.Select(&table, queryString)
	exp := Export{table, tableName}
	tmpl, _ := template.ParseFiles("./template/pojoTemplate.txt")
	f, err := os.OpenFile("./src/user/"+tableName+".go", os.O_RDWR|os.O_CREATE, os.ModePerm)

	if err != nil {
		log.Println(err.Error())
	}
	err = tmpl.Execute(f, exp)
	if err != nil {
		log.Println(err.Error())
	}
}
