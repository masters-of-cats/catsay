package main

import (
	"fmt"

	"github.com/fouralarmfire/xsay"
)

func main() {
	cat := xsay.New("catsay.ascii", "MEOW!")

	fmt.Print("\n")
	cat.Say()
	fmt.Print("\n")
}
