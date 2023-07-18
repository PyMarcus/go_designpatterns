package main

import (
	"fmt"
	"errors"
)

/*
Bridge is a design patterns that allow to separate
the implementation from abstration.
*/
type IDevice interface { // the bridge
	IsEnable() bool
	On()
	Off()
}

type Tv struct {
	enable bool
}

func (tv *Tv) On() {
	tv.enable = true
}

func (tv *Tv) Off() {
	tv.enable = false
}

func (tv Tv) IsEnable() bool {
	return tv.enable
}

type RemoteControl struct {
	device IDevice
}

func (r *RemoteControl) togglePower() error{
	if r.device == nil{
		return errors.New("Device is nil")
	}
	if r.device.IsEnable(){
		r.device.Off()
		fmt.Println("Device off")
	}else{
		r.device.On()
		fmt.Println("Device on")
	}
	return nil 
}

func main() {
	tv := &Tv{true}
	rc := RemoteControl{tv}
	if v := rc.togglePower(); v != nil{
		fmt.Println(v)
	}

}
