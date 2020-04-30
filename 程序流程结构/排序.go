func main(){
	/*
	冒泡排序
	*/

	arr1 := [10]int{22,67,32,45,98,75,19,54,39,11}
	for i:= 1; i<len(arr1);i++{
		for j:=0; j<cap(arr1)-i; j++{
			if arr1[j] > arr1[j+1]{
				arr1[j],arr1[j+1] = arr1[j+1],arr1[j]
			}
		}
	}
	fmt.Println(arr1)

}
