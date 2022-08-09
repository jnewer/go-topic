package main

import (
	"fmt"

	"github.com/spf13/cast"
)

func main() {
	fmt.Printf("cast.ToString(\"kang\"): %T\n", cast.ToString("kang"))
	fmt.Printf("cast.ToString(8): %T\n", cast.ToString(8))
	fmt.Printf("cast.ToString(8.31): %v\n", cast.ToString(8.31))
	fmt.Printf("cast.ToString([]byte(\"one time\")): %v\n", cast.ToString([]byte("one time")))
	fmt.Printf("cast.ToString(nil): %v\n", cast.ToString(nil)) //""

	var foo interface{} = "one more time"
	fmt.Printf("cast.ToString(foo): %v\n", cast.ToString(foo))

	fmt.Printf("cast.ToInt(8): %v\n", cast.ToInt(8))
	fmt.Printf("cast.ToInt(8.31): %v\n", cast.ToInt(8.31))   //8
	cast.ToInt("8")                                          //8
	fmt.Printf("cast.ToInt(true): %v\n", cast.ToInt(true))   //1
	fmt.Printf("cast.ToInt(false): %v\n", cast.ToInt(false)) //0

	var eight interface{} = 8
	fmt.Printf("cast.ToInt(eight): %v\n", cast.ToInt(eight)) //8
	fmt.Printf("cast.ToInt(nil): %v\n", cast.ToInt(nil))     //0
}
