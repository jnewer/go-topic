package main

import (
	"bytes"
	"fmt"
	"regexp"
)

func main() {
	fmt.Println(regexp.Match("H.* ", []byte("Hello World!"))) //true

	r := bytes.NewReader([]byte("Hello World!"))
	fmt.Println(regexp.MatchReader("H.* ", r)) //true

	fmt.Println(regexp.MatchString("H.* ", "Hello World!")) //true

	fmt.Println(regexp.QuoteMeta("(?P:Hello) [a-z]")) //\(\?P:Hello\) \[a-z\]

	reg, err := regexp.Compile(`\w+`)
	fmt.Printf("%q,%v\n", reg.FindString("Hello World!"), err) //"Hello"

	reg2, err2 := regexp.CompilePOSIX(`[[:word:]]+`)
	fmt.Printf("%q,%v\n", reg2.FindString("Hello World!"), err2) //"Hello"

	reg3 := regexp.MustCompile(`\w+`)
	fmt.Println(reg3.FindString("Hello World!")) //Hello

	reg4 := regexp.MustCompilePOSIX(`[[:word:]]+`)
	fmt.Printf("%q\n", reg4.FindString("Hello World!")) //"Hello"

	fmt.Printf("%q\n", reg3.Find([]byte("Hello World"))) //"Hello"

	fmt.Println(reg3.FindString("Hello World!")) //Hello

	fmt.Printf("%q\n", reg3.FindAll([]byte("Hello World!"), -1)) //["Hello" "World"]

	fmt.Printf("%q\n", reg3.FindAllString("Hello World!", -1)) //["Hello" "World"]

	fmt.Println(reg3.FindIndex([]byte("Hello World!"))) //[0 5]

	fmt.Println(reg3.FindStringIndex("Hello World!")) //[0 5]

	r2 := bytes.NewReader([]byte("Hello World!"))
	fmt.Println(reg3.FindReaderIndex(r2)) //[0 5]

	fmt.Println(reg3.FindAllIndex([]byte("Hello World!"), -1)) //[[0 5] [6 11]]

	fmt.Println(reg3.FindAllStringIndex("Hello World!", -1)) //[[0 5] [6 11]]

	reg5 := regexp.MustCompile(`(\w)(\w)+`)
	fmt.Printf("%q\n", reg5.FindSubmatch([]byte("Hello World!"))) //["Hello" "H" "o"]

	fmt.Printf("%q\n", reg5.FindStringSubmatch("Hello World!")) //["Hello" "H" "o"]

	fmt.Printf("%q\n", reg5.FindAllSubmatch([]byte("Hello World!"), -1)) //[["Hello" "H" "o"] ["World" "W" "d"]]

	fmt.Printf("%q\n", reg5.FindAllStringSubmatch("Hello World!", -1)) //[["Hello" "H" "o"] ["World" "W" "d"]]

	fmt.Println(reg5.FindSubmatchIndex([]byte("Hello World!"))) //[0 5 0 1 4 5]
	fmt.Println(reg5.FindStringSubmatchIndex("Hello World!"))   //[0 5 0 1 4 5]

	r3 := bytes.NewReader([]byte("Hello World!"))
	fmt.Println(reg5.FindReaderSubmatchIndex(r3)) //[0 5 0 1 4 5]

	fmt.Println(reg5.FindAllStringSubmatchIndex("Hello World!", -1)) //[[0 5 0 1 4 5] [6 11 6 7 10 11]]

	reg6 := regexp.MustCompile(`(\w+),(\w+)`)
	src := []byte("Golang,World!")
	dst := []byte("Say: ")
	template := []byte("Hello $1, Hello $2")
	match := reg6.FindSubmatchIndex(src)
	fmt.Printf("%q\n", reg6.Expand(dst, template, src, match)) //"Say: Hello Golang, Hello World"

	src2 := "Golang,World!"
	dst2 := []byte("Say: ")
	template2 := "Hello $1, Hello $2"
	match2 := reg6.FindStringSubmatchIndex(src2)
	fmt.Printf("%q\n", reg6.ExpandString(dst2, template2, src2, match2)) //"Say: Hello Golang, Hello World"

	reg7 := regexp.MustCompile(`Hello[\w\s]+`)
	fmt.Println(reg7.LiteralPrefix()) //Hello false

	reg8 := regexp.MustCompile(`Hello`)
	fmt.Println(reg8.LiteralPrefix()) //Hello true

	text := `Hello World, 123 Go!`
	pattern := `(?U)H[\w\s]+o` // 正则标记“非贪婪模式”(?U)
	reg9 := regexp.MustCompile(pattern)
	fmt.Printf("%q\n", reg9.FindString(text)) //"Hello"

	b := []byte(`Hello World`)
	reg10 := regexp.MustCompile(`Hello\w+`)
	fmt.Println(reg10.Match(b)) //false

	reg11 := regexp.MustCompile(`Hello[\w\s]+`)
	fmt.Println(reg11.Match(b)) //true

	r12 := bytes.NewReader([]byte("Hello World!"))
	reg12 := regexp.MustCompile(`Hello\w+`)
	fmt.Println(reg12.MatchReader(r12)) //false

	r12.Seek(0, 0)
	reg13 := regexp.MustCompile(`Hello[\w\s]+`)
	fmt.Println(reg13.MatchReader(r12)) //true

	s14 := `Hello World!`
	reg14 := regexp.MustCompile(`Hello\w+`)
	fmt.Println(reg14.MatchString(s14)) //false

	reg15 := regexp.MustCompile(`Hello[\w\s]+`)
	fmt.Println(reg15.MatchString(s14)) //true

	reg16 := regexp.MustCompile(`(?U)(?:Hello)(\s+)(\w+)`)
	fmt.Println(reg16.NumSubexp()) //2

	b17 := []byte("Hello World, 123 Go!")
	reg17 := regexp.MustCompile(`(Hell|G)o`)
	rep17 := []byte("${1}ooo")
	fmt.Printf("%q\n", reg17.ReplaceAll(b17, rep17)) //"Hellooo World, 123 Gooo!"

	b18 := "Hello World, 123 Go!"
	reg18 := regexp.MustCompile(`(Hell|G)o`)
	rep18 := "${1}ooo"
	fmt.Printf("%q\n", reg18.ReplaceAllString(b18, rep18)) //"Hellooo World, 123 Gooo!"

	b19 := []byte("Hello World, 123 Go!")
	reg19 := regexp.MustCompile(`(Hell|G)o`)
	rep19 := []byte("${1}ooo")
	fmt.Printf("%q\n", reg19.ReplaceAllLiteral(b19, rep19)) //"${1}ooo World, 123 ${1}ooo!"

	b20 := "Hello World, 123 Go!"
	reg20 := regexp.MustCompile(`(Hell|G)o`)
	rep20 := "${1}ooo"
	fmt.Printf("%q\n", reg20.ReplaceAllLiteralString(b20, rep20)) //"${1}ooo World, 123 ${1}ooo!"

	s21 := []byte("Hello World!")
	reg21 := regexp.MustCompile("(H)ello")
	rep21 := []byte("$0$1")
	fmt.Printf("%s\n", reg21.ReplaceAll(s21, rep21)) //HelloH World!

	fmt.Printf("%q\n", reg21.ReplaceAllFunc(s21,
		func(b []byte) []byte {
			rst := []byte{}
			rst = append(rst, b...)
			rst = append(rst, "$1"...)
			return rst
		}))
	//"Hello$1 World!"

	s22 := "Hello World!"
	reg22 := regexp.MustCompile("(H)ello")
	rep22 := "$0$1"
	fmt.Printf("%s\n", reg22.ReplaceAllString(s22, rep22)) //HelloH World!

	fmt.Printf("%q\n", reg22.ReplaceAllStringFunc(s22,
		func(b string) string {
			return b + "$1"
		},
	))
	//"Hello$1 World!"

	s23 := "Hello World\tHello\nGolang"
	reg23 := regexp.MustCompile(`\s`)
	fmt.Printf("%q\n", reg23.Split(s23, -1)) //["Hello" "World" "Hello" "Golang"]

	reg24 := regexp.MustCompile("Hello.*$")
	fmt.Printf("%q\n", reg24.String()) //"Hello.*$"

	reg25 := regexp.MustCompile(`(?P<Hello>[a-zA-Z]+) (?P<World>[a-zA-Z]+)`)
	fmt.Printf("%q\n", reg25.SubexpNames()) //["" "Hello" "World"]

}
