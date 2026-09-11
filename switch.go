package main

import (
	"fmt"
)

func main() {

	i := 3
	fmt.Print("Write ", i, "as ")
	switch i {
	case 1:
		println(" one")
	case 2:
		println(" two")
	case 3:
		println(" four")
		fallthrough
	case 4:

		fmt.Println(" You won bronze ")

	}
}
