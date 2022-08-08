package main

import (
	"fmt"
	"github.com/spf13/viper"
)

func ReadIni() {

	v := viper.New()
	v.AddConfigPath("./conf")
	v.SetConfigName("config")
	v.SetConfigType("ini")

	err := v.ReadInConfig()

	if err != nil {
		panic(err)
	}
	s := v.GetString("db.username")
	fmt.Printf("s: %v\n", s)
}
func ReadYaml() {
	v := viper.New()
	v.AddConfigPath("./conf")
	v.SetConfigName("config2")
	v.SetConfigType("yml")

	err := v.ReadInConfig()

	if err != nil {
		panic(err)
	}

	//fmt.Println(v)
	s := v.GetString("db.username")
	fmt.Printf("s: %v\n", s)
}
func main() {
	ReadIni()
	ReadYaml()
}
