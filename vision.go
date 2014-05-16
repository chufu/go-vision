package main

import (
	"fmt"
	"image"
	"log"
	"os"

	"github.com/skelterjohn/go.uik"
	"github.com/skelterjohn/go.uik/widgets"
	"github.com/skelterjohn/go.wde"
	_ "github.com/skelterjohn/go.wde/init"
)

func main() {
	go showPic(loadJpeg("a.jpg"))
	wde.Run()
}

func loadJpeg(filename string) image.Image {
	reader, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer reader.Close()

	ret, _, err := image.Decode(reader)
	if err != nil {
		log.Fatal(err)
	}
	return ret
}

func showPic(m image.Image) {
	wbounds := m.Bounds()

	w, err := uik.NewWindow(nil, wbounds.Max.X, wbounds.Max.Y)
	if err != nil {
		fmt.Println(err)
		return
	}
	w.W.SetTitle("pic")

	im := widgets.NewImage(widgets.ImageConfig{
		Image: m,
	})

	w.SetPane(&im.Block)

	w.Show()

	done := make(chan interface{}, 1)
	isDone := func(e interface{}) (accept, done bool) {
		_, accept = e.(uik.CloseEvent)
		done = accept
		return
	}
	w.Block.Subscribe <- uik.Subscription{isDone, done}

	<-done

	w.W.Close()

	wde.Stop()
}
