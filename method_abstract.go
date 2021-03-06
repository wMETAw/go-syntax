package main

import "fmt"

type Human struct {
	name string
	age int
	phone string
}

type Student struct {
	Human
	school string
}

type Employee struct {
	Human
	company string
}

// human上でメソッドを定義
func (h *Human) SayHi(){
	fmt.Printf("Hi, I am %s you can call me on %s\n", h.name, h.phone)
}

// EmployeeのmethodでHumanのmethodを書き直す
func (e *Employee) SayHi(){
	fmt.Printf("Hi, I am %s, I work at %s. Call me on %s\n", e.name,
		e.company, e.phone) //Yes you can split into 2 lines here.
}


func main() {
	mark := Student{Human{"Mark",25,"080-1111-2222"}, "MIT"}
	sam := Employee{Human{"Sam",45,"111-888-XXXX"}, "Go Inc"}
	mark.SayHi()
	sam.SayHi()
}
