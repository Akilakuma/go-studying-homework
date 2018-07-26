package main

import "fmt"

type Iphone interface {
	turnON()
	fingerPrint()
	faceCheck()
}

type Iphone5 struct {
	name string
}

func (i Iphone5) turnON() {
	fmt.Println("I'm " + i.name + " power on, but it feel little slow!")
}

func (Iphone5) fingerPrint() {
	fmt.Println("Iphone5 no fingerPrint!")
}

func (Iphone5) faceCheck() {
	fmt.Println("Iphone5 no faceCheck!")
}

type Iphone8 struct {
	name string
}

func (i Iphone8) turnON() {
	fmt.Println("I'm " + i.name + " power on, is very fast!")
}

func (Iphone8) fingerPrint() {
	fmt.Println("put your fingerPrint!")
}

func (Iphone8) faceCheck() {
	fmt.Println("Iphone8 no faceCheck!")
}

type IphoneX struct {
	name string
}

func (i IphoneX) turnON() {
	fmt.Println("I'm " + i.name + " power on, is very fast!")
}

func (IphoneX) fingerPrint() {
	fmt.Println("put your fingerPrint!")
}

func (IphoneX) faceCheck() {
	fmt.Println("Give me your faceCheck!")
}

func main() {
	// phoneList := []Iphone{
	// 	new(Iphone5),
	// 	new(Iphone8),
	// 	new(IphoneX),
	// }

	a := Iphone5{"iphone5"}
	b := Iphone8{"iphone8"}
	c := IphoneX{"iphoneX"}
	phoneList := []Iphone{a, b, c}

	for _, myphone := range phoneList {
		myphone.turnON()
		myphone.fingerPrint()
		myphone.faceCheck()
	}
}
