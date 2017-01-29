package main

func fibo(n int) (res int) {
	if n <= 1 {
		res = 1
	} else {
		res = fibo(n-1) + fibo(n-2)
	}
	return
}
