package base

import (
	"fmt"
	"os"
	"os/signal"
	"reflect"
	"sync"
	"syscall"
	"time"
	_ "time"

	"k8s.io/apimachinery/pkg/util/wait"
)

func TestGoRecv(ch chan int) int {
	//out := <-ch

	//fmt.Println("TestGo", out)
	i := 1
	for v := range ch {
		fmt.Printf("GoRecv finish %d...\n", v)
		i += 1
		//fmt.Println(v)
	}
	/*
		fmt.Printf("GoRecv  %d start...\n",  <-ch)
		i += 1
		fmt.Printf("GoRecv %d start...\n",  <-ch)
		i += 1
		fmt.Printf("GoRecv %d start...\n", <-ch)
		i += 1
		fmt.Println("GoRecv exiting...")
	*/
	return 0
}
func TestChanSend(ch chan int) int {
	//c := make(chan int, n)
	i := 1
	//go func() {
	ch <- i
	fmt.Printf("GoSend %d finish...\n", i)
	i += 1
	ch <- i
	fmt.Printf("GoSend %d finish...\n", i)
	i += 1
	ch <- i
	fmt.Printf("GoSend %d finish...\n", i)
	fmt.Println("GoSend exiting...")
	//time.Sleep(2 * time.Second)
	//c <- 200
	//}()

	/*time.Sleep(1 * time.Second)
	fmt.Println(<-c)
	fmt.Println(<-c)*/
	/*for v := range c {
		fmt.Println(v)
	}*/
	// 保持持续运行
	//holdRun()
	return 0
}

func TestChanSend2(wg *sync.WaitGroup, ch chan int) chan int {

	//i := 1
	//go func() {
	//ch <- i+1
	fmt.Printf("GoSend %d finish2...\n", <-ch)
	/*i += 1
	ch <- i
	fmt.Printf("GoSend %d finish...\n", i)
	i += 1
	ch <- i
	fmt.Printf("GoSend %d finish...\n", i)*/
	fmt.Println("GoSend exiting2...")
	defer wg.Done()
	return ch
}

func TestChanSend3() {
	var wg sync.WaitGroup
	ch := make(chan int, 3)
	wg.Add(1)

	go TestChanSend2(&wg, ch)

	ch <- 7
	ch <- 8
	ch <- 9

	wg.Wait()

	fmt.Println("GoSend exiting3...")

	return
}

type person struct {
	name string
	age  int
}
type persons struct {
	Ps []person
}

func newPerson(name string) *person {

	p := person{name: name}
	p.age = 42
	return &p
}

func newPerson2(name string) person {

	p := person{name: name}
	p.age = 42
	return p
}

func TestStructInitialize() {

	var (
		s   person
		spp person
	)

	fmt.Println(person{"Bob", 20})

	fmt.Println(person{name: "Alice", age: 30})

	fmt.Println(person{name: "Fred"})

	fmt.Println(&person{name: "Ann", age: 40})

	fmt.Println(newPerson("Jon"))
	spp = newPerson2("Jon222")
	fmt.Println(spp)

	s = person{name: "Sean", age: 50}
	fmt.Println(s.name)

	sp := &s
	fmt.Println(sp.age)

	sp.age = 51
	fmt.Println(sp.age)

	primes := []int{2, 3, 5, 7, 11, 13}
	var pri []int = primes[1:4]
	fmt.Println(pri)
	primes[2] = 99
	fmt.Println(pri)

	var pri2 []int = make([]int, len(primes[1:4]))
	copy(pri2, primes[1:4])
	var pri3 *[]int
	pri3 = &pri2
	fmt.Println("len:", len(pri2))
	fmt.Println(pri2)
	primes[2] = 990
	fmt.Println(pri2)
	(*pri3)[2] = 1111
	fmt.Println(pri2)
	ps := persons{}
	fmt.Println("append info:%+v, %+v, %+v", ps, len(ps.Ps), ps.Ps == nil)

	ps.Ps = append(ps.Ps, person{
		name: "aa",
		age:  33,
	})

	fmt.Println("append info:%+v.", ps)

}

func TestArray() {
	//a := []string{}
	var a []string
	a = append(a, "aa")
	var arrb []string
	for _, s := range a {
		fmt.Println("result: ", s)
	}
	arrb = a
	a[0] = "dd"
	arrb = append(arrb, "mmm")
	for _, s := range a {
		fmt.Println("new resulta: ", s)
	}

	//copy(arrb, a)
	for _, s := range arrb {
		fmt.Println("new resultb: ", s)
	}

	copy(a, arrb)
	for _, s := range a {
		fmt.Println("new resulta: ", s)
	}

	b := []person{
		{
			name: "a",
			age:  22,
		},
		{
			name: "b",
			age:  23,
		},
	}

	c := []person{
		{
			name: "a",
			age:  22,
		},
		{
			name: "b",
			age:  24,
		},
	}

	d := map[string]string{
		"a": "b",
		"c": "d",
		"m": "n",
	}

	e := map[string]string{
		"a": "b",
		"c": "d",
	}

	fmt.Println("Equal check:", reflect.DeepEqual(b, c), reflect.DeepEqual(d, e))

}

func TestRepeat() {

	stop := make(chan struct{})
	// Define a function to be executed.
	fn := func() {
		// Do something.
		fmt.Println("Hello, world!")
	}

	fn2 := func() {
		// Do something.
		fmt.Println("Hello, world222!")
	}
	// Wait until the condition is met or the duration has elapsed.
	go wait.Until(fn, 3*time.Second, stop)
	go wait.Until(fn2, 3*time.Second, stop)

	fmt.Println("Waiting for stop signals...")
	// Wait for all goroutines to finish.
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	<-sigs
	close(stop)
	time.Sleep(2 * time.Second)
}
