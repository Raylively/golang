
package main

import "fmt"

func main(){
	/*
	for 表达式1: 表达式2: 表达式3{
		循环体
	}
	*/

	sum := 0
	for i := 0; i<11; i++ {
		sum += i
	}
	fmt.Println(sum)

	/*

	*/
	for i := 1; i<101; i++ {
		fmt.Print(i,"\t")
		if i % 10 == 0 {
			fmt.Println()
		}
	}
	fmt.Println()

	/**/
	count := 0
	for i := 1; i < 101 ; i++ {
		if i % 3 == 0 && i % 5 !=0 {
			fmt.Print(i,"\t")
			count += 1
			if count % 5 == 0 {
				fmt.Println()
			}
		}
	}
	fmt.Printf("\n一共%d个数\n",count)

	/*
	for 循环后的所有表达式 都可以缺省
	表达式2 如果没有，默认为true， 及无限循环
	*/
	i := 5
	for ; i<=10 ;  {
		println(i)
	}

	/*
	死循环
	for {}
	for true {}
	for ;;{}
	*/

}
