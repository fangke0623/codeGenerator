package main

import (
	"archive/zip"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

func init() {
	//generator.SqlInit()
}

func main() {
	//DeviceTypeToPid(1)
	err := Zip("./src/generator/", "000000000000000000000000f45f0100.zip")
	if err != nil {
		fmt.Printf("%s", err)

	}
}

// srcFile could be a single file or a directory
func Zip(srcFile string, destZip string) error {
	zipfile, err := os.Create(destZip)
	if err != nil {
		return err
	}
	defer zipfile.Close()

	archive := zip.NewWriter(zipfile)
	defer archive.Close()

	filepath.Walk(srcFile, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Printf("%s", err)
			return err
		}

		header, err := zip.FileInfoHeader(info)
		if err != nil {
			return err
		}

		header.Name = strings.TrimPrefix(path, filepath.Dir(srcFile)+"/")
		// header.Name = path
		if info.IsDir() {
			header.Name += "/"
		} else {
			header.Method = zip.Deflate
		}

		writer, err := archive.CreateHeader(header)
		if err != nil {
			return err
		}

		if !info.IsDir() {
			file, err := os.Open(path)
			if err != nil {
				return err
			}
			defer file.Close()
			_, err = io.Copy(writer, file)
		}
		return err
	})

	return err
}

func DeviceTypeToPid(deviceType int) {
	pack := [16]byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	buf := make([]byte, 4)
	binary.LittleEndian.PutUint32(buf, uint32(deviceType))
	copy(pack[12:], buf)

	fmt.Printf("%x", pack)
}

func Generator() {
	//s := "d_discuss"
	//table := SelectTable(s)
	//tableName := TableName{}
	//tableName.TableName = s
	//tableName.PackageName = util.GetPackageName(s)
	//tableName.ClassName = util.GetClassName(s)
	//export := Export{Table: table, TableName: tableName}
	export := Export{}
	//PojoGenerator(export)
	//FormGenerator(export)
	//DaoGenerator(export)
	VtGenerator(export)
}

type Export struct {
	Vt Vt
}
type AttributesExtend struct {
	DefaultValue interface{} `json:"defaultValue"`
	IsNull       int         `json:"isNull"`
	OnlyRead     int         `json:"onlyRead"`
	IsShow       int         `json:"isShow"`
	IsMust       int         `json:"isMust"`
	DefaultShow  int         `json:"defalutShow"`
	IsChoosed    int         `json:"isChosed"`
}
type (
	Vt struct {
		Name         string    `json:"name"`
		Pid          string    `json:"pid"`
		ModelRelated []VtModel `json:"modelRelated"`
	}

	VtModel struct {
		AttributesRelated []VtField `json:"attributesRelated"`
	}
	VtField struct {
		Name         string      `json:"name"`
		Key          string      `json:"key"`
		FieldType    int         `json:"fieldType"`
		FieldExpand  interface{} `json:"fieldExpand"`
		EnumExpand   []FieldEnum
		StringExpand FieldString
		NumExpand    FieldNum
		NumSize      int
		Extend       AttributesExtend `json:"attributesExtend"`
	}
	FieldEnum struct {
		EnumKey   int
		EnumValue string
	}
	FieldString struct {
		Length     int    `json:"length"`
		StringType string `json:"stringType"`
	}
	FieldNum struct {
		Multiple int    `json:"multiple"`
		Step     int    `json:"step"`
		Unit     string `json:"unit"`
		ValueMax int    `json:"value_max"`
		ValueMin int    `json:"value_min"`
	}
)

func JsonToObject(jsonObject interface{}, object interface{}) {

	resByre, err := json.Marshal(jsonObject)
	if err != nil {
		println("Object to json error:", err)
	}
	err = json.Unmarshal(resByre, &object)

	//err := json.Unmarshal(jsonObject, &object)
	if err != nil {
		println("Object to json error:", err)
	}
}

type Response struct {
	Status int         `json:"status"`
	Data   interface{} `json:"data"`
}

func VtGenerator(export Export) {
	var response Response
	res, _ := ioutil.ReadFile("./template/data.json")
	_ = json.Unmarshal(res, &response)

	JsonToObject(response.Data, &export.Vt)
	vtHandle(&export)
	pojoTmpl, err := template.ParseFiles("./template/vtTemplate.txt")
	if err != nil {
		log.Println(err.Error())
	}
	_ = os.Mkdir("./generator/", os.ModePerm)

	pojoFile, err := os.OpenFile("./generator/"+export.Vt.Pid+".json", os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		log.Println(err.Error())
	}
	err = pojoTmpl.Execute(pojoFile, export.Vt)
	if err != nil {
		log.Println(err.Error())
	}
}

func vtHandle(export *Export) {
	var models []VtModel
	for _, m := range export.Vt.ModelRelated {
		var fields []VtField
		for _, a := range m.AttributesRelated {
			if a.FieldType == 0 {
				JsonToObject(a.FieldExpand, &a.EnumExpand)
				a.NumSize = len(a.EnumExpand) - 1
			}
			if a.FieldType == 1 {
				JsonToObject(a.FieldExpand, &a.NumExpand)

			}
			if a.FieldType == 2 {
				JsonToObject(a.FieldExpand, &a.StringExpand)

			}
			fields = append(fields, a)
		}
		m.AttributesRelated = fields
		models = append(models, m)
	}
	export.Vt.ModelRelated = models
}
