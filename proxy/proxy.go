package main

import "fmt"

// Proxy is a design pattern that provides a substitute or placeholder
// for another object.
type IYoutubeLib interface {
	ListVideos()
	GetVideoInfo()
}

type CachedYoutube struct {
}

func (c CachedYoutube) ListVideos() {
	fmt.Println("Video A")
	fmt.Println("Video B")
}

func (c CachedYoutube) GetVideoInfo() {
	fmt.Println("Cache request")
}

type UserInteraction struct{
	proxy IYoutubeLib
}

func (u UserInteraction) access(option int){
	if option == 1{
		u.proxy.GetVideoInfo()
	}else{
		u.proxy.ListVideos()
	}
}
func main() {
	ui := UserInteraction{
		proxy: CachedYoutube{},
	}

	ui.access(1)
}
