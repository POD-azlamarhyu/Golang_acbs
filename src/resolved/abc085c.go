package main

import (
	"fmt"
)

func main(){
	var n,y int
	const a int = 10000
	const b int = 5000
	const c int = 1000

	fmt.Scanf("%d %d",&n,&y)
	var v int
	var w int
	var x int
	var isStopIloop bool =false
	for i := 0; i <= n; i++{
		for j := 0; j<=n; j++{
			z := y -  (a*i + b*j)
			if z < 0{
				break
			}
			q := z % c
			k := z / c
			if q == 0 && i+j+k == n {
				v = i
				w = j
				x = k
				isStopIloop = true
				break
			}
		}
		if isStopIloop{
			break
		}
	}

	if isStopIloop {
		fmt.Println(v,w,x)
	}else{
		fmt.Println(-1,-1,-1)
	}
}