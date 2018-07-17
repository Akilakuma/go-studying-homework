package main

import (
	"fmt"
	"rocket"
)

func main() {
	aRock := rocket.AppleRocket{
		Name:  "apple",
		Power: 56666,
	}

	bRock := rocket.BananaRocket{
		Name:  "banana",
		Power: 777,
	}

	rocket.Lanch(aRock)
	fmt.Println()
	rocket.Lanch(bRock)
}
