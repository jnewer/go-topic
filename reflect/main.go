package main

import (
	"fmt"
	"reflect"
)

func say(text string) {
	fmt.Println(text)
}
func Call(m map[string]interface{}, name string, params ...interface{}) (result []reflect.Value) {
	f := reflect.ValueOf(m[name])
	in := make([]reflect.Value, len(params))
	for k, param := range params {
		in[k] = reflect.ValueOf(param)
	}
	result = f.Call(in)

	return
}
func main() {
	var x int = 1
	fmt.Println("type: ", reflect.TypeOf(x))                              //type:  int
	fmt.Println("value: ", reflect.ValueOf(x))                            //value:  1
	fmt.Println("Kind: ", reflect.TypeOf(x).Kind())                       //Kind:  int
	fmt.Println("Kind is Int? ", reflect.TypeOf(x).Kind() == reflect.Int) //Kind is Int?  true

	var f float64 = 3.4
	v := reflect.ValueOf(f)
	//v.SetFloat(7.3)//panic: reflect: reflect.Value.SetFloat using unaddressable value
	fmt.Println("v can set: ", v.CanSet()) //v can set:  false
	p := reflect.ValueOf(&f)
	fmt.Println("p can set: ", p.CanSet()) //p can set:  false
	v = p.Elem()
	fmt.Println("v can set: ", v.CanSet()) //v can set:  true
	v.SetFloat(7.1)
	fmt.Println(v.Interface()) //7.1
	fmt.Println(v)             //7.1

	type T struct {
		A int
		B string
	}
	t := T{1, "golang"}
	s := reflect.ValueOf(&t).Elem()
	typeOfT := s.Type()
	for i := 0; i < s.NumField(); i++ {
		field := s.Field(i)
		fmt.Printf("%d %s %s=%v\n", i, typeOfT.Field(i).Name, field.Type(), field.Interface())
	}
	//0 A int=1
	//1 B string=golang

	type S struct {
		F string `species:"gopher" color:"blue"`
	}

	st := S{}
	stt := reflect.TypeOf(st)
	field := stt.Field(0)
	fmt.Println(field.Tag.Get("color"), field.Tag.Get("species")) //blue gopher

	//var funcMap = make(map[string]func(string))
	//var funcMap = make(map[string]interface{})//invalid operation: cannot call non-function funcMap["say"] (map index expression of type interface{})
	//funcMap["say"] = say
	//funcMap["say"]("hello")

	var funcMap = make(map[string]interface{})
	funcMap["say"] = say
	Call(funcMap, "say", "hello") //hello
}
