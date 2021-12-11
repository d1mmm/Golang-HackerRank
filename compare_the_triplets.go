package main

// import (
// 	"fmt"
// )

func compareTriplets(a []int32, b []int32) []int32 {
	var res []int32
	var temp [2]int32
	for i := 0; i < len(a); i++ {
		if a[i] > b[i] {
			temp[0] += 1
		} 
		if a[i] < b[i] {
			temp[1] += 1	
		}
	}
	for i:=0; i < len(temp); i++{
		res = append(res, temp[i])
	}
	return res
}

// func main() {
// 	a := []int32{17, 28, 30}
// 	b := []int32{99, 16, 8}
// 	for i := range a {
// 		fmt.Print("\t", a[i])
// 	}
// 	fmt.Println()
// 	for i := range b {
// 		fmt.Print("\t", b[i])
// 	}
// 	fmt.Println()

// 	fmt.Println(compareTriplets(a, b))
// }
