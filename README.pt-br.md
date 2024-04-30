[![pt-br](https://img.shields.io/badge/language-pt--br-green.svg)](https://github.com/kauemurakami/go-inheritance/blob/main/README.pt-br.md)
[![en](https://img.shields.io/badge/language-en-orange.svg)](https://github.com/kauemurakami/go-inheritance/blob/main/README.md)

## "Herança"
Como GO não é uma linguagem POO, falaremos sobre o que temos mais próximo de uma "Herança" convencional.    

Vamos começar criando dois ```structs``` onde um "herdará" o outro, nossos dois novos tipos serão ```person``` e ```student```:  
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
Todos concordam que o nosso ```student``` também possui os mesmos campos de uma ```person```, ou seja o estudando é uma pessoa. Então para *NÃO* ser redundante fazendo algo como:
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
Para solucionar isso, podemos fazer o seguinte, aplicando a "herança" em GO:  
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
Você pode passar o nome de seu outro ```struct``` sem tipo no ```struct``` que deseja que "herde" os campos do ```struct``` declarado da forma mostrada acima.  
Então agora nosso struct ```student``` tem um campo ```person``` e pode usar *diretamente* todos seus atributos.  
*Por que diretamente?*
Pois não necessitamos "acessar" ```person``` para usar seus atributos, exemplo:  
~~student.person.nome = "Name"~~  
Podemos acessar *diretamente* dessa forma:  
```go
student.name = "Name"
```  
Por isso omitimos o segundo ```person``` na declaração em ```student```  
Outra observação, o GO não permite manter variáveis com os tipos primitivos no código sem usar, mas permite um struct que não seja usado, pois o struct é um tipo.
*Usando*
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




