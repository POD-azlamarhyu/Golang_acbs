package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {

	var n int
	fmt.Scanf("%d\n", &n)

	sc := bufio.NewScanner(os.Stdin)

	var t []int
	var x []int
	var y []int

	for i := 0; i < n; i++ {
		sc.Scan()
		inputs := strings.Split(sc.Text(), " ")
		tt, _ := strconv.Atoi(inputs[0])
		xx, _ := strconv.Atoi(inputs[1])
		yy, _ := strconv.Atoi(inputs[2])
		t = append(t, tt)
		x = append(x, xx)
		y = append(y, yy)
	}

	t0 := 0
	x0 := 0
	y0 := 0

	flag := true
	for j := 0; j < n; j++ {
		dist:= int(math.Abs(float64(x0-x[j])) + math.Abs(float64(y0-y[j])))
		dt := t[j] - t0
		if dist > dt {
			flag = false
			break
		}
		if (dt-dist) %2  == 1 {
			flag = false
			break
		}
		t0 = t[j]
		x0 = x[j]
		y0 = y[j]
	}
	if flag {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
