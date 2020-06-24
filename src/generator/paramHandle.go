package generator

import (
	"codeGenerator/src/util"
	"log"
	"os"
	"text/template"
)

func Generator() {
	s := "f_user"
	table := SelectTable(s)
	tableName := TableName{}
	tableName.TableName = s
	tableName.PackageName = util.GetPackageName(s)
	tableName.ClassName = util.GetClassName(s)
	export := Export{Table: table, TableName: tableName}
	PojoGenerator(export)
	FormGenerator(export)
	DaoGenerator(export)
}

func SelectTable(s string) []Table {
	var table []Table
	queryString := "SELECT t.COLUMN_NAME as columnName, t.COLUMN_COMMENT as comment, t.DATA_TYPE as dataType , t.COLUMN_KEY as priKey FROM `COLUMNS` t WHERE t.TABLE_SCHEMA = 'chatroom' AND t.TABLE_NAME =\"" + s + "\""

	err := Mysql.Select(&table, queryString)
	if err != nil {
		log.Println(err.Error())
	}
	for k, v := range table {
		table[k].JsonName = util.GetJsonParam(v.ColumnName)
		table[k].PojoName = util.GetPojoParam(v.ColumnName)
	}
	return table
}
func PojoGenerator(export Export) {

	pojoTmpl, _ := template.ParseFiles("./template/pojoTemplate.txt")
	pojoFile, err := os.OpenFile("./user/"+export.TableName.PackageName+".go", os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		log.Println(err.Error())
	}
	err = pojoTmpl.Execute(pojoFile, export)
	if err != nil {
		log.Println(err.Error())
	}
}
func FormGenerator(export Export) {
	formTmpl, _ := template.ParseFiles("./template/formTemplate.txt")
	formFile, err := os.OpenFile("./user/"+export.TableName.PackageName+"Form.go", os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		log.Println(err.Error())
	}
	err = formTmpl.Execute(formFile, export)
	if err != nil {
		log.Println(err.Error())
	}
}
func DaoGenerator(export Export) {
	daoTmpl, _ := template.ParseFiles("./template/daoTemplate.txt")
	daoFile, err := os.OpenFile("./user/"+export.TableName.PackageName+"Dao.go", os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		log.Println(err.Error())
	}
	err = daoTmpl.Execute(daoFile, export)
	if err != nil {
		log.Println(err.Error())
	}
}
