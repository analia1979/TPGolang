package model

import (
	"strconv"
)

type Result struct {
	Type   string
	Length int
	Value  string
}

func NewResult(t string, l int, v string) Result {
	return Result{t, l, v}
}

func containsCharValid(c string) bool {

	for _, v := range c {
		if !(v >= 'a' && v <= 'z' || v >= 'A' && v <= 'Z') {

			return false
		}

	}
	return true

}

func ConverterTo(palabra string) Result {
	//la long de la palabra minimo tiene que tener seis caracteres AA03AAA
	//y no contener caracteres especiales
	var r Result
	if len(palabra) >= 4 {
		tipo := palabra[:2]
		if containsCharValid(tipo) {
			longitud := palabra[2:4]
			sv, error := strconv.Atoi(longitud)
			//fmt.Println(error, sv)
			if error == nil {
				valor := palabra[4:]
				if containsCharValid(valor) {
					cant := len(valor)
					if sv == cant {
						r.Type = tipo
						r.Length = sv
						r.Value = valor

					} else {
						panic("La longitud Value no corresponde con el Length ingresado")
					}
				} else {
					panic("Value invalido---Contine caracteres invalidos")
				}

			} else {
				panic("Length incorrecta-Deben ingresar valores numericos")
			}
		} else {
			panic("Type incorrecto--Contiene caracteres invalidos")
		}
	}
	return r
}
