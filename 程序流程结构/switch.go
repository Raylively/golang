package main

import "fmt"


func main() {
	/*
	default 不是必须的，不满足条件就会执行default
	switch 后如果没有表达式，相当于布尔类型ture
	case 是无序的
	case 可以有多个值/条件
	switch 可以设置初始化语句
	*/
	a := 7
	switch a {
	default:
		fmt.Println("我也不知道")
	case 1:
		fmt.Println("第一季度")
	case 2:
		fmt.Println("第二季度")
	case 3:
		fmt.Println("第三季度")
	case 4:
		fmt.Println("第四季度")
	}
	
	var num int = 2
	switch num {
	case 2,4,6,8,10:
		fmt.Printf("%d是偶数\n",num)
	}
	
    switch lan := "python";lan{
	case "java":
		fmt.Println("java语言")
	case "c++":
		fmt.Println("c++语言")
	case "python":
		fmt.Println("python语言")
	case "golang":
		fmt.Println("go语言")
	}

	switch score := 45/10;score {
	case 9:
		fmt.Println("优秀")
	case 8:
		fmt.Println("良好")
	case 7:
		fmt.Println("中等")
	case 6:
		fmt.Println("及格")
	case 0,1,2,3,4,5:
		fmt.Println("不及格")
	}

	/*
	fallthrough  穿透switch
	某个case执行后如果有fallthrough,下一个case不判断条件，直接执行
	break 打破switch,用于强制结束某个case的执行
	*/
	day := 0
	switch 2 {
	default:
		fmt.Println("我也不知道")
	case 1,3,5,7,8,10,12:
		day = 31 
	case 4,6,9,11:
		day = 30
	case 2:
		day = 28
		break 
		fmt.Println(2,"月有%d天",day)
	}
	fmt.Println(day)

}
         


