package main 

import "fmt"

func main() {
	/* 
		常量习惯上字母都大写
		标识符如果首字母大写表示共用，小写表示私有
		常量可以直接使用内置函数来赋值
	*/

	// 显示定义
	const pURL string = "http://baidu.com"

	// 隐式定义
	const PI = 3.14
	
	fmt.Println(pURL,PI)

	// 常量组
	const (
		M = 100
		N = "world"
	)
	fmt.Println(M,N)
	// 常量组中如果常量没有赋值，将会使用上一个常量的值
	const (
		A int = 10
		B
		C string = "jack"
		D
	)
	fmt.Println(A,B,C,D)

	// 常量可以做枚举
	const (
		SPRING = "春天"
		SUMMER = "夏天"
		AUTUMN = "秋天"
		WINTER = "冬天"
	)
	fmt.Println(AUTUMN)

	// 特殊常量iota, 默认为0  go的内置结束起
	// 常量组中每定义一个常量都会+1，当遇到下一个常量组定义iota是会清0
	const (
		a = iota 
		b = iota
		c = iota
	)
	const (
		d = iota
		e
		f
	)
	fmt.Println(a,b,c)
	fmt.Println(d,e,f)
	
}
