/*
 http://www.lintcode.com/zh-cn/problem/trailing-zeros/
 设计一个算法，计算出n阶乘中尾部零的个数

 思路：
	一个数 n 的阶乘末尾有多少个 0 取决于从 1 到 n 的各个数的因子中 2 和 5 的个数
	而 2 的个数是远远多余 5 的个数的, 因此求出 5 的个数即可
	例如, 100/5 = 20, 20/5 = 4, 4/5 = 0
	100中包含5,10,15,20...100,共100/5=20
	但是25=5*5,含有2个5,100中包含25,50,75,100,共100/5/5=4
 	以此类推125=5*5*5,625=5*5*5*5...,直到结果为0
	所以100中包含5的个数为20+4=24
*/
package main

import "fmt"

func main() {
	//a := big.NewInt(1)
	//for i := 1; i <= 26; i++ {
	//	a.Mul(a, big.NewInt(int64(i)))
	//	fmt.Printf("%v\t%v\n", i, a)
	//}
	//fmt.Println(a.String())
	fmt.Println(trailingZeros(100))
	//
	//a := big.NewInt(1)
	//b := big.NewInt(2)
	//c := big.NewInt(3)
	////d := a.Mul(b, c)
	//d := big.NewInt(11).Mul(b, c)
	//fmt.Println(a, b, c, d)
}

func trailingZeros(n int64) int64 {
	z := int64(0)
	for {
		n = n / 5
		z = z + n
		if n == 0 {
			return z
		}
	}
}
