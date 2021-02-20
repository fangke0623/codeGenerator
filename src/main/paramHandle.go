package main

//func Generator() {
//	//s := "d_discuss"
//	//table := SelectTable(s)
//	//tableName := TableName{}
//	//tableName.TableName = s
//	//tableName.PackageName = util.GetPackageName(s)
//	//tableName.ClassName = util.GetClassName(s)
//	//export := Export{Table: table, TableName: tableName}
//	export := Export{}
//	//PojoGenerator(export)
//	//FormGenerator(export)
//	//DaoGenerator(export)
//	VtGenerator(export)
//}
//func JsonToObject(jsonObject interface{}, object interface{}) {
//
//	resByre, err := json.Marshal(jsonObject)
//	if err != nil {
//		println("Object to json error:", err)
//	}
//	err = json.Unmarshal(resByre, &object)
//
//	//err := json.Unmarshal(jsonObject, &object)
//	if err != nil {
//		println("Object to json error:", err)
//	}
//}
//type Response struct {
//	Status int `json:"status"`
//	Data interface{} `json:"data"`
// }
//func VtGenerator(export Export) {
//	var response Response
//	res,_ := ioutil.ReadFile("./template/data.json")
//	_ = json.Unmarshal(res,&response)
//
//	JsonToObject(response.Data,&export.Vt)
//
//	pojoTmpl, err := template.ParseFiles("./template/vtTemplate.txt")
//	if err != nil {
//		log.Println(err.Error())
//	}
//	_ = os.Mkdir("./generator/", os.ModePerm)
//
//	pojoFile, err := os.OpenFile("./generator/"+export.Vt.Pid+".json", os.O_RDWR|os.O_CREATE, os.ModePerm)
//	if err != nil {
//		log.Println(err.Error())
//	}
//	err = pojoTmpl.Execute(pojoFile, export.Vt)
//	if err != nil {
//		log.Println(err.Error())
//	}
//}
//func SelectTable(s string) []Table {
//	var table []Table
//	queryString := "SELECT t.COLUMN_NAME as columnName, t.COLUMN_COMMENT as comment, t.DATA_TYPE as dataType , t.COLUMN_KEY as priKey FROM `COLUMNS` t WHERE t.TABLE_SCHEMA = 'chatroom' AND t.TABLE_NAME =\"" + s + "\""
//
//	err := Mysql.Select(&table, queryString)
//	if err != nil {
//		log.Println(err.Error())
//	}
//	for k, v := range table {
//		table[k].JsonName = util.GetJsonParam(v.ColumnName)
//		table[k].PojoName = util.GetPojoParam(v.ColumnName)
//	}
//	return table
//}
//func PojoGenerator(export Export) {
//
//	pojoTmpl, _ := template.ParseFiles("./template/pojoTemplate.txt")
//	_ = os.Mkdir("./generator/", os.ModePerm)
//	err := os.Mkdir("./generator/"+export.TableName.PackageName, os.ModePerm)
//	if err != nil {
//		log.Println(err.Error())
//	}
//	pojoFile, err := os.OpenFile("./generator/"+export.TableName.PackageName+"/"+export.TableName.PackageName+".go", os.O_RDWR|os.O_CREATE, os.ModePerm)
//	if err != nil {
//		log.Println(err.Error())
//	}
//	err = pojoTmpl.Execute(pojoFile, export)
//	if err != nil {
//		log.Println(err.Error())
//	}
//}
//func FormGenerator(export Export) {
//	formTmpl, _ := template.ParseFiles("./template/formTemplate.txt")
//
//	formFile, err := os.OpenFile("./generator/"+export.TableName.PackageName+"/"+export.TableName.PackageName+"Form.go", os.O_RDWR|os.O_CREATE, os.ModePerm)
//	if err != nil {
//		log.Println(err.Error())
//	}
//	err = formTmpl.Execute(formFile, export)
//	if err != nil {
//		log.Println(err.Error())
//	}
//}
//func DaoGenerator(export Export) {
//	daoTmpl, _ := template.ParseFiles("./template/daoTemplate.txt")
//	daoFile, err := os.OpenFile("./user/"+export.TableName.PackageName+"Dao.go", os.O_RDWR|os.O_CREATE, os.ModePerm)
//	if err != nil {
//		log.Println(err.Error())
//	}
//	err = daoTmpl.Execute(daoFile, export)
//	if err != nil {
//		log.Println(err.Error())
//	}
//}
