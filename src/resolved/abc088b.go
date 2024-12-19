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
	var a []int
	var n int

    fmt.Scanf("%d", &n)
    a  = createArray(n)

    for i := 0; i<len(a);i++{
        for j := 0; j < len(a)-i-1;j++{
            if a[j] < a[j+1]{
                temp := a[j]
                a[j] = a[j+1]
                a[j+1] = temp
            }
        }
    }
    var sumAlice int = 0
    var sumBob int = 0

    for i := 0; i < len(a);i++{
        if i % 2 == 0{
            sumAlice += a[i]
        }else{
            sumBob += a[i]
        }
    }

    fmt.Println(sumAlice-sumBob)

}