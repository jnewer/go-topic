package main

import (
	"errors"
	"fmt"
)

func Div(a int, b int) (int, error) {
	if b == 0 {
		return -1, errors.New("除数不能为零")
	}
	return a / b, nil
}

type WrapError struct {
	msg string
	err error
}

func (e *WrapError) Error() string {
	return e.msg
}

func (e *WrapError) Unwrap() error {
	return e.err
}

type ErrorString struct {
	s string
}

func (e *ErrorString) Error() string {
	return e.s
}

func main() {
	_, err := Div(1, 0)
	fmt.Printf("err.Error(): %v\n", err.Error())
	fmt.Printf("err: %v\n", err)

	err1 := errors.New("new error")
	err2 := fmt.Errorf("err2: [%w]", err1)
	err3 := fmt.Errorf("err3: [%w]", err2)
	fmt.Println(err3) //err3: [err2: [new error]]

	fmt.Println(errors.Unwrap(err3))                //err2: [new error]
	fmt.Println(errors.Unwrap(errors.Unwrap(err3))) //err2: [new error]

	fmt.Println(errors.Is(err3, err2)) //true
	fmt.Println(errors.Is(err3, err1)) //true

	var targetError *ErrorString
	err4 := fmt.Errorf("new error:[%w]", &ErrorString{s: "target error"})
	fmt.Println(errors.As(err4, &targetError)) //true
}
