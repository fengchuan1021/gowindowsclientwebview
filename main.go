package main
import "github.com/yamnikov-oleg/wingo"
func  onclose(w *wingo.Window) bool{
	w.Hide()
	return false
}
func ontrayclick(w *wingo.Window){
	w.Show()
}
func ontrayrightclick(w *wingo.Window){
traymenu.StartContext(w)

}
var traymenu *wingo.Menu

func main(){
	w:=wingo.NewWindow(true,true)
	size:=wingo.Vector{500,500}
	w.SetSize(size)
	w.OnClose=onclose
	w.OnTrayClick=ontrayclick
	w.OnTrayRightClick=ontrayrightclick
	icon:=wingo.LoadIcon(101)
	w.SetIcon(icon)
	w.AddTrayIcon(icon,"hellofromaiel")
	traymenu=wingo.NewContextMenu()
	exitbtn:=traymenu.AppendItemText("退出")
	exitbtn.OnClick= func(item *wingo.MenuItem) {
		w.Destroy()
	}
	w.Show()
	wingo.Start()
}
