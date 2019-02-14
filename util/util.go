package util

import (
	"fmt"
	"unicode"
)

func Unhex(c byte) byte {
	switch {
	case '0' <= c && c <= '9':
		return c - '0'
	case 'a' <= c && c <= 'f':
		return c - 'a' + 10
	case 'A' <= c && c <= 'F':
		return c - 'A' + 10
	}
	return 0
}

func DyncmicType(t interface{}) {
	//var t interface{}
	//t = functionOfSomeType()
	//t = 3
	switch t := t.(type) {
	default:
		fmt.Printf("unexpected type %T\n", t)     // %T prints whatever type t has
	case bool:
		fmt.Printf("boolean %t\n", t)             // t has type bool
	case int:
		fmt.Printf("integer %d\n", t)             // t has type int
	case *bool:
		fmt.Printf("pointer to boolean %t\n", *t) // t has type *bool
	case *int:
		fmt.Printf("pointer to integer %d\n", *t) // t has type *int
	}
}

// A similar approach obviates the need to pass a pointer to a return value to simulate a reference parameter.
// Here's a simple-minded function to grab a number from a position in a byte slice,
// returning the number and the next position.
func NextInt(b []byte, pos int) (value, nextPos int) {
	for ; pos < len(b) && !unicode.IsDigit(rune(b[pos])); pos++ {
	}
	
	for ; pos < len(b) && unicode.IsDigit(rune(b[pos])); pos++ {
		value = value*10 + int(b[pos]) - '0'
	}
	nextPos = pos
	return value, nextPos
}
