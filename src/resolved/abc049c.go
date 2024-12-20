package main

import (
	"fmt"
	"strings"
)


func main(){
	var s string
	t := []string{"dream", "dreamer", "erase", "eraser"}
	
	
	fmt.Scanf("%s",&s)

	for {
		if strings.HasSuffix(s,t[0]){
			s = strings.TrimSuffix(s,t[0])
		}else if strings.HasSuffix(s,t[1]){
			s = strings.TrimSuffix(s,t[1])
		}else if strings.HasSuffix(s,t[2]){
			s = strings.TrimSuffix(s,t[2])
		}else if strings.HasSuffix(s,t[3]){
			s = strings.TrimSuffix(s,t[3])
		}else if len(s) == 0{
			fmt.Println("YES")
			break
		}else{
			fmt.Println("NO")
			break
		}

	}
}