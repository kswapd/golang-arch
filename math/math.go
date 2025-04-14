package math

import (
	"fmt"
	"math"
	"math/big"

	"github.com/ALTree/bigfloat"
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

func StartMyPiCalc() {
	var index1 = new(big.Float).SetFloat64(2 * math.Sqrt2 / (99 * 99))
	var result = new(big.Float).SetFloat64(0)
	var precision uint = 5000
	//var loopTimes = 300
	var times = 10
	var printInterval = 1
	//for times := 1; times <= loopTimes; times++ {
	result = new(big.Float).SetPrec(precision).SetFloat64(0)
	for k := 0; k < times; k++ {
		/*var temp1 = bigFactorial(4 * k)
		//var temp2 = new(big.Float).SetPrec(precision).SetFloat64(math.Pow(float64(factorial(k)), 4))

		var temp2 = bigfloat.Pow(bigFactorial(k), new(big.Float).SetPrec(precision).SetFloat64(4))
		//var temp33 = new(big.Float).SetPrec(precision).SetFloat64(float64(26390*k+1103) / math.Pow(396, float64(4*k)))
		var temp3 = new(big.Float).SetPrec(precision).Quo(new(big.Float).SetPrec(precision).SetFloat64(float64(26390*k+1103)), bigfloat.Pow(new(big.Float).SetPrec(precision).SetFloat64(396), new(big.Float).SetPrec(precision).SetFloat64(float64(4*k))))

		fmt.Printf("times %3d: %e, %e %e.\n", k, temp1, temp2, temp3)
		var temp = new(big.Float).SetPrec(precision).Mul(new(big.Float).SetPrec(precision).Quo(temp1, temp2), temp3)*/

		/*var temp = new(big.Float).SetPrec(precision).Quo(
			new(big.Float).SetPrec(precision).Quo(
				new(big.Float).SetPrec(precision).Mul(
					bigFactorial(4*k),
					new(big.Float).SetPrec(precision).SetFloat64(float64(26390*k+1103))
				),
				bigfloat.Pow(bigFactorial(k), new(big.Float).SetPrec(precision).SetFloat64(4))
			),
			bigfloat.Pow(new(big.Float).SetPrec(precision).SetFloat64(396), new(big.Float).SetPrec(precision).SetFloat64(float64(4*k)))
		)*/

		var temp = new(big.Float).SetPrec(precision).Quo(
			new(big.Float).SetPrec(precision).Quo(
				new(big.Float).SetPrec(precision).Mul(
					bigFactorial(4*k),
					new(big.Float).SetPrec(precision).SetFloat64(float64(26390*k+1103))),
				bigfloat.Pow(bigFactorial(k), new(big.Float).SetPrec(precision).SetFloat64(4))),
			bigfloat.Pow(new(big.Float).SetPrec(precision).SetFloat64(396), new(big.Float).SetPrec(precision).SetFloat64(float64(4*k))))

		result = new(big.Float).SetPrec(precision).Add(result, temp)
		var pi = new(big.Float).SetPrec(precision).Quo(new(big.Float).SetPrec(precision).SetFloat64(1), new(big.Float).SetPrec(precision).Mul(result, index1))
		if k%printInterval == 0 {
			fmt.Printf("times %3d: %e, %e %e.\n", k, pi, new(big.Float).SetPrec(precision).Sub(pi, new(big.Float).SetPrec(precision).SetFloat64(math.Pi)), temp)
		}
	}
	//var pi = 1 / new(big.Float).SetFloat64(result*index1)

	//}

}

func StartMyPiCalcBk() {
	var index1 = new(big.Float).SetFloat64(2 * math.Sqrt2 / (99 * 99))
	var result = new(big.Float).SetFloat64(0)
	var precision uint = 5000
	//var loopTimes = 300
	var times = 10
	var printInterval = 1
	//for times := 1; times <= loopTimes; times++ {
	result = new(big.Float).SetPrec(precision).SetFloat64(0)
	for k := 0; k < times; k++ {
		/*var temp1 = bigFactorial(4 * k)
		//var temp2 = new(big.Float).SetPrec(precision).SetFloat64(math.Pow(float64(factorial(k)), 4))

		var temp2 = bigfloat.Pow(bigFactorial(k), new(big.Float).SetPrec(precision).SetFloat64(4))
		//var temp33 = new(big.Float).SetPrec(precision).SetFloat64(float64(26390*k+1103) / math.Pow(396, float64(4*k)))
		var temp3 = new(big.Float).SetPrec(precision).Quo(new(big.Float).SetPrec(precision).SetFloat64(float64(26390*k+1103)), bigfloat.Pow(new(big.Float).SetPrec(precision).SetFloat64(396), new(big.Float).SetPrec(precision).SetFloat64(float64(4*k))))

		fmt.Printf("times %3d: %e, %e %e.\n", k, temp1, temp2, temp3)
		var temp = new(big.Float).SetPrec(precision).Mul(new(big.Float).SetPrec(precision).Quo(temp1, temp2), temp3)*/

		/*var temp = new(big.Float).SetPrec(precision).Quo(
			new(big.Float).SetPrec(precision).Quo(
				new(big.Float).SetPrec(precision).Mul(
					bigFactorial(4*k),
					new(big.Float).SetPrec(precision).SetFloat64(float64(26390*k+1103))
				),
				bigfloat.Pow(bigFactorial(k), new(big.Float).SetPrec(precision).SetFloat64(4))
			),
			bigfloat.Pow(new(big.Float).SetPrec(precision).SetFloat64(396), new(big.Float).SetPrec(precision).SetFloat64(float64(4*k)))
		)*/

		var temp = new(big.Float).SetPrec(precision).Quo(
			new(big.Float).SetPrec(precision).Quo(
				new(big.Float).SetPrec(precision).Mul(
					bigFactorial(4*k),
					new(big.Float).SetPrec(precision).SetFloat64(float64(26390*k+1103))),
				bigfloat.Pow(bigFactorial(k), new(big.Float).SetPrec(precision).SetFloat64(4))),
			bigfloat.Pow(new(big.Float).SetPrec(precision).SetFloat64(396), new(big.Float).SetPrec(precision).SetFloat64(float64(4*k))))

		result = new(big.Float).SetPrec(precision).Add(result, temp)
		var pi = new(big.Float).SetPrec(precision).Quo(new(big.Float).SetPrec(precision).SetFloat64(1), new(big.Float).SetPrec(precision).Mul(result, index1))
		if k%printInterval == 0 {
			fmt.Printf("times %3d: %e, %e %e.\n", k, pi, new(big.Float).SetPrec(precision).Sub(pi, new(big.Float).SetPrec(precision).SetFloat64(math.Pi)), temp)
		}
	}
	//var pi = 1 / new(big.Float).SetFloat64(result*index1)

	//}

}

func StartMyPiCalc2() {
	var index1 float64 = 2 * math.Sqrt2 / (99 * 99)
	var result float64 = 0
	var loopTimes = 300
	for times := 1; times <= loopTimes; times++ {
		for k := 0; k < times; k++ {
			var temp1 = factorial(4 * k)
			var temp2 float64 = math.Pow(float64(factorial(k)), 4)
			var temp3 float64 = float64(26390*k+1103) / math.Pow(396, float64(4*k))
			var temp float64 = float64(temp1) / temp2 * temp3
			result += temp
		}
		var pi float64 = 1 / (result * index1)
		fmt.Printf("times %3d: %.24f, %.24f.\n", times, pi, (math.Pi - pi))

		result = 0
	}

}

func StartMyPiCalc333() {
	var index1 float64 = 2 * math.Sqrt2 / (99 * 99)
	var result float64 = 0
	var loopTimes = 300

	for times := 1; times <= loopTimes; times++ {
		result = 0 // 将 result 初始化放在外层循环内部
		for k := 0; k < times; k++ {
			// 缓存 factorial(4 * k) 和 math.Pow(float64(factorial(k)), 4) 的结果
			fact4k := factorial(4 * k)
			factK := float64(factorial(k))
			powFactK4 := math.Pow(factK, 4)

			temp3 := float64(26390*k+1103) / math.Pow(396, float64(4*k))
			temp := float64(fact4k) / powFactK4 * temp3
			result += temp
		}
		var pi float64 = 1 / (result * index1)
		fmt.Printf("times %3d: %.24f, %.24f.\n", times, pi, (math.Pi - pi))
	}
}
