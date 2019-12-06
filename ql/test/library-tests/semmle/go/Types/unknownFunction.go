package main

// This file tests type inference for expressions referencing undeclared entities.
// It is therefore expected to produce extractor warnings.

import "fmt"

func unknownFunctionTest() {
	e := unknownFunction()
	f := "hi " + e
	fmt.Println(e, f)
}
