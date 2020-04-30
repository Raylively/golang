package main 

import (
    "fmt"
)


func main(){
	/*
	数组： 存储一组相同数据类型的数据
	特点： 有序，数据可以重复，数组长度固定
	*/
	var arr1[5] int
	fmt.Println(arr1)
	fmt.Println(len(arr1))
	fmt.Printf("%v,%T\n",arr1,arr1)
	// 赋值
	arr1[0] = 1
	arr1[1] = 3
	arr1[2] = 5
	arr1[3] = 7
	arr1[4] = 9
	var b = [5]int{6,7,8,9,10}
	var c = [5]int{1,2,3}
	var d = [5]int{4:100}
	var e = [5]int{1:2,3:99}
	var f = [...]int{1,2,3,4,5}
	var g = [...]int{1:100,9:88}
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
	fmt.Println(g)
	/*
	访问数组: 存储获取数据
	*/

	fmt.Println(arr1)
	for i:=0;i<len(arr1);i++ {
		fmt.Printf("数组中第%d个数为%d\n",i+1,arr1[i])
	}

	arr2 := [10]int{1,2,3,4,5,6,7,8,9,10}
	sum2 := 0
	for i:=0;i<cap(arr2);i++ {
		sum2 += arr2[i]
	}
	fmt.Println(sum2)

	/*
	使用range遍历
	*/
	for i,v := range arr2  {
		fmt.Println(i,v)
	}
	sum3 := 0
	for _,v := range arr2 {
		sum3+= v
	}
	fmt.Println(sum3)

}
