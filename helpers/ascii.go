package helpers

import (
	"fmt"
	"strconv"
)


func SliceToInt(arg interface{}) []int {
	var val []int
	switch arg.(type) {
	case []string:
		for _, str := range arg.([]string) {
			num, err := strconv.Atoi(str)
			if err != nil {
				panic("error converting string to int" + err.Error())
			}
			val = append(val, num)
		}
	default:
	  panic(fmt.Sprintf("unhandled type %T", arg))
	}
	return val
}


func ToInt(arg interface{}) int {
	var val int
	switch arg.(type) {
	case string:
		var err error
		val, err = strconv.Atoi(arg.(string))
		if err != nil {
			panic("error converting string to int" + err.Error())
		}
	case rune:
		val = int(arg.(rune))
	default:
	  panic(fmt.Sprintf("unhandled type %T", arg))
	}
	return val
}

func ToString(arg interface{}) string {
	var str string
	switch arg.(type) {
	case int:
		str = strconv.Itoa(arg.(int))
	case byte:
		b := arg.(byte)
		str = string(rune(b))
	case rune:
		str = string(arg.(rune))
	default:
	  panic(fmt.Sprintf("unhandled type %T", arg))
	}
	return str
}

const (
	ASCIICodeA = int('A')
	ASCIICodeZ = int('Z')
	ASCIICodea = int('a')
	ASCIICodez = int('z')
)

func ToASCIICode(arg interface{}) int{
	var asciiVal int
	switch arg.(type) {
	case string:
		str := arg.(string)
		if len(str) != 1 {
			panic("string must be length 1")
		}
		asciiVal = int(str[0])
	case byte:
		asciiVal = int(arg.(byte))
	case rune:
		asciiVal = int(arg.(rune))
	}
	return asciiVal
}

func ASCIIIntToChar(arg int) string {
	return string(rune(arg))
}
