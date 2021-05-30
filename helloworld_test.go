package main

import "testing"

func TestHelloWorld(t *testing.T){
	result := HelloWorld("Sandi Pradana")
	if result != "Hello World, Sandi Pradana" {
		panic("Result is not \"Hello World, Sandi Pradana\"")
	}
}