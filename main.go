package main

import (
	"fmt"
	"modulos/model"
)

func main() {

	palabra := "ab04AAAA"

	r := model.ConverterTo(palabra)
	fmt.Println(r.Type, r.Length, r.Value)
	//	fmt.Println(model.ContainsChar("AAAAA&BBBBA"))

}
