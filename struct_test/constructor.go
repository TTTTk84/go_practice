// goではコンストラクタがないため、Newを使って初期値の代入を行う
// 基本、知らないパッケージで.New() New~() みたいなのでてきたら、
// 構造体のポインタ返してると思ってよいと思う

package main

import "fmt"


type person struct {
	name string
	id	 int
}

func NewPerson(name string, id int) *person {
	return &person{name,id}
}


func main() {
	p := NewPerson("tk84",1)
	fmt.Printf("%v\n", p) // &{tk84 1}

}
