package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func createArrayloopInput(n int) []int{
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	line := strings.Fields(sc.Text())

	a := make([]int, n)
	for i := range line {
		a[i], _ = strconv.Atoi(line[i])
	}

	return a
}

func createArraySplitSpace(n int) []int{
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	line := strings.Split(sc.Text()," ")
	var a []int
	for _, input := range line{
		p, _ := strconv.Atoi(input)
		a = append(a, p)
	}
	return a
}

func create2dimArraySplitSpace(n int) [][]int{
	sc := bufio.NewScanner(os.Stdin)
	
	var a [][]int

	for i := 0; i<n; i++{
		sc.Scan()
		line := strings.Split(sc.Text()," ")
		for j := 0; j < len(line); j++{
			p, _ := strconv.Atoi(line[j])
			a[i] = append(a[i], p)
		}
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