package gosayhello

import "fmt"

func SayHello(name string) {
	fmt.Println("Hello " + name)
}

func SayHelloWorld() {
	SayHello("Audyari Wiyono")
}

func PanggilLagi() {
	SayHelloWorld()
}
