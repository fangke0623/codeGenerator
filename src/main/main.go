package main

import (
	"codeGenerator/src/generator"
)

func init() {
	generator.SqlInit()
}

func main() {
	generator.Generator()
}
