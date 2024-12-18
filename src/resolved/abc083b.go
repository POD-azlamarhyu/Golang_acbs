package main

import (
	"fmt"
	"strconv"
	"strings"
)
func calc_i2s_sum(i2s string) int{
	var sum int = 0
	slice_str := strings.Split(i2s,"")
	for i := 0; i<len(slice_str);i++{
		s2i, _ := strconv.Atoi(slice_str[i])
		sum += s2i
	}
	return sum
}

func main(){
	var a,b,n int

	fmt.Scanf("%d %d %d",&n,&a,&b)

	var sum int = 0
	for i := 1; i<=n; i++{
		i2s := strconv.Itoa(i)

		i2s_sum := calc_i2s_sum(i2s)

		if a <= i2s_sum && i2s_sum <= b{
			sum += i
		}
	}
	fmt.Println(sum)
}