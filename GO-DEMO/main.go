package main

import (
	"fmt"

	"example.com/go-demo/mascot"
)

func main() {
	var m string = mascot.GetBestMascot()
	fmt.Println(m)
}
