package main


import "fmt"


func main() {

	/*
	基本类型：
		数值类型（整数，浮点，复数)
		布尔类型(bool)
			true
			false
		字符串类型(string)
	*/

	// bool
	b1 := true
	b2 := false
	fmt.Printf("b1的数值是：%t,类型是：%T\n",b1,b1)
	fmt.Printf("b2的数值是：%t,类型是：%T\n",b2,b2)
	
	// 数值型--整数-有符号
	// int8 别名：byte
	// int32 别名 rune
	var d1 int8 = 100
	// int8 0111 1111 -128-(2^7-1)
	// int16 0111 1111 1111 1111 -32768-(2^15-1)
	fmt.Printf("数值d1的值是：%d,类型是：%T",d1,d1)

	// 数值型--证书-无符号
	var d2 uint8 = 200
	// uint8 1111 1111 0-(2^8-1)
	fmt.Printf("数值d2的值是：%d,类型是：%T",d2,d2)

	// int uint 根据操作系统相关（32位操作系统，64位操作系统）
	var d3 int = 300
	fmt.Printf("数值d3的值是：%d,类型是：%T",d3,d3)

	// 小数
	// float32  单精度
	// float64  双精度
	var f1 float32 = 3.14
	var f2 float64 = 6.28
	fmt.Printf("数值f1的值是：%f,类型是：%T",f1,f1)
	fmt.Printf("数值f2的值是：%f,类型是：%T",f2,f2)
	fmt.Println()
	
	// 字符串
	// 使用“” ``定义
	var s1  string = `hello world `
	fmt.Printf("s1的值是：%s,类型是：%T",s1,s1)
}
