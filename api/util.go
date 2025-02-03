package api

import "math"

const (
	ODD       = "odd"
	EVEN      = "even"
	ARMSTRONG = "armstrong"
)

func isPrime(n int) bool {
	if n <= 1 {
		return false
	}

	for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func isPerfect(n int) bool {
	sum := 1

	for i := 2; i*i <= n; i++ {
		if n%i == 0 {

			if i*i != n {
				sum = sum + i + n/i
			} else {
				sum += i
			}

		}
	}

	return sum == n && n != 1

}

func isOdd(num int) bool {
	return num%2 != 0
}

func order(n int) int {
	ord := 0

	for n != 0 {
		ord++
		n = n / 10
	}
	return ord
}

func digitSum(n int) int {
	sum := 0
	temp := n

	if n < 0 {
		temp = -1 * temp

	}

	for temp != 0 {
		r := temp % 10
		sum += r
		temp = temp / 10
	}

	return sum
}

func isArmstrong(n int) bool {
	sum := 0
	ord := order(n)
	temp := n

	for temp != 0 {
		r := temp % 10
		sum += int(math.Pow(float64(r), float64(ord)))
		temp = temp / 10
	}

	return sum == n

}

func isEven(num int) bool {
	return num%2 == 0
}
