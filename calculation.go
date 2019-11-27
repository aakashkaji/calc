package calc

import "fmt"

func Factorial(n int64) {
	var val, i int64
	val = 1


	if n == 0 {
		fmt.Println("factorial of ", n, "is = ", 1)
	}else {
		for i = 1; i <= n; i++ {
			val = val*i

		}

		fmt.Println("factorial of ", n, "is = ", val)


	}
}


