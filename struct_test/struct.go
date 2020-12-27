//構造体型のポインタ型まわり

package main

import "fmt"

type person struct {
	name string
	id 	 int
}

func main() {
	p1 := person{"taro", 1}
	fmt.Printf("%v\n",p1) // {taro 1}

	p2 := new(person)
	*p2 = person{"jiro", 2}
	fmt.Printf("%v\n",p2) // &{jiro 2}

	// p2を短く書く
	p3 := &person{"saburo", 3}
	fmt.Printf("%v\n", p3)// &{saburo 3}
}
