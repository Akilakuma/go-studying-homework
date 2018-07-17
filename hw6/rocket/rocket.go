package rocket

import (
	"fmt"
	"strconv"
)

type Rocket interface {
	standby()
	onReady()
	fire()
	upSky()
}

type AppleRocket struct {
	Name   string
	color  string
	Power  int
	weight int
}

func (a AppleRocket) standby() {
	fmt.Println(a.Name + " is standBy!")
}

func (a AppleRocket) onReady() {
	fmt.Println(a.Name + " is onReady!")
}

func (a AppleRocket) fire() {
	fmt.Println(a.Name + " fire!!!, It has power:" + strconv.Itoa(a.Power))
}

func (a AppleRocket) upSky() {
	fmt.Println(a.Name + " is upSky!")
}

type BananaRocket struct {
	Name     string
	category string
	Power    int
	weight   int
}

func (b BananaRocket) standby() {
	fmt.Println(b.Name + " is standBy!")
	fmt.Println("for banana king!")
}

func (b BananaRocket) onReady() {
	fmt.Println(b.Name + " is onReady!")
	fmt.Println("for banana king!")
}

func (b BananaRocket) fire() {
	fmt.Println(b.Name + " fire!!!, It has power: " + strconv.Itoa(b.Power))
	fmt.Println("for banana king!")
}

func (b BananaRocket) upSky() {
	fmt.Println(b.Name + " is upSky!")
	fmt.Println("for banana king!")
}

func Lanch(r Rocket) {
	r.standby()
	r.onReady()
	r.fire()
	r.upSky()
}
