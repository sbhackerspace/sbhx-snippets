// Steve Phillips / elimisteve
// 2012.06.27

package main

import (
	"fmt"
	"reflect"
)

func main() {
	type Person struct { First, Last string; Age int }
	t := Person{First: "Steve", Last: "P", Age: 26}

	s := reflect.ValueOf(&t).Elem()
    typeOfT := s.Type()
    for i := 0; i < s.NumField(); i++ {
        f := s.Field(i)
        fmt.Printf("%s: %v\n", typeOfT.Field(i).Name, f.Interface())

		// Note: This prints more info, like the type information:
        // fmt.Printf("%d: %s %s = %v\n", i,
        //     typeOfT.Field(i).Name, f.Type(), f.Interface())
    }
}
