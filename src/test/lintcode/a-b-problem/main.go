/*
 http://www.lintcode.com/zh-cn/problem/a-b-problem/
 给出两个整数a和b, 求他们的和, 但不能使用 + 等数学运算符。

 思路：
	与运算，获得需要进位的位，左移进位
	异或运算，相当于加运算，但不包括进位运算
*/
package main

import "fmt"

func main() {
	a, b := 0, 0
	for i := 0; i <= 100; i++ {
		b = aplusb(i, b)
		a += i
	}
	fmt.Println(a, b)
}

func aplusb(a int, b int) int {
	for {
		a, b = a&b<<1, a^b
		if a == 0 {
			return b
		}
	}
}

func aplusb2(a, b int) int {
	if a == 0 {
		return b
	}
	a, b = a&b<<1, a^b
	return aplusb2(a, b)
}
