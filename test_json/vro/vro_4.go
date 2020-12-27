package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

type Person struct {
  Id            int			`json:"id"`
  Name          string	`json:"name"`
	Birthday      *string	`json:"birthday"`
	Vivid			struct {
		Color	string	`json:color`
		Weapon string `json:weapon`
	} `json:"vivid_info"`
}

func (p *Person) GetInfo() string {
	if p.Birthday != nil {
		return fmt.Sprintf("%d : %s [%s]", p.Id, p.Name, *p.Birthday)
	}
	return fmt.Sprintf("%d : %s", p.Id, p.Name)
}

func main() {
  // json読み込み
  bytes, err := ioutil.ReadFile("vro.json")
  if err != nil {
    log.Fatal(err)
  }

  // jsonデコード
  var persons []Person
  if err := json.Unmarshal(bytes, &persons); err != nil {
    log.Fatal(err)
  }

  for _, p := range persons {
		fmt.Printf(p.GetInfo())
  }

}
