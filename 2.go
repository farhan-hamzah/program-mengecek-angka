package main

import (
	"fmt"
)

func bilangan(){
	var x, y, z int
	var b bool
	b = false
	fmt.Scan(&x, &y)
	for x > 0{
		z = x%10
		if z == y {
			b = true
		}
		x = x/10
	}
	if b == true{
		fmt.Print("ADA")
	}else{
		fmt.Print("TIDAK")
	}
	
}
func main(){
	bilangan()
}