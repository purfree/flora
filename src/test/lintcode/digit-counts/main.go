/*
 http://www.lintcode.com/zh-cn/problem/digit-counts/
 计算数字k在0到n中的出现的次数，k可能是0~9的一个值

n
k在个位出现的次数c
a = n/10
b = n%10
if b >= k
	c = a + 1
else
	c = a

k在十位出现的次数
a = n/100 * 10
b = n%100
d = b%10
e = b/10
if e > k
	c = a + 10
if e < k
	c = a
if e == k
	c = a + d + 1

k在百位出现的次数
a = n/1000 * 100
b = n%1000
d = b%100
e = b/100
if e > k
	c = a + 100
if e < k
	c = a
if e == k
	e = a + d + 1

m为n的位数
k!=0的公式
x = 10^(m+1)
y = 10^m
a = n/x*y
b = n/x
d = b%y
e = b/y
if e > k
	c = a + y
if e < k
	c = a
if e ==k
	c = a + d + 1


k==0
个位0出现次数
n/10+1
其他位0出现的次数
x = 10^(m+1)
y = 10^m
a = (n/x - 1) * y
d = n % x
e = d / y
if e > k
	c = a + y
if e < k
	c = a
if e == k
	c = a + d + 1

*/
package main

import (
	"fmt"
	"math"
)

func main() {
	//fmt.Println(digitCounts(1, 1114))
	//fmt.Println(CalculationPlace(10))
	//fmt.Println(CalculationPlace(100))
	//fmt.Println(CalculationPlace(1000))
	//fmt.Println(math.Pow10(-1))
	//fmt.Println(CalculationCountsByPlace(0, 119, 0))
	//fmt.Println(CalculationCountsByPlace(0, 119, 1))
	//fmt.Println(CalculationCountsByPlace(0, 119, 2))

	fmt.Println(digitCounts(0, 112))
	fmt.Println(digitCounts_1(0, 112))

}

func digitCounts_1(k int, n int) int {
	c := 0
	for i := 0; i <= n; i++ {
		c = c + CalculationTimes(k, i)
	}
	return c
}

func CalculationTimes(k int, n int) int {
	c := 0
	for {
		r := n % 10
		if k == r {
			c++
		}
		n = n / 10
		if n == 0 {
			break
		}
	}
	return c
}

//计算数字的位数
func CalculationPlace(n int) int {
	return int(math.Log10(float64(n)))
}

func CalculationCountsByPlace(k int, n int, p int) int {
	c := 0
	x := int(math.Pow10(p + 1))
	y := int(math.Pow10(p))
	a := n / x * y
	b := n % x
	d := b % y
	e := b / y
	if e > k {
		c = a + y
	} else if e < k {
		c = a
	} else if e == k {
		c = a + d + 1
	}
	return c
}

func CalculationCounts0ByPlace(k int, n int, p int) int {
	c := 0
	x := int(math.Pow10(p + 1))
	y := int(math.Pow10(p))
	a := (n/x - 1) * y
	d := n % x
	e := d / y
	if e > k {
		c = a + y
	} else if e < k {
		c = a
	} else if e == k {
		c = a + d + 1
	}
	return c
}

func digitCounts(k int, n int) int {
	p := CalculationPlace(n)
	c := 0
	if k == 0 {
		c = n/10 + 1
		for i := 1; i < p; i++ {
			c = c + CalculationCounts0ByPlace(k, n, i)
		}
	} else {
		for i := 0; i <= p; i++ {
			c = c + CalculationCountsByPlace(k, n, i)
		}
	}
	return c
}
