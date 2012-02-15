// Steve Phillips / elimisteve
// 2012.02.07

package main

import (
	"fmt"
	"reflect"
)

func main() {
	//
	// Determine the type of a variable
	//
	var i interface{} = 5 // i is an interface{} value
	t := reflect.TypeOf(i)
	fmt.Printf("i's underlying value is %v\n\n", t) // t stores i's type

	//
	// .Type() v .Kind()
	//
	type MyInt int
	myVal := reflect.ValueOf(MyInt(10))
	fmt.Printf("myVal is a       %T\n", myVal)
	fmt.Printf("myVal.Type() is  %v\n", myVal.Type()) // Static, declared type
	fmt.Printf("myVal.Kind() is  %v\n", myVal.Kind()) // Underlying type/"kind"

	// Learn more at
	// http://blog.golang.org/2011/09/laws-of-reflection.html and
	// http://stackoverflow.com/questions/6395076/in-golang-using-reflect-how-do-you-set-the-value-of-a-struct-field
}