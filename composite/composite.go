package main

import (
	"reflect"
	"fmt"
)

// Composite is a design pattern that compose objects into tree structures and then work with
// these structures as if they were individual objects
type Branch struct {
	leafs []Leaflet
	name  string
}

type Leaflet struct {
	name string
}

type IComposite interface {
	process()
}

func (b *Branch) add(leaf Leaflet) {
	b.leafs = append(b.leafs, leaf)
}

func (b *Branch) remove(positionOrName interface{}) {
	var index int 
	if reflect.TypeOf(positionOrName) == reflect.TypeOf("") {
		for i := 0; i < len(b.leafs); i++{
			if(b.leafs[i].name == positionOrName){
				index = i 
				break
			}
		}
		b.leafs = append(b.leafs[:index], b.leafs[index+1:]...)
	}else{
		index = positionOrName.(int)
		b.leafs = append(b.leafs[:index], b.leafs[index+1:]...)
	}
}

func (b Branch) process(){
	for _, v := range b.leafs{
		fmt.Printf("%s ", v.name)
	}
}
func main() {
	leaf := Leaflet{"1"}
	leaf2 := Leaflet{"2"}
	leaf3 := Leaflet{"3"}
	leaf4 := Leaflet{"4"}
	branch := Branch{name: "1"}
	branch.add(leaf)
	branch.add(leaf2)
	branch.remove(1)
	branch.add(leaf3)
	branch.add(leaf4)
	branch.process()
}
