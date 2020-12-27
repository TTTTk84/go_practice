/*
{
  "Title": "Title",
  "Url": "Url",
  "LikedAt": "LikedAt"
}


{
  "content": "**タイトル:** \n**URL :** \n**日時  :** "
}
直す
*/
package main

import (
	"fmt"
)

type Person struct {
  Id            int
  Name          string
  Birthday      string
}

func main() {
  // json読み込み
  // bytes, err := ioutil.ReadFile("vro.json")
  // if err != nil {
  //   log.Fatal(err)
  // }

  // // jsonデコード
  // var persons []Person
  // if err := json.Unmarshal(bytes, &persons); err != nil {
  //   log.Fatal(err)
  // }

  // for _, p := range persons {
  //   fmt.Printf("%d : %s\n", p.Id, p.Name)
  // }

  var p = new(Person)
  *p = Person{1,"hoge","20200404"}
  fmt.Printf("%v\n", p)


}
