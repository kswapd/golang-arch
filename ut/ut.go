package ut

import (
	"fmt"
	"math/big"
	"sync"
	"time"
)

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

func bigFactorial(n int) *big.Float {
	var precision uint = 5000
	if n == 0 {
		return new(big.Float).SetPrec(precision).SetFloat64(1)
	}
	return new(big.Float).SetPrec(precision).Mul(new(big.Float).SetPrec(precision).SetFloat64(float64(n)), bigFactorial(n-1))
}

type testResult struct {
	name     string
	args     []string
	expected string
	wantErr  bool
}

func blockingMultipleFunction(tr *testResult, wg *sync.WaitGroup) error {
	defer wg.Done()
	fmt.Printf("result %v.\n", tr)
	return nil
}

func StartMyUT() {

	var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

	for iv := range pow {
		fmt.Printf("%d", iv)
	}
	var wg sync.WaitGroup
	tests := []testResult{
		/*{
			name:     "ui",
			args:     []string{"ui"},
			expected: "processed\n",
			wantErr:  false,
		},*/
		{
			name:     "server cmd",
			args:     []string{"server", "-c", "./configs/server_ut.yaml"},
			expected: "processed: server\n",
			wantErr:  false,
		},
		{
			name:     "controller cmd",
			args:     []string{"controller", "-c", "./configs/k8s_config_dev"},
			expected: "processed: controller\n",
			wantErr:  false,
		},
	}
	/*for i := 0; i < len(tests); i++ {
		wg.Add(1)
		go worker(&wg)
	}*/
	for i := range tests {
		wg.Add(1)
		fmt.Printf("i:%d.\n", i)
		go blockingMultipleFunction(&tests[i], &wg)
	}

	done := make(chan error)
	go func() {
		wg.Wait()
		close(done)
	}()
	select {
	case err := <-done:
		if err != nil {
			fmt.Printf("blockingMultipleFunction returned an error: %v", err)
		} else {
			fmt.Printf("blockingMultipleFunction all done.\n")
		}
	case <-time.After(5 * time.Second):
		fmt.Printf("blockingMultipleFunction timed out")
	}

}
