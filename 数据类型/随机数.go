package main 

import (
    "fmt"
	"math/rand"
	"time"
)


func main(){

	num1 := rand.Intn(10)
	fmt.Println(num1)

	/*
	设置种子数
	*/
	rand.Seed(time.Now().Unix())
	num2 := rand.Intn(101)
	fmt.Println(num2)

	/*
	参数游戏
	*/
	rand.Seed(time.Now().Unix())
	num3 := rand.Intn(100) + 1
	fmt.Printf("系统产生的随机数为%d",num3)

	guessNum := 0
	for i:=0;i<3;i++ {
		fmt.Println("请输入猜测的数字：")
		fmt.Scanln(&guessNum)
		if guessNum > num3 {
			fmt.Println("猜大了")
		} else if guessNum < num3 {
			fmt.Println("猜小了")
		} else {
			fmt.Println("猜对了")
			break
		}
	}

}
