
package main 

import (
	"fmt"
	"sort"
	// "math/rand"
	// "time"
)


func main(){
	/*
	map(映射)  用于存储键值对
			   无序
			   key唯一
			   引用类型
	*/

	/*
	定义map,没有赋初值的map是nil，无法数据
	var mapName map[keyType]valueType  // 仅声明
	mapName := map[keyType]valueType{k:v,k:v,...}
	mapName := make(map[keyType]valueType)
	*/

	var map1 map[int]string
	if map1 == nil {
		fmt.Printf("%v是nil的，需要先创建。。。\n",map1)
		map1 = make(map[int]string)
		map1[1] = "jack"
		map1[2] = "lili"
		fmt.Println(map1)
		v,ok := map1[3]
		if ok {
			fmt.Println("获取到的数据是:",v)
		} else {
			fmt.Println("key不存在，无法获取数据")
		}
	}
	fmt.Printf("%v,%T\n",map1,map1)
	map2 := map[string]int{"Go":88,"c":78,"python":98,"java":83}
	map3 := make(map[int]string)
	fmt.Println(map2)
	fmt.Println(map3)
	/*
	获取数据
	value, ok := mapName[key]
	*/
	v,ok := map2["python"]
	fmt.Println(v,ok)
	v1 := map2["Go"]
	fmt.Println(v1)
	/*
	删除
	*/
	delete(map2,"java")
	fmt.Println(map2)

	/*
	修改
	*/
	map2["c"] = 87
	fmt.Println(map2)

	/*
	长度
	*/
	fmt.Println(len(map2))

	/*
	遍历
	*/

	userInfo := make(map[string]string)
	userInfo["name"] = "go"
	userInfo["age"] = "30"
	userInfo["sex"] = "man"
	userInfo["address"] = "beijin"
	fmt.Println(userInfo)
	userInfo["sex"] = "women"
	fmt.Println(userInfo)

	for k,v := range userInfo{
		fmt.Println(k,v)
	}

	game := map[int]string{1:"王者荣耀",4:"绝地求生",3:"LOL",2:"DNF"}
	s1 := make([]int,0,len(game))
	for k := range game {
		s1 = append(s1, k)
	}
	fmt.Println(s1)
	// 使用sort包下的方法Ints()给int型的切片排序
	sort.Ints(s1)
	fmt.Println(s1)
	for _,v := range s1{
		fmt.Println(v,game[v])
	}

	/*
	将map存入slice中
	*/
	s2 := make([]map[string]string,0,10)
	s2 = append(s2,userInfo)
	fmt.Println(s2)

	for i:=0; i<len(s2); i++ {
		for k,v := range s2[i]{
			fmt.Println(k,v)
		}
	}
}
