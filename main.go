package main

import (
	"awesomeProject1/data"
	"fmt"
	assetfs "github.com/elazarl/go-bindata-assetfs"
	"github.com/webview/webview"
	"github.com/yamnikov-oleg/wingo"
	"net"
	"net/http"
	"runtime"
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

var mainview webview.WebView

func onresize(w *wingo.Window, xy wingo.Vector) {
	if nil != mainview {
		mainview.SetSize(xy.X, xy.Y, webview.HintNone)

	}
}
func realopenurl(url string) {
	runtime.LockOSThread()
	w := webview.New(true)
	defer w.Destroy()
	w.Navigate(url)
	w.Run()

}
func openurl(url string) int {
	fmt.Println(url)
	go realopenurl(url)
	return 1
}
func oncreate(w *wingo.Window, url string) {

	hand := w.GetHandle()
	mainview = webview.NewWindow(true, unsafe.Pointer(&hand))

	//mainview=&wv
	mainview.Navigate(url + "/static/index.html")
	mainview.Bind("open_url", openurl)
	mainview.Run()
	defer mainview.Destroy()
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
	w.OnSizeChanged = onresize
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
