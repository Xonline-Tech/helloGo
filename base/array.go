package main

import "fmt"

func main() {

	// 声明数组 var arrayName [size]dataType
	var array [10]int
	fmt.Printf("array = %d\n", array) //balance = [0 0 0 0 0 0 0 0 0 0]

	// 初始化数组 var arrayName [size]dataType{elements,...}
	var numbers = [5]int{1, 2, 3, 4, 5}
	numbers1 := [5]int{1, 2, 3, 4, 5}
	fmt.Printf("numbers = %d\n", numbers)   //numbers = [1 2 3 4 5]
	fmt.Printf("numbers1 = %d\n", numbers1) //numbers1 = [1 2 3 4 5]

	// 不定长度数组 var arrayName [...]dataType{elements,...}
	strings := [...]string{"a", "b", "c", "d"}
	fmt.Printf("strings = %s\n", strings) //strings = [a b c d]
	// strings[4] = "e" 不定长度数组根据初始化元素个数创建，在初始化后无法改变长度

	// 初始化固定下下标元素 var arrayName [size]dataType{position:elements,...}
	float32s := [5]float32{0: 1.111, 4: 4.444}
	fmt.Printf("float32s = %f\n", float32s) // float32s = [1.111000 0.000000 0.000000 0.000000 4.444000]

	// 修改固定下标的元素
	strings1 := [...]string{"a", "b", "c", "d"}
	fmt.Printf("strings1修改前 = %s\n", strings1) //strings1修改前 = [a b c d]
	strings1[1] = "b1"
	fmt.Printf("strings1修改后 = %s\n", strings1) //strings1修改后 = [a b1 c d]

}
