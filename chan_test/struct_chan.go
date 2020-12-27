// 構造体を絡めたchannel間の通信

package main

import "fmt"

type user1 struct {
	name string
	id	 int
}

type user2 struct {
	name string
	id   int
}

func NewUser2(name string, id int) *user2 {
	return &user2{name,id}
}

func (u *user1) chat() {
	chat := make(chan string)
	u2 := NewUser2("jiro", 2)

	go u2.chat(chat)

	res := <- chat
	fmt.Printf("jiroからの送信: %s\n", res)
}

func (u *user2) chat(c chan string) {
	c <- "taro 元気ですか？"
}

func main() {
	u1 := &user1{"taro",1}
	u1.chat()
}
