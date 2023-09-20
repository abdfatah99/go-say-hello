package go_module_say_hello

import (
	"fmt"
)

func SayHello(name string) string {
	fmt.Println("Say Hello From Go-Module")
	return fmt.Sprintf("Hello, I'm %s; from SayHello() func", name)
}

// add new feature to module, Introduction function representing information
// about my personal
type Person struct {
	Name, Address string
	Age           int8
}

func Introdution() string {
	fatah := Person{"Abdul Fatah", "Bogor", 23}

	return fmt.Sprintf("My name is %s, from %s, and I'm %d years old",
		fatah.Name, fatah.Address, fatah.Age)
}
