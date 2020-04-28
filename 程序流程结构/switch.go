package main

import "fmt"


func main() {
	/*
	case 是无序的
	default 不是必须的，不满足条件就会执行default
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

}
         

