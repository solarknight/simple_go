package main

import (
	"fmt"
	"reflect"
)

func tryReflect() {
	fmt.Println()
	var x = 3.4
	fmt.Println("type:", reflect.TypeOf(x))
	v := reflect.ValueOf(&x)
	fmt.Println("value:", v)
	fmt.Println("type:", v.Type())
	fmt.Println("kind:", v.Kind())
	// get element by pointer
	v = v.Elem()
	fmt.Println("can set value:", v.CanSet())
	fmt.Println(v.Interface())
	fmt.Printf("value is %5.2e\n", v.Interface())
	y, ok := v.Interface().(float64)
	fmt.Println("ok:", ok, "y:", y)

	// try use reflect on struct
	reflectOnStruct()
}

type NotknownType struct {
	s1, s2, s3 string
}

func (n NotknownType) String() string {
	return n.s1 + " - " + n.s2 + " - " + n.s3
}

// variable to investigate:
var secret interface{} = NotknownType{"Ada", "Go", "Oberon"}

func reflectOnStruct() {
	fmt.Println(secret)
	// alternative:
	//typ := value.Type()  // main.NotknownType
	value := reflect.ValueOf(secret)
	typ := value.Type()
	fmt.Println(typ)
	knd := value.Kind() // struct
	fmt.Println(knd)

	// iterate through the fields of the struct:
	for i := 0; i < value.NumField(); i++ {
		fmt.Printf("Field %d: %v\n", i, value.Field(i))
		// error: panic: reflect.Value.SetString using value obtained using unexported field
		//value.Field(i).SetString("C#")
	}

	// call the first method, which is String():
	results := value.Method(0).Call(nil)
	fmt.Println(results) // [Ada - Go - Oberon]
}
