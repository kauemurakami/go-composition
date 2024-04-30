package main

import "fmt"

type person struct {
	name, lastname string
	age, height    uint8
}
type student struct {
	person
	course, university string
}

func main() {
	p := person{"Kauê", "Murakami", 28, 180}
	fmt.Println(p)
	s := student{p, "Medicine", "Harvard"}
	fmt.Println(s)
	fmt.Println(s.name)
}
