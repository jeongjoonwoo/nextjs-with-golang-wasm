package main

import (
	"context"
	"fmt"
	"log"
	"syscall/js"

	"github.com/chromedp/chromedp"
)

var c chan bool

func init() {
	c = make(chan bool)
}

func main() {
	fmt.Println("Hello ZiPIDA !!!!! Hello ZiPIDA !!!!! Hello ZiPIDA !!!!! Hello ZiPIDA !!!!! Hello ZiPIDA !!!!!")
	js.Global().Set("printMessage", js.FuncOf(printMessage))
	js.Global().Set("sum", js.FuncOf(addFunction))
	js.Global().Set("searchUrl", js.FuncOf(searchUrl))
	<-c
}

func printMessage(_ js.Value, i []js.Value) interface{} {
	callback := i[len(i)-1:][0]
	message := i[0].String()

	callback.Invoke(js.Null(), "Did you say "+message)

	return nil
}

func addFunction(this js.Value, p []js.Value) interface{} {
    sum := p[0].Int() + p[1].Int()
    return js.ValueOf(sum)
}

func searchUrl(_ js.Value, i []js.Value) interface{} {
	// var html string
	// callback := i[0]
	// targetURL := callback.String()
	
	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.Flag("headless", false),
	)

	execCtx, _ := chromedp.NewExecAllocator(context.Background(), opts...)
	// defer execCancel()

	ctx, _ := chromedp.NewContext(execCtx)
	// defer cancel()

	err := chromedp.Run(ctx)
	if err != nil {
		log.Println(err)
		return js.ValueOf(err.Error())
	}
	
	return js.ValueOf("done")
}
