package main

import (
	"fmt"
	"path/filepath"
)



func main() {
	f := filepath.Join("test", "a.txt")
	fmt.Printf("%s\n", f) // 出力: test/a.txt
}
