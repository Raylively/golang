
package main 

import (
	"fmt"
	// "sort"
	// "math/rand"
	// "time"
)

func showMessage(){
	fmt.Println("hello wolrd")
}

// func getAdd(x int,y int) {
func getAdd(x,y int)(int) {
	/*
	求和
	*/
	return x+y
}

func getAreaAndPerimeter(width,height float64)(area,perimeter float64){
	area = width*height
	perimeter = 2*(width+height)
	return 
}

func main(){
	/*
	函数： 一段具有特殊功能的代码，可以多次被调用
	函数的参数: 形参 实参
	*/
	sum := getAdd(10, 20)
	fmt.Println(sum)
	fmt.Println(getAreaAndPerimeter(2,10))
	showMessage()

}
