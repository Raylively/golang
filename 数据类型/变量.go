package main

import "fmt"

func main() {

	// 变量定于的语法格式

	//  常规写法
	var age int
	age = 30
	fmt.Println(age)

	// 简化
	var name string = "paul"
	fmt.Println(name)

	// 类型推断
	var float1 = 13145.20
	fmt.Printf("%f,%T",float1, float1)

	// 简短写法，只能在函数中使用
	reward := 1000
	fmt.Println(reward)


	// 多个变量
	var a,b,c int
	a = 1
	b = 2
	c = 3

	var d,e,f string = "zhang","li","wang"

	var m,n = 1 , "zhao"

	num1,num2 := 100,200  // 左边的多个变量必须有一个是新的

	// 变量的集合
	var (
		h = 8
		j =9
	)

	fmt.Println(a,b,c,d,e,f,num1,num2,m,n,h,j)

   /* 变量的默认值（零值）
	int,float---> 0
	string-----> ""
   */


   var x int
   var y float64
   var z string
   fmt.Println(x,y,z)


   /* 格式化
	%d, %f,%s, %T,\n
   */
   var n1,f1,s1 = 520,1314.520,"i love you"
   fmt.Printf("%d,%T\n",n1,n1)
   fmt.Printf("%.2f,%T\n",f1,f1)
   fmt.Printf("%s,%T\n",s1,s1)

}


