package main

import (
	"context"
	"log"
	"syscall/js"

	"github.com/chromedp/chromedp"
)

var c chan bool

func init() {
	c = make(chan bool)
}

func main() {
	println("Hello ZiPIDA !!!!! Hello ZiPIDA !!!!! Hello ZiPIDA !!!!! Hello ZiPIDA !!!!! Hello ZiPIDA !!!!!")
	js.Global().Set("printMessage", js.FuncOf(printMessage))
	js.Global().Set("sum", js.FuncOf(sum))
	js.Global().Set("searchUrl", js.FuncOf(searchUrl))
	<-c
}

func printMessage(_ js.Value, i []js.Value) interface{} {
	callback := i[len(i)-1:][0]
	message := i[0].String()

	callback.Invoke(js.Null(), "Did you say "+message)

	return nil
}

func sum(_ js.Value, i []js.Value) interface{} {
	calc := i[0].Int() + i[1].Int()

	js.Global().Set("sum : ", js.ValueOf(calc))
	return js.ValueOf(calc).String()
}

// func temp() {
// 	println("FOO!")
// }

func searchUrl(_ js.Value, i []js.Value) interface{} {
	callback := i[len(i)-1:][0]
	targetURL := i[0].String()
	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.DisableGPU,
		chromedp.Flag("headless", true),
	)
	contextVar, cancelFunc := chromedp.NewExecAllocator(context.Background(), opts...)
	defer cancelFunc()

	contextVar, cancelFunc = chromedp.NewContext(contextVar)
	defer cancelFunc()

	var strVar string

	err := chromedp.Run(contextVar,
		chromedp.Navigate(targetURL),
		chromedp.InnerHTML(`body`, &strVar),
	)
	if err != nil {
		log.Println(err)
	}
	// js.Global().Set("searchUrl : ", js.ValueOf(strVar))

	callback.Invoke(js.Null(), "strVar "+strVar)
	return js.ValueOf(strVar).String()
}
