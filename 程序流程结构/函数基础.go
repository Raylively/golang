

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

func getMul(nums ... int)(sum int){
	sum = 1
	for i:=1; i<=len(nums); i++ {
		sum *= i
	}
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

	/*
	可变参数
	func funcName(variablesName ... type){}
	一个函数只能有一个可变参数
	如果函数中还有其他参数，可变参数必须放在最后
	*/
	res := getMul(1,2,3,4,5)
	fmt.Println(res)

	s1 := []int{1,2,3,4,5,6,7,8,9}
	res1 := getMul(s1...)
	fmt.Println(res1)

}
