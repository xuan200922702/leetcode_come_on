package main

import "fmt"

func main(){
	I: for i :=0;i<2;i++{
		for j:=0;j<5;j++{
			if  j==2{
				continue  I
			}
			fmt.Println("hello")
		}
	}
	fmt.Println("hi")
}