package main


func nthUglyNumber(n int) int {
	if n == 1 {
		return 1
	}
	queue := make([]int, n)
	i := 0
	queue[i] = 1
	for {
		if i == n-1 {
			break
		}
		min := queue[i] * 5
		temp := queue[i] + 1
		for j := 0; j <= i; j++ {
			if t := queue[j] * 2; t > queue[i] {
				if t < min {
					min = t
				} else {
					continue
				}
			} else if t = queue[j] * 3; t > queue[i] {
				if t < min {
					min = t
				} else {
					continue
				}
			} else if t = queue[j] * 5; t > queue[i] {
				if t < min {
					min = t
				}
			}
			if temp == min {
				break
			}
		}
		i++
		queue[i] = min
	}
	return queue[i]
}

//暴力解法 时间太长
func nthUglyNumber_1(n int) int  {
	queue := make([]int, n)
	i := 0
	v := 1
	queue[i] = v
	for {
		if i == n-1 {
			break
		}
		v++
		if v%2 == 0 {
			if r := v / 2; r == 1 {
				i++
				queue[i] = v
			} else {
				index := i
				for ; index >= 0; index-- {
					if queue[index] == r {
						i++
						queue[i] = v
					} else if queue[index] < r {
						break
					}
				}
			}
		} else if v%3 == 0 {
			if r := v / 3; r == 1 {
				i++
				queue[i] = v
			} else {
				index := i
				for ; index >= 0; index-- {
					if queue[index] == r {
						i++
						queue[i] = v
					} else if queue[index] < r {
						break
					}
				}
			}
		} else if v%5 == 0 {
			if r := v / 5; r == 1 {
				i++
				queue[i] = v
			} else {
				index := i
				for ; index >= 0; index-- {
					if queue[index] == r {
						i++
						queue[i] = v
					} else if queue[index] < r {
						break
					}
				}
			}
		}
	}
	return queue[i]
}
