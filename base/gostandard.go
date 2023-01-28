package base

import (
	"fmt"
	"time"
)

type Operation int

const (
	Add Operation = iota + 1
	Subtract
	Multiply
)

type F interface {
	f()
}
type S1 struct{}

func (s S1) f() {
}

type S2 struct{}

func (s *S2) f() {}
func MyInterface() {
	// f1.f() 无法修改底层数据
	// f2.f() 可以修改底层数据，给接口变量 f2 赋值时使用的是对象指针
	var f1 F = S1{}
	var f2 F = &S2{}
	var f3 F = &S1{}
	//cannot use (S2 literal) (value of type S2) as F value in variable declaration: S2 does not implement F (method f has pointer receiver)
	//var f3 F = S2{}

	fmt.Printf("type:%T, %T, %T\n", f1, f2, f3)

}

type S struct {
	data1 string
	data2 string
}

func (s S) Read() string {
	fmt.Printf("Read:%s\n", s.data1)
	return s.data1
}

func (s S) VWrite(str string) S {
	fmt.Printf("VWrite:%s-->%s\n", s.data1, str)
	s.data1 = str
	return s
}

func (s *S) PWrite(str string) (S, *S) {
	fmt.Printf("PWrite:%s-->%s\n", s.data1, str)
	s.data1 = str
	return *s, s
}

func MyStructValue() {

	//too few values in struct{data1 string; data2 string}{…}
	sVals := map[int]S{1: {}}
	fmt.Printf("sVals:%+v\n", sVals)
	sVals[1].Read()
	sVals[1].VWrite("A")
	//cannot call pointer method PWrite on S
	//sVals[1].PWrite("ddd")
	sVals[1].Read()
	//invalid operation: cannot take address of (sVals[1]) (map index expression of type S)
	//(&(sVals[1])).PWrite("D")
	s3 := sVals[1]
	s3.PWrite("D")
	sVals[1].Read()

	arr := []S{S{}}
	s4, s6 := arr[0].PWrite("DDD")
	s4.VWrite("FFF")
	s5 := arr[0].VWrite("EEE")

	arr[0].Read()
	s4.Read()
	s5.Read()
	fmt.Printf("Struct addr:%p, (%p,%p) %p.\n", &arr[0], &s4, s6, &s5)

	s2 := S{}
	fmt.Printf("s2 type: %T\n", s2)
	s2.Read()
	s2.PWrite("B")
	s2.VWrite("C")
	s2.Read()

}

func MySlice() {
	arr1 := []int{1, 2, 3, 4, 5}
	arr2 := arr1
	arr3 := make([]int, len(arr1))
	copy(arr3, arr1)
	arr1[0] = 100
	fmt.Printf("arr1: %+v, arr2: %+v, arr3: %+v, %p, %p, %p.\n", arr1, arr2, arr3, arr1, arr2, arr3)
}

type E1 struct {
	s1 string
}

func (e *E1) Read() {
	fmt.Printf("E1 Read:%s\n", e.s1)
}

type E2 struct {
	E1
	E3 struct {
		s3 string
	}
}

func (e *E2) Read() {
	fmt.Printf("E2 Read:%s, %s\n", e.s1, e.E3.s3)
}
func MyEmbeddedStruct() {
	e2 := E2{
		E1: E1{
			s1: "aaa",
		},
	}
	e2.s1 = "sss"
	e2.E3.s3 = "ddd"
	e2.Read()

}

func myflush() {
	fmt.Println("flush")
}
func MyGoroutine() {
	stop := make(chan struct{})
	testChan := make(chan struct{}, 1)
	done := make(chan struct{})

	go func(delay time.Duration) {
		/*for {
			myflush()
			time.Sleep(delay * time.Millisecond)
		}*/

		defer close(done)
		for {
			select {
			case <-stop:
				fmt.Println("Get stop signal, exit.")
				return
			case <-testChan:
				fmt.Println("testChan read.")
			//case testChan <- struct{}{}:
			//	fmt.Println("testChan write.")
			default:
				myflush()
				time.Sleep(delay * time.Millisecond)
			}
		}
		//close(stop)
		//done <- struct{}{}

	}(100)
	time.Sleep(2 * time.Second)
	stop <- struct{}{}
	<-done
	return
}
