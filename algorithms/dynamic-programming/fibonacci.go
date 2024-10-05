package main

func main() {
	n := 15
	memo := make([]int, n)
	println(Fib(n))
	println(Fib2(n, memo))
	println(Fib2(n, make([]int, n+1)))
}

//1 1 2 3 5 8 ...
//O(2 ^n)
func Fib(n int) int {
	if n <= 1 {
		return 1
	}
	return Fib(n-1) + Fib(n-2)
}

//Memoized solution
// 0 -> 1
// 1 -> 1
// 2 -> 2
// 3 -> 3
// 4 -> 5
// [0, 2, 3, 5]
func Fib2(n int, memo []int) int {
	if n <= 1 {
		return 1
	}

	if memo[n-1] != 0 {
		return memo[n-1]
	}

	res := Fib2(n-1, memo) + Fib2(n-2, memo)

	memo[n-1] = res

	return res
}

//Bottom up
func Fib3(n int, k []int) int {
	if n <= 1 {
		return 1
	}

	k[0] = 1
	k[1] = 1

	for i := 2; i <= n; i++ {
		k[i] = k[i-1] + k[i-2]
	}
	return k[n]
}
