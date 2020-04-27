package main

import "fmt"


func main(){
	/*
	if语句
	*/

	a := -1
	if a>0 { 
		fmt.Printf("%d是正数",a)
	}
	if a<0 {
		fmt.Printf("%d是负数",a)
	}
	fmt.Println("========")

	if a>0 {
		fmt.Printf("%d是正数",a)
	} else {
		fmt.Printf("%d是负数",a)
	}

	sex := "男"
	if sex == "男" {
		fmt.Println("你是男生去男厕所")
	} else {
		if sex == "女" {
			fmt.Println("你是女生去女厕所")
		}else {
			fmt.Println("不知道你该去哪里")
		}
	}

	if sex == "1" {
		fmt.Println("你是男生去男厕所")
	} else if sex == "女" {
			fmt.Println("你是女生去女厕所")
	} else {
			fmt.Println("不知道你该去哪里")
	}

	score := 67
	if score < 60 {
		fmt.Println("优秀")
	} else if score < 70 {
		fmt.Println("良好")
	} else if score < 80 {
		fmt.Println("中等")
	} else if score < 90 {
		fmt.Println("及格")
	} else if score <= 100 {
		fmt.Println("不及格")
	}
	
	/*
	if初始化语句
	*/

	if b :=6; b>0 {
		fmt.Println("正数")
	}

    
}
