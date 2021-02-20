package main

import (
	"awesomeProject1/data"
	"fmt"
	assetfs "github.com/elazarl/go-bindata-assetfs"
	"github.com/webview/webview"
	"github.com/yamnikov-oleg/wingo"
	"net"
	"net/http"
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
func oncreate(w *wingo.Window, url string) {

	hand := w.GetHandle()
	wv := webview.NewWindow(true, unsafe.Pointer(&hand))
	wv.Navigate(url + "/static/index.html")
	wv.Run()
	defer wv.Destroy()
}

var traymenu *wingo.Menu

func myhttp(urlchan chan string) {
	mux := http.NewServeMux()

	files := assetfs.AssetFS{
		Asset:     data.Asset,
		AssetDir:  data.AssetDir,
		AssetInfo: data.AssetInfo,
		Prefix:    "",
	}
	mux.Handle("/", http.FileServer(&files))
	//mux.HandleFunc("/start", start)
	//mux.HandleFunc("/frame", getFrame)
	//mux.HandleFunc("/key", captureKeys)

	// get an ephemeral port, so we're guaranteed not to conflict with anything else
	listener, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		panic(err)
	}
	portAddress := listener.Addr().String()
	urlchan <- "http://" + portAddress
	listener.Close()
	server := &http.Server{
		Addr:    portAddress,
		Handler: mux,
	}
	server.ListenAndServe()
}

func main() {
	urlchan := make(chan string)
	go myhttp(urlchan)
	prefix := <-urlchan
	fmt.Println(prefix)
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

	oncreate(w, prefix)
	w.Show()

	wingo.Start()
}
