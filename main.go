package main

import (
	"fmt"
	"math"
	"strings"
)

func sayGreeting(n string) {
	fmt.Printf("Good Morning! %v\n", n)
}
func sayBye(n string) {
	fmt.Printf("Good Bye! %v\n", n)
}

func cycleName(n []string, f func(string)) {
	for _, value := range n {
		f(value)
	}
}

func areaOfCircle(r float64) float64 {
	return math.Pi * r * r
}

func getInitials(n string) (string, string) {
	s := strings.ToUpper(n)
	names := strings.Split(s, " ")

	var initials []string
	for _, value := range names {
		initials = append(initials, value[:1])
	}

	if len(initials) > 1 {
		return initials[0], initials[1]
	}
	return initials[0], "_"

}

func updateName(n *string) {
	*n = "deepak"
}

func main() {
	// var ages [3]int = [3]int{1, 2, 3}
	// names := [3]string{"vishal", "deepak", "chhedi"}

	// fmt.Println(ages, len(ages))
	// fmt.Println(names, len(names))
	// x := 0
	// for x < 5 {
	// 	fmt.Println("index value:", x)
	// 	x++
	// }
	// for i := 0; i < 5; i++ {
	// 	fmt.Println("index of:", i)
	// }

	// names := []string{"vishal", "deepak", "chhedi"}

	// for i := 0; i < 3; i++ {
	// 	fmt.Println(names[i])
	// }
	// for i := 0; i < len(names); i++ {
	// 	fmt.Println(names[i])
	// }

	// sayGreeting("Deepak")
	// sayBye("Deepk")

	// cycleName(names, sayGreeting)

	// area1 := areaOfCircle(5)
	// area2 := areaOfCircle(10)

	// fmt.Println(area1, area2)
	// fmt.Printf("area1: %0.2f and area2: %0.3f", area1, area2)

	// fn1, sn1 := getInitials("deepak yadav")
	// fmt.Println(fn1, sn1)
	// fn2, sn2 := getInitials("aditya")
	// fmt.Println(fn2, sn2)

	// sayHello("Deepooo")
	// for _, value := range points {
	// 	println(value)
	// }

	// menu := map[string]float64{
	// 	"soup":    4.99,
	// 	"pie":     7.99,
	// 	"salad":   6.99,
	// 	"pudding": 3.55,
	// }

	// fmt.Println(menu)
	// fmt.Println(menu["pie"])

	// for key, value := range menu {
	// 	fmt.Printf("%v: %v\n", key, value)
	// }
	// name := "vishal"
	// m := &name
	// fmt.Println("before value of name is: ", name)
	// updateName(m)
	// fmt.Println("before value of name is: ", name)
	// fmt.Println("value of name is: ", name)
	// fmt.Println("memory address of name is: ", &name)
	// fmt.Println("m is: ", m)
	// fmt.Println("de-referencing m is: ", *m)

	myBill := newBill("Deepooo")
	myBill.updateTip(9)
	myBill.addItem("pastry", 24)

	fmt.Println(myBill.format())

}
