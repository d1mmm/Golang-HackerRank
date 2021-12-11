package main

// import (
// 	"fmt"
// )

// func plusMinus(arr []int32) {
// 	var count [3]int32
// 	var result float32
// 	for i:= 0; i < len(arr); i++{
// 		if arr[i] > 0{
// 			count[0] += 1  
// 		}
// 		if arr[i] < 0{
// 			count[1] += 1 
// 		}
// 		if arr[i] == 0{
// 			count[2] += 1 
// 		}
// 	}

// 	for i := 0; i < len(count); i++{
// 		result = float32(count[i]) / float32(len(arr))
// 		fmt.Printf("%f\n", result)
// 	}
// }

// func main(){
// 	arr := []int32{-4,3,-9,0,4,1}
// 	plusMinus(arr)
// }