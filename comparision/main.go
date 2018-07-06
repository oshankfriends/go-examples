package main

import (
	"errors"
	"fmt"
)

var(
	ErrNoname = errors.New("No Name is present")
)

func GetError()error{
	return ErrNoname
}

func main() {
	if GetError() == ErrNoname {
		fmt.Println("identical")
		return
	}
	fmt.Println("non-identical")
}
