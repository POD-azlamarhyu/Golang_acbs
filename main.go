package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func createArray(n int) []int{
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	line := strings.Fields(sc.Text())

	a := make([]int, n)
	for i := range line {
		a[i], _ = strconv.Atoi(line[i])
	}

	return a
}

func main(){
	var a int
	var b, c int
	var s string

	fmt.Scanf("%d",&a)
	fmt.Scanf("%d %d",&b,&c)
	fmt.Scanf("%s",&s)

	fmt.Println(a+b+c, s)
}