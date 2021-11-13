package main

import "fmt"

func main() {
	claster1 :=  [[3,0], [4,3], [6,4], [7,1]]
	claster2 := [[3,0], [4,3], [6,4], [7,1]]
	for i := range claster1{
		v := claster1[i] + claster2[i]
		fmt.Println(v)
	}
}
