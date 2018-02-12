package main

import (
	"fmt"
)

type Human struct {
	name string
	age int
	phone string
}

type Student struct {
	Human //匿名フィールドHuman
	school string
	loan float32
}

type Employee struct {
	Human //匿名フィールドHuman
	company string
	money float32
}

//HumanオブジェクトにSayHiメソッドを実装します。
func (h Human) SayHi() {
	fmt.Printf("Hi, I am %s you can call me on %s\n", h.name, h.phone)
}

// Humanオブジェクトにsingメソッドを実装
func (h Human) Sing(lyrics string) {
	fmt.Println("La la, la la la, la la la la la...", lyrics)
}

// HumanメソッドにGuzzleメソッドを実装
func (h Human) Guzzle(beerStein string) {
	fmt.Println("Guzzle Guzzle Guzzle...", beerStein)
}

// EmployeeはHumanのSayHiメソッドをオーバーロードします。
func (e Employee) SayHi(){
	fmt.Printf("Hi, I am %s, I work at %s. Call me on %s\n", e.name, e.company, e.phone) //この行は複数に渡ってもかまいません。
}

// StudentはBorrowMoneyメソッドを実装します。
func (s Student) BorrowMoney(amount float32) {
	s.loan += amount
}

// EmployeeはSpendSalaryメソッドを実装します。
func (e *Employee) SpendSalary(amout float32) {
	e.money -= amout
}

type Men interface {
	SayHi()
	Sing(lyrics string)
	Guzzle(beerStein string)
}

type YoungChap interface {
	SayHi()
	Sing(song string)
	BorrowMoney(amount float32)
}

type ElderlyGent interface {
	SayHi()
	Sing(song string)
	SpendSalary(amount float32)
}

func main() {
	mike := Student{Human{"Mike", 25, "222-222-XXX"}, "MIT", 0.00}
	paul := Student{Human{"Paul", 26, "111-222-XXX"}, "Harvard", 100}
	sam := Employee{Human{"Sam", 36, "444-222-XXX"}, "Golang Inc.", 1000}
	tom := Employee{Human{"Tom", 37, "222-444-XXX"}, "Things Ltd.", 5000}

	// Men型の変数i
	var i Men

	//iにはStudentを保存できます。
	i = mike
	fmt.Println("This is Mike, a Student:")
	i.SayHi()
	i.Sing("November rain")

	// iにはEmployeeを保存することもできます。
	i = tom
	fmt.Println("This is Tom, an Employee:")
	i.SayHi()
	i.Sing("Born to be wild")

	//sliceのMenを定義します。
	fmt.Println("Let's use a slice of Men and see what happens")
	x := make([]Men,3)

	//この３つはどれも異なる要素ですが、同じインターフェースを実装しています。
	x[0],x[1],x[2] = paul, sam, mike

	for _,v := range x {
		v.SayHi()
	}
}
