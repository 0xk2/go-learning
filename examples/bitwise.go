package examples

import (
	"fmt"
)

func BitWise() {
	var n1, i int
	n1 = 200 // 255=128+64+32+16+8+4+2+1 - 11111111

	for i = 1; i <= 8; i++ {
		clearNSmallestBit(n1, i)
	}
	for i = 1; i <= 8; i++ {
		keepNSmallestBit(n1, i)
	}
}

// clear n bit
func clearNSmallestBit(num int, n int) int {
	r := (num >> n) << n
	fmt.Println(num, "clear ", n, " smallest bit, turn into", r)
	return r
}

func keepNSmallestBit(num int, n int) int {
	mask := 1<<n - 1 // e.g: n=6 -> 1<<6= 1000000b, 1<<6-1 = 0111111b
	r := num & mask
	fmt.Println(num, "keep ", n, " smallest bit, mask is ", mask, "; turn into", r)
	return r
}
