// 関数のみのchannel間の通信

package main

import "fmt"

func chan_1(){
	ch := make(chan string)

	go chan_2(ch)

	finish := <- ch
	fmt.Printf("%s\n",finish)
}

func chan_2(c chan string){
	c <- "chan_2"
}

func main() {
	chan_1()
}
