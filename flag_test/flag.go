// コマンドライン引数を使うためのもの
package main

import (
	"flag"
	"fmt"
)

func main() {
	  // flag.String(-nameの形で入力させる, デフォルトの値, コメント)
    var name = flag.String("name", "taro", "デフォルトだとtaro")
    flag.Parse()
		fmt.Println("あなたの名前は:", *name)
		/*
		./flag
		出力: あなたの名前は: taro
		./flag -name jiro
		出力: あなたの名前は: jiro
		*/
}
