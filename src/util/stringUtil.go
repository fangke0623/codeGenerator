package util

import (
	"fmt"
	"strings"
)

func GetJsonParam(s string) string {
	temp := strings.Split(s, "_")
	var upperStr string
	for y := 0; y < len(temp); y++ {
		vv := []rune(temp[y])
		if y != 0 {
			for i := 0; i < len(vv); i++ {
				if i == 0 {
					vv[i] -= 32               // string的码表相差32位
					upperStr += string(vv[i]) // + string(vv[i+1])
				} else {
					upperStr += string(vv[i])
				}
			}
		}
	}
	return temp[0] + upperStr
}
func GetPackageName(s string) string {
	temp := strings.Split(s, "_")
	if len(temp) > 1 {
		return temp[1]
	} else {
		return temp[0]
	}
}
func GetClassName(s string) string {
	return Capitalize(GetPackageName(s))
}

func GetPojoParam(s string) string {
	return Capitalize(GetJsonParam(s))
}
func Capitalize(str string) string {
	var upperStr string
	vv := []rune(str)
	for i := 0; i < len(vv); i++ {
		if i == 0 {
			if vv[i] >= 97 && vv[i] <= 122 {
				vv[i] -= 32 // string的码表相差32位
				upperStr += string(vv[i])
			} else {
				fmt.Println("Not begins with lowercase letter,")
				return str
			}
		} else {
			upperStr += string(vv[i])
		}
	}
	return upperStr
}
