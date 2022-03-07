package main

import "syscall/js"

var c chan bool

func init() {
	c = make(chan bool)
}

func main() {
	println("Hello ZiPIDA !!!!! Hello ZiPIDA !!!!! Hello ZiPIDA !!!!! Hello ZiPIDA !!!!! Hello ZiPIDA !!!!!")
	js.Global().Set("printMessage", js.FuncOf(printMessage))
	js.Global().Set("sum", js.FuncOf(sum))
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

func temp() {
	println("FOO!")
}
