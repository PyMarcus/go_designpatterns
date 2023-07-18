package main

import "fmt"

/*
Adapter pattern is a structural GOF design pattern that translate
incompatible types.
Adapter translate a incompatible interface to a compatible
interface to the client.
*/
type Adapter struct {
	adaptee Adaptee
}

type Adaptee struct {
	adapterType int
}

type IProcess interface {
	process()
}

func (a Adapter) process(){
	fmt.Println("Converting...")
	a.adaptee.convert()
}

func (a Adaptee) convert(){
	fmt.Println("adapting...")
}

func convert(class IProcess){
	class.process()
}

func main() {
	process := Adapter{}
	convert(process)
}
