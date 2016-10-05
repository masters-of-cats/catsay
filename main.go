package main

import (
	"fmt"

	xsaypkg "github.com/fouralarmfire/xsay/mainframe"
)

func main() {
	cat := xsaypkg.NewMainframe(catAscii(), defaultMessage())

	fmt.Print("\n")
	cat.Say()
	fmt.Print("\n")
}
