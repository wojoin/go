package main

import (
	"log"
	"fmt"
	"github.com/wojoin/go/src/util"
)



func main() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	for pos, char := range "日本\x80語" { // \x80 is an illegal UTF-8 encoding
		fmt.Printf("character %#U starts at byte position %d\n", char, pos)
	}

	for c := 'a'; c <= 'f'; c = c + 1 {
		fmt.Print(util.Unhex(byte(c)))
	}
	fmt.Println("========")
	for c := 'A'; c <= 'F'; c = c + 1 {
		fmt.Print(util.Unhex(byte(c)))
	}

	fmt.Println("========")

	util.DyncmicType(3)

	log.Println("hello")

	num, pos := util.NextInt([]byte("a123a456"), 0)
	fmt.Println(num)
	fmt.Println(pos)
	
	// defer statement is likely to stack, put statement into stack, it execute before the function executing the defer returns
	// Deferred functions are executed in LIFO order
	for i := 0; i < 5; i++ {
		defer fmt.Printf("%d ", i) // output: 4 3 2 1 0
	}

}


