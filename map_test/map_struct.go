// 構造体の要素に、mapを入れる

package main

import "fmt"

type person struct {
	id map[int]bool
	name string
}

type group struct {
	gid			int
	persons map[*person]bool
}

func (g *group) persons_shedules() {
	for p, p_bool := range g.persons {
		if p_bool {
			fmt.Printf("今予定が空いているのは %v さん\n",p.name)
		}
	}
}

func main() {
	// 要素にmapを使った例
	p := &person{id: map[int]bool{1: true}, name: "tk84"}
	p2 := &person{id: map[int]bool{2: true}, name: "taro"}
	p3 := &person{id: map[int]bool{3: true}, name: "jiro"}
	fmt.Printf("%v\n%v\n%v\n", p, p2, p3)

	// 要素にmapを使い、keyをほかの構造体にした例
	g := &group{gid: 1, persons: map[*person]bool{
		p:  true,
		p2: false,
	}}
	g.persons[p3] = true

	g.persons_shedules()
}
