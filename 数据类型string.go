
package main 

import (
	"fmt"
	"strings"
	// "sort"
	// "math/rand"
	// "time"
)


func main(){
	/*
	数据类型
		基本类型 数值，布尔，字符串
		符合类型 数组，切片，映射，函数，结构提，接口。。。
	值类型，引用类型 
	*/
	/*
	字符串 字节的集合（字符序列）    //字节（byte）  : 8个01码
	字符串内容不可改变
	*/
	num1 := 65  //整型变量 数值65
	num2 := 'A' //整型变量 赋值为A的编码，数值也是65
	num3 := "A" //字符变量 A
	fmt.Printf("%v,%T\n",num1,num1)
	fmt.Printf("%v,%T\n",num2,num2)
	fmt.Printf("%v,%T\n",num3,num3)

	s1 := "hello World!"
	s2 := "hello 世界！"
	fmt.Println(s1,len(s1))
	fmt.Println(s2,len(s2))
	// 取值
	fmt.Println(s1[1],s2[7])  // 返回字符对应的编码值
	fmt.Printf("%q,%c,%d\n",s1[7],s1[7],s1[7]) 
	// 遍历
	for i:=0; i<len(s1); i++ {
		fmt.Printf("%c\n", i)
	}
	for _,c := range s1 {
		fmt.Printf("%c\n", c)
	}
	
	/*
	将字符串一字节的形式存储在切片中
	*/
	s3 := "abcdefg"
	arr1 := []byte(s3)
	fmt.Printf("%c,%T\n", arr1,arr1)

	/*
	将切片存入字符串中
	*/
	arr2 := []uint8{65,66,67,68}
	s4 := string(arr2)
	fmt.Printf("%v,%T",s4,s4)

	arr3 := []int32{97,98,99,100}
	s5 := string(arr3)
	fmt.Println(s5)
	fmt.Printf("%v,%T\n", s5,s5)

	/*
	strings.Contains()  判断字符串是否包涵指定的内容
	strings.Contains() 判断字符串中是否包涵指定指定内容中任意一个
	*/
	ss1 := "hello world"
	fmt.Println(strings.Contains(ss1,"hello")) 
	fmt.Println(strings.ContainsAny(ss1,"abcd")) 
	fmt.Println(strings.Index(ss1,"llo"))  // 返回子串第一次出现的下标索引位置，不存在返回-1
	fmt.Println(strings.LastIndex(ss1,"l")) //返回子串最后一次出现的位置

	ss2 := "2020-04-30笔记.txt"
	fmt.Println(strings.HasPrefix(ss2,"2020")) // 是否以指定内容开始
	fmt.Println(strings.HasSuffix(ss2,"txt")) // 是否以指定内容结尾
	
	// 统计
	fmt.Println(strings.Count(ss2,"txt"))  //统计子串出现的次数
	// 切割,根据指定切割符切割，切割完的数据存储在切片中
	fmt.Println(strings.Split(ss2,"-"))
	fmt.Println(strings.Split(ss2,"."))
	fmt.Printf("%v,%T\n",strings.Split(ss2,""),strings.Split(ss2,""))
	fmt.Println(strings.SplitN(ss2, "",5)) //指定切割成几部分，多余的部分会拼接在一起

	// 大小写转换
	ss3 := "hello123WorLD"
	fmt.Println(strings.ToUpper(ss3))
	fmt.Println(strings.ToLower(ss3))

	//  替换
	fmt.Println(strings.Replace(ss3, "1", "*", -1)) // 将1替换成* 

	// 去除指定内容  首尾
	ss4 := "    zhang san   "
	fmt.Println(ss4)
	fmt.Println(strings.Trim(ss4, " +")) // 去除首位的空格和+
	fmt.Println(strings.TrimSpace(ss4))   // 去除首位空格

	// 重复
	fmt.Println(strings.Repeat("hello",5))

	// 截取
	ss5 := "hellowolrd"
	ss6 := ss5[:5]
	fmt.Printf("%v,%T\n", ss6,ss6)

}
