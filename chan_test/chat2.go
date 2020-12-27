// 複数ユーザの処理をselectで待ち受ける
// mapでユーザの入退室を管理、

package main

import (
	"fmt"
	"strconv"
)

type client struct {
	name string
	id	 int
	room *room			 // 入ってるルーム
}


type room struct {
	name string								// ルーム名
	forward chan string				// 送られてきたメッセージ
	join chan *client					// 参加するクライアント
	leave chan *client				// 退出するクライアント
	clients map[*client]bool  // いま入室してるクライアント
}

func NewClient(name string, id int, room *room) *client{
	return &client{
		name: name,
		id:		id,
		room: room,
	}
}

func NewRoom(name string) *room {
	return &room{
		name: name,
		forward: make(chan string),
		join: make(chan *client),
		leave: make(chan *client),
		clients: make(map[*client]bool),
	}
}

func (c *client) write() {
	c.room.join <- c
	for i := 0; i < 5; i++ {
		mes := (c.name + "さんの発言: " + strconv.Itoa(i))
		c.room.forward <- mes
	}
	c.room.leave <- c

}

func (r *room) run() {
	// count := 0
	for {
		select{
		case client := <- r.join:  //入室
			fmt.Printf("%sさん(id: %d)が参加しました\n", client.name, client.id)
			r.clients[client] = true
		case client := <- r.leave: //退室
			fmt.Printf("%sさん(id: %d)が退室しました\n", client.name, client.id)
			delete(r.clients, client)
		case msg := <- r.forward:  //メッセージをもらう
			fmt.Printf("%s\n", msg)
		}
	}
}

func main() {
	room := NewRoom("chat room1")
	c1 := NewClient("taro",1,room)
	c2 := NewClient("jiro",2,room)

	go c1.write()
	go c2.write()
	room.run()
}
