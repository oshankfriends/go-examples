package main

import (
	"fmt"
	"reflect"
)

type Myint int

type Rectangle struct {
	Len, Br float64
}

func (r *Rectangle) Area() float64 {
	return r.Len * r.Br
}

/*func (r *Rectangle) String() string {
	return fmt.Sprintf("len: %f br: %f area: %f", r.len, r.br, r.Area())
}
*/

func main() {
	var a Myint = 8
	myIntVal := reflect.ValueOf(a)
	var r = &Rectangle{3, 4}
	val := reflect.ValueOf(r)
	fmt.Println(val)
	fmt.Println(val.String())
	fmt.Println("type is :",val.Type())
	fmt.Println("kind is :",val.Kind())
	fmt.Println("Kind is ptr :",val.Kind() == reflect.Ptr)
	fmt.Println("kind of Myint :",myIntVal.Kind())
	fmt.Println("type of Myint :",myIntVal.Type())
	fmt.Println("Settebility of myIntVal :",myIntVal.CanSet())
	fmt.Println("settebility of  val :",val.CanSet())
	val = val.Elem()
	fmt.Println("settebility of  val :",val.CanSet())
	val.Field(0).SetFloat(3.2)
	val.Field(1).SetFloat(4.2)
	fmt.Println(val)
	fmt.Println(r)
}
