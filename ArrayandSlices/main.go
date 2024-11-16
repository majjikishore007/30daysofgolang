package main

import "fmt"

// func Sum(numbers [5]int) int{
// 	sum :=0

// 	for i:=0;i<5;i++{
// 		sum+=numbers[i]
// 	}
// 	return sum
// }

//  refator

func Sum(numbers []int) int{
	sum:=0
	for _,num:= range numbers{
		sum+=num
	}
	return sum
}

// func SumAll(listOfList ...[]int) ([]int) {
// 	len := len(listOfList)
// 	// new way to create a slice 
// 	sums:=make([]int,len)
	
// 	for i,numbers:=range listOfList{
// 		sums[i]=Sum(numbers)
// 	}
// 	return sums
// }

// refactored

func SumAll(listOfList ...[]int) ([]int) {
	var sums []int
	for _,numbers:=range listOfList{
		sums=append(sums,Sum(numbers))
	}
	return sums
}


func SumAllTails(lot ...[]int)[]int  {
	var sums []int
	for _,numbers:=range lot{
		if len(numbers)==0{
			sums=append(sums,0)
		}else {
			tail:= numbers[1:]
			sums=append(sums,Sum(tail))
		}
	}
	return sums
}


func main()  {
	arr := [5]int{1,2,3,4,5}
	fmt.Println(len(arr))
	fmt.Println(cap(arr))
	s := arr[:3]
	s[0]=10
	fmt.Println(len(s))
	fmt.Println(cap(s))
	fmt.Println(s)
	fmt.Println(arr)
}

