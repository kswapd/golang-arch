package base

import (
	"context"
	"fmt"
	"time"
)

func doSomething(ctx context.Context) {
	//ctx, cancelCtx := context.WithCancel(ctx)
	ctx, cancelCtx := context.WithTimeout(ctx, 1500*time.Millisecond)

	printCh := make(chan int)
	fmt.Printf("doSomething: myKey's value is %s\n", ctx.Value("myKey"))
	go doAnother(ctx, printCh)

	/*for num := 1; num <= 3; num++ {
		printCh <- num
	}*/
	for num := 1; num <= 3; num++ {
		select {
		case printCh <- num:
			time.Sleep(1 * time.Second)
		case <-ctx.Done():
			break
		}
	}

	defer cancelCtx()

	time.Sleep(100 * time.Millisecond)

	fmt.Printf("doSomething: finished\n")
}

func doAnother(ctx context.Context, printCh <-chan int) {
	for {
		select {
		case <-ctx.Done():
			if err := ctx.Err(); err != nil {
				fmt.Printf("doAnother err: %s\n", err)
			}
			fmt.Printf("doAnother: finished\n")
			return
		case num := <-printCh:
			fmt.Printf("doAnother: %d\n", num)
		}
	}
}

func TestContext() {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "myKey", "myValue")
	doSomething(ctx)
}
