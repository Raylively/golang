
package main 

import (
	"fmt"
	// "sort"
	// "math/rand"
)

func getSum(n int)(int){
	if n == 1{
		return 1
	}
	return getSum(n-1) + n
}


func getMul(n int)int{
	// 阶乘
	if n == 1 || n == 0{
		return 1
	}
	return getMul(n-1)*n
}

func getTot(n int)int{
	/*
	裴伯纳切数列
	*/
	if n == 1 || n == 2 {
		return 1
	}
	return getTot(n-1) + getTot(n-2)

}

func func1(a,b int)(int){
	return a+b
}

func main(){
	/*
	递归函数
	*/
	res := getSum(5)
	fmt.Println(res)
	
	res1 := getMul(6)
	fmt.Println(res1)

	res2 := getTot(12)
	fmt.Println(res2)

	/*
	定义函数类型的变量
	*/
	var a func(int,int)(int)
	a = func1
	res3 := a(1,2)
	fmt.Println(res3)

	/*
	匿名函数
	*/
	func (){
		fmt.Println("我是匿名函数")
	}()

	res4 := func (a,b int)(int){
		return a+b
	}(1,2)
	fmt.Println(res4)

	for i:=0; i<3; i++ {
		func(){
			fmt.Println("我是匿名函数:",i)
		}()
	}
	/*
	函数作为参数
	*/

	fmt.Println(oper(20,10,add))
	fmt.Println(oper(20,10,sub))
	fmt.Println(oper(20,10,mul))
	fmt.Println(oper(20,10,div))

	/*
	将函数作为返回值   
	应用： 闭包
	*/
	res5 := increment()
	fmt.Println(res5())
	fmt.Println(res5())


}

func oper(a,b int, fun func(int,int)(int))int{
	res := fun(a,b)
	return res
}

func add(a,b int)(int){
	return a+b
}
func sub(a,b int)(int){
	return a-b
}
func mul(a,b int)(int){
	return a*b
}
func div(a,b int)(int){
	return a/b
}

func increment()(func ()int){
	x := 0
	fun := func()int{
		x++
		return x
	}
	return fun
}
