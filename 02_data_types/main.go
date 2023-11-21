package main

import (
	"fmt"
	"reflect"
)

func main() {

	// объявление пременных
	var a int8 = 4
	var b float64 = 0.0
	var c byte

	// объявление констант
	const fl float64 = 0.01
	// const str string 	// ошибка, константам необходиомо присваивать знавечние
	const isEnabled bool = false
	const comp complex64 = 1 + 2i

	//Неявное объявление типов
	var name = "Svjatoslav" //string
	const new_a = 38        //int64
	var new_b = 0.0         //float64
	// var c  				//ошибка, переменной должно быть присвоено значение, либо объявлен ее тип

	//Объявление нескольких переменных
	var (
		my_name = "Sasha"
		age     = 21
	)
	var city, year = "SPB", 2021

	fmt.Println(reflect.TypeOf(a))
	fmt.Println(reflect.TypeOf(b))
	fmt.Println(reflect.TypeOf(c))
	fmt.Println(reflect.TypeOf(fl))
	fmt.Println(reflect.TypeOf(isEnabled))
	fmt.Println(reflect.TypeOf(comp))
	fmt.Println(reflect.TypeOf(name))
	fmt.Println(reflect.TypeOf(new_a))
	fmt.Println(reflect.TypeOf(new_b))

}
