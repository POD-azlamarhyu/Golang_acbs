package main

import (
	"fmt"
)



func main(){
	var n int
	

	fmt.Scanf("%d",&n)
	var d []int = make([]int, n)


	for i := 0; i < n; i++{
		fmt.Scanf("%d",&d[i])
	}


	for i := 0; i<len(d);i++{
        for j := 0; j < len(d)-i-1;j++{
            if d[j] < d[j+1]{
                temp := d[j]
                d[j] = d[j+1]
                d[j+1] = temp
            }
        }
    }
	
	var sum int = 1
	for i := 0; i<len(d)-1;i++{
		if d[i] > d[i+1] {
			sum += 1
		}
	}
	fmt.Println(sum)
}