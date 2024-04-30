[![pt-br](https://img.shields.io/badge/language-pt--br-green.svg)](https://github.com/kauemurakami/go-inheritance/blob/main/README.pt-br.md)
[![en](https://img.shields.io/badge/language-en-orange.svg)](https://github.com/kauemurakami/go-inheritance/blob/main/README.md)

## "Inheritance"
Since GO is not an OOP language, we will talk about what we have closest to a conventional "Inheritance".

Let's start by creating two ```structs``` where one will "inherit" the other, our two new types will be ```person``` and ```student```:  
```go
package main

import "fmt"

type person struct {
	name, lastname string
	age, height    uint8
}
type student struct {

}

func main() {

	fmt.Println("x")
}
```  
Everyone agrees that our ```student``` also has the same fields as a ```person```, that is, the student is a person. So to *NOT* be redundant by doing something like:  
```go
...
type person struct {
	name, lastname string
	age, height    uint8
}
type student struct {
	name, lastname string
	age, height    uint8
  course, university string
}
func main ...
```  
To solve this, we can do the following, applying "inheritance" in GO:  
```go
...
type person struct {
	name, lastname string
	age, height    uint8
}
type student struct {
  person
  course, university string
}
func main ...
```  
You can pass the name of your other untyped ```struct``` into the ```struct``` that you want to "inherit" the fields of the ```struct``` declared in the way shown above.
So now our ```student``` struct has a ```person``` field and can *directly* use all its attributes.
*Why directly?*
Because we don't need to "access" ```person``` to use its attributes, example:
~~student.person.name = "Name"~~
We can access it *directly* like this:  
```go
student.name = "Name"
```  
That's why we omitted the second ```person``` in the declaration in ```student```
Another observation, GO does not allow keeping variables with primitive types in the code without using them, but it allows a struct that is not used, as the struct is a type.
*Using*  
```go
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
  //or
  s := student{person{"Kauê", "Murakami", 28, 180}, "Medicine", "Harvard"}
	fmt.Println(s)
	fmt.Println(s.name)
}
```  





