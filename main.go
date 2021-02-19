package main

import (
	"fmt"

	"github.com/webview/webview"
	"github.com/yamnikov-oleg/wingo"
	"unsafe"
)

func onclose(w *wingo.Window) bool {
	fmt.Println("whhy")
	w.Hide()
	return false
}
func ontrayclick(w *wingo.Window) {
	w.Show()
}
func ontrayrightclick(w *wingo.Window) {
	traymenu.StartContext(w)

}
func oncreate(w *wingo.Window) {
	fmt.Println("111111")
	hand := w.GetHandle()
	wv := webview.NewWindow(true, unsafe.Pointer(&hand))
	wv.Navigate("http://www.baidu.com")
	wv.Run()
	defer wv.Destroy()
}

var traymenu *wingo.Menu

func main() {
	w := wingo.NewWindow(true, true)
	size := wingo.Vector{500, 500}
	w.SetSize(size)
	w.OnClose = onclose
	w.OnTrayClick = ontrayclick
	w.OnTrayRightClick = ontrayrightclick
	//w.OnCreate=oncreate
	icon := wingo.LoadIcon(101)
	w.SetIcon(icon)
	w.AddTrayIcon(icon, "hellofromaiel")
	traymenu = wingo.NewContextMenu()
	exitbtn := traymenu.AppendItemText("退出")
	exitbtn.OnClick = func(item *wingo.MenuItem) {
		w.Destroy()
	}

	oncreate(w)
	w.Show()

	wingo.Start()
}
