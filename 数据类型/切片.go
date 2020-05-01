
package main 

import (
	"fmt"
	// "math/rand"
	// "time"
)


func main(){
	/*
	切片： 变长数组
	切片是引用类型
	*/
	var sclice1 []int
	sclice1 = []int{1,2,3,4}
	s2 := []int{5,6,7,8}
	fmt.Println(sclice1,s2)

	/*
	使用make创建切片
	make([]type,len,cap)
	*/
	s3 := make([]int,3,8)
	fmt.Printf("%v,%d,%d\n",s3,len(s3),cap(s3))

	/*
	在数组或切片上创建切片
	*/
	b := [10]int{1,2,3,4,5,6,7,8,9,10}
	fmt.Println(b)
	s4 := b[2:6] 
	// 切片容量从创建切片的启始下标开始到数组末尾
	fmt.Printf("%v,%d,%d\n",s4,len(s4),cap(s4))
	s5 := b[6:]
	fmt.Printf("%v,%d,%d\n",s5,len(s5),cap(s5))
	s6 := b[:6]  //从头开始
	fmt.Printf("%v,%d,%d\n",s6,len(s6),cap(s6))
	/*
	引用类型，存储的是数据的内存地址
	任何一边数据发生变化，数据都会发生变化
	*/
	s7 := []int{1,2,3,4,5}
	s8 := s7
	s7[1] = 7
	fmt.Println(s7,s8)
	fmt.Printf("%p,%p\n",s7,s8)
	s8[4] = 8
	fmt.Println(s7,s8)

	/*
	向切片追加数据
	容量不够时会扩容，扩容后切片容量扩大了2倍
	扩容后会返回一个新的切片
	*/
	sclice2 := []int{}
	fmt.Println(sclice2)
	sclice2 = append(sclice2, 1)
	sclice2 = append(sclice2, 2)
	fmt.Println(sclice2)
	sclice2 = append(sclice2, 3)
	fmt.Printf("%v,%d,%d\n",sclice2,len(sclice2),cap(sclice2))

	/*
	修改切片内容也会修改切片所引用的数组的数据
	扩容也会修改引用数组的数据，如果扩容超过源数组的长度，会重新建立一个切片
	*/
	arr1 := [10]int{1,2,3,4,5,6,7,8,9,10}
	fmt.Printf("%p,%v\n",&arr1,arr1)
	sclice3 := arr1[2:8]
	fmt.Printf("%p,%d,%d,%v\n",sclice3,len(sclice3),cap(sclice3),sclice3)
	sclice3[3] = 11
	fmt.Println(arr1)

	/*
	将一个切片加到另外一个切片中
	*/
	ss1 := []int{1,2,3,4}
	ss2 := []int{5,6,7,8}
	ss1 = append(ss1,ss2...)
	fmt.Println(ss1)

	/*
	删除切片
	*/
	del := 3
	ss1 = append(ss1[:del], ss1[del+1:]...)
	fmt.Println(ss1)
	fmt.Println("--------")

	/*
	copy(n,m)  将m中的数据拷贝到n中
	深拷贝： 拷贝数值 （数值型数据）
	浅拷贝: 拷贝地址（引用型数据）
	*/
	
	n := []int{1,2,3,4,5}
	m := []int{7,8,9}
	// copy(n,m)
	copy(m[1:],n[3:])
	fmt.Println(n,m)

}
