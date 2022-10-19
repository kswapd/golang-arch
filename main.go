package main

import (
	"fmt"
	"golang-arch/base"
	_ "golang-arch/httpclient"
)

func main() {

	/*httpclient.GetNextQuery()

	base.TestChan(1)
	base.TestGo()
	base.TestMath()
	c := make(chan int)
	go base.TestChanSend(c)
	go base.TestGoRecv(c)
	time.Sleep(time.Second * 3)
	close(c)

	base.TestStructInitialize()
	base.TestChanSend3()*/
	//proxy.ProxyHttp()
	//base.GoStructValue()
	base.GoInterface()
	fmt.Println("finished. ")

}
