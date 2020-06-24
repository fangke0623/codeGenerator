package generator

type Table struct {
	ColumnName string `db:"columnName"`
	Comment    string `db:"comment"`
	DataType   string `db:"dataType"`
	PriKey     string `db:"priKey"`
	JsonName   string
	PojoName   string
}
type TableName struct {
	TableName   string
	PackageName string
	ClassName   string
}

type Export struct {
	Table     []Table
	TableName TableName
}
