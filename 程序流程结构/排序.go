package main 

import (
    "fmt"
)


func main(){
	/*
	冒泡排序
	*/

	arr1 := [10]int{22,67,32,45,98,75,19,54,39,11}
	for j:= cap(arr1)-1; j>0;j--{
		for i:=0; i<cap(arr1)-1; i++{
			if arr1[i] > arr1[i+1]{
				arr1[i],arr1[i+1] = arr1[i+1],arr1[i]
			}
		}
	}
	fmt.Println(arr1)

}
