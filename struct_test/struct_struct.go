// 構造体のメンバを、構造体型で宣言する

package main

import "fmt"

type skill struct {
	name string
	level int
}

type person struct {
	name 	string
	id 	 	int
	skill *skill
}

func (s *skill) output_skill() {
	fmt.Printf("スキル名: %s, スキルレベル: %d\n", s.name,s.level)
}

func main() {
	// 多分一番短い
	p := &person{"tk84", 1, &skill{"草むしり", 3}}
	p.skill.output_skill()

	// newを使うやり方
	p2 := new(person)
	p2_skill := new(skill)
	*p2_skill = skill{"ランニング", 3}
	*p2 = person{"taro", 2, p2_skill}
	p2.skill.output_skill()


}
