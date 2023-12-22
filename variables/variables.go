// In Go, variables are declared explicitly and used by the compiler
// use can use the "var" keyword to declare one or more variables.
// you can also declare more than one variable at a time.
// Go will infer the type of variable that is declared.
// variable can also be declared using the shorthand format(:=)

package main

import "fmt"

func main() {
	var topic string = "This is how variables are declared in Golang"
	fmt.Println(topic)
	var age int = 45
	fmt.Println(age)
	var isPresent bool = true
	fmt.Println(isPresent)

	// declared using the shorthand
	f := 56
	fmt.Println(f)
}