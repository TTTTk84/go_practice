// 複数ユーザの処理をselectで待ち受ける

package main

import "fmt"

type client struct {
	name string
	id	 int
	room *room
}


type room struct {
	message chan string
	join chan *client
	done chan bool
}

func (c *client) chat() {
	c.room.join <- c
	for i := 0; i < 5; i++{
		c.room.message <- c.name + "さんのchat: hello"
	}
	c.room.done <- true
}

func (r *room) run() {
	count := 0
	for {
		select{
		case client := <- r.join: //参加client
			fmt.Printf("%sさん(id: %d)が参加しました\n", client.name, client.id)
		case mes := <- r.message:
			fmt.Printf("%s\n", mes)
		case <- r.done:
			count += 1
			if count > 1 {
				fmt.Printf("chat終了\n")
				return
			}
		}
	}
}

// selectを使うには、make()しなくてはいけない
func newRoom() *room {
	return &room{
		message: make(chan string),
		join: make(chan *client),
		done: make(chan bool),
	}
}

func main() {
	room := newRoom()
	c1 := &client{"taro",1,room}
	c2 := &client{"jiro",2,room}

	go c1.chat()
	go c2.chat()
	room.run()


}
