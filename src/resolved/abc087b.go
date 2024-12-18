package main

import (
	"fmt"
)


func main(){
	var a, b, c,x int
	const coina int = 500
	const coinb int = 100
	const coinc int = 50
	fmt.Scanf("%d",&a)
	fmt.Scanf("%d",&b)
	fmt.Scanf("%d",&c)
	fmt.Scanf("%d",&x)

	var sum int = 0
	for i := 0; i <= a; i++{
		for j := 0; j<= b; j++{
			for k := 0; k<= c;k++{
				if x == coina*i+coinb*j+coinc*k{
					sum++
				}else{
					continue
				}
			}
		}
	}

	fmt.Println(sum)
}