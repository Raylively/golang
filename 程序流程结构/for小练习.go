package main 

import "fmt"

func main(){

	/*
	*
	**
	***
	****
	*****	
	*/

	for i:=0;i<5;i++{
		for j:=0;j<=i;j++{
			fmt.Print("*")
		}
		fmt.Println()
	}


	/*
		*
	   **
	  ***
	 ****
	*****
	*/

	for i:=0;i<5;i++{
		for j:=0;j<5;j++{
			if j>3-i{
				fmt.Print("*")
			} else  {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
	/*
		*
	   ***
	  *****	
	 *******
	  *****
	   ***
	    *
	*/
	for i :=1; i<5; i++ {
		for j:=1; j<8; j++ {
			if j<=4-i {
				fmt.Print(" ")
			} else if j>=4+i {
				fmt.Print(" ")
			} else {
				fmt.Print("*")
			}
		}
		fmt.Println()
	}
	for i :=4; i>0; i-- {
		for j:=1; j<8; j++ {
			if j<=4-i {
				fmt.Print(" ")
			} else if j>=4+i {
				fmt.Print(" ")
			} else {
				fmt.Print("*")
			}
		}
		fmt.Println()
	}
	/*
	乘法表
	*/
	for i:=1; i<10; i++ {
		for j:=1; j<=i; j++ {
			fmt.Printf("%2d*%2d=%2d\t",j,i,i*j)
		}
		fmt.Println()
	}

	/*
	百元白鸡
	100元买100只鸡
	公鸡5元一只，母鸡3元1只，小鸡1元3只
	*/
	for i:=0; i<=20; i++{ //公鸡数量
		for j:=0; j<=33; j++ { //母鸡数量
			k:= 100 -i -j //小鸡
			if i*5+j*3+k/3 == 100 && k%3 == 0 {
				fmt.Printf("小鸡%d只，母鸡%d只，公鸡%d只",k,j,i)
				fmt.Println()
			}
		}
	}

}
