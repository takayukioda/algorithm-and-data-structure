package main

import "fmt"

func main() {
	m, n := 18, 30
	fmt.Printf("%d と %d の最大公約数は %d\n", m, n, gcd(m, n))
}

func gcd(m, n int) int {
	if n == 0 {
		return m
	}
	return gcd(n, m%n)
}
