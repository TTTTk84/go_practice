package main

import (
	"os"
	"text/template"
)

// 使いたい要素を入れる
type Person struct {
	Name string
	Id   int
}

func main() {
	// 値を入れたインスタンスを入れる
	person := Person{"tk84", 20}
	// *template.Template.Parse() で テンプレートを作る
	// *template.Template は template.Newでとれる
	tmpl, err := template.New("test").Parse("name: {{.Name}}, id: {{.Id}}\n")
	if err != nil { panic(err) }
	// *template.Template.Executeで出力 (出力先,構造体)
	err = tmpl.Execute(os.Stdout, person)
	if err != nil { panic(err) }

}
