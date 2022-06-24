package main

//go:generate go run github.com/webview/webview/cmd/webview-go fetch-libs --set-env
//go:generate windres -o resources.syso resources.rc

import "github.com/webview/webview"

func main() {
	w := webview.New(false)
	defer w.Destroy()
	w.SetTitle("Basic Example")
	w.SetSize(480, 320, webview.HintNone)
	w.SetHtml("Thanks for using webview!")
	w.Run()
}
