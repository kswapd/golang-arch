package base

import (
	"fmt"
	//"encoding/json"
	mr "golang-arch/rect"
	"math"
	"time"
	//"strconv"
)

func goProducer(ch chan int) {
	var data int
	fmt.Println("producing...")
	for n := 0; n < 5; n++ {
		//time.Sleep(time.Second*1)
		data = n * 10
		ch <- data
		fmt.Println("produce:", data)
	}
}
func goConsumer(ch chan int) {
	fmt.Println("consuming...")
	var data int
	var ok bool
	_ = ok
	/*for {
		//time.Sleep(time.Second*1)
		data,ok = <- ch
		fmt.Println("consume:", data)
		if ok == false{
			fmt.Println("Channel is closed, exit.")
			break
		}
	}*/
	for v := range ch {
		data = v
		fmt.Println("consume:", data)
	}
	//fmt.Println("consume:", data)
	fmt.Println("Channel is closed, exit.")
}

func TestChan(n int) int {
	//c := make(chan int, 1)
	c := make(chan int, 1)
	var data int
	//go func() {
	c <- 48
	data = <-c
	fmt.Println("TestGo", data)
	c <- 96

	data = <-c
	fmt.Println("TestGo", data)
	//time.Sleep(2 * time.Second)
	c <- 200
	data = <-c
	fmt.Println("TestGo", data)
	//}()

	time.Sleep(1 * time.Second)
	//for v := range c {
	//	fmt.Println(v)
	//}

	// 保持持续运行
	//holdRun()
	return 0
}

func TestPack() {
	var width, height = 10, 20
	area := mr.Area(width, height)
	fmt.Printf("area of rect (%d x %d ) is %d.\n", width, height, area)
}

func TestMath() {
	var c1 complex128 = 2 + 5i
	c2 := 3 + 8i
	c3 := c1 + c2
	var (
		c4, c5 = 6, math.Sqrt(144)
		s1     = "hi,goworld"
	)
	const ca = 16

	fmt.Printf("get result:%#v, %T ,%v,%v,%v,%v,%T\n", c3, c3, c4, c5, ca, s1, ca)

	var i = 5
	var f = 5.6
	var c = float64(i) + f
	fmt.Println(c)
}

func TestGo() {

	var ch = make(chan int, 0)
	go goProducer(ch)
	go goConsumer(ch)
	time.Sleep(time.Second * 3)
	close(ch)
	time.Sleep(time.Second * 3)
}
