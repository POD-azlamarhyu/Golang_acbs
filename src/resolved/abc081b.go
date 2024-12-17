package main

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
	"strings"
)


func is_array_even(a []int,n int) bool{
	flag := true
	for i := 0; i < n; i++{
		if a[i] % 2 == 1 {
			flag = false
			break
		}else{
			continue
		}
	}
	return flag
}

func main(){
	var n int

	fmt.Scanf("%d",&n)

	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	line := strings.Fields(sc.Text())

	a := make([]int, n)
	for i := range line {
		a[i], _ = strconv.Atoi(line[i])
	}

	sum := 0
	for {
		if !is_array_even(a,n){
			break
		}else{
			for i := 0; i < n; i++{
				a[i] =a[i]/2
			}
			
		}
		sum ++
	}

	fmt.Println(sum)
}