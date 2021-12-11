package main

// import (
// 	"fmt"
// )

func miniMaxSum(arr []int32) {
	var minel, maxel int64
	var resmin []int64
	var resmax []int64

	min := int64(arr[0])
	max := int64(arr[0])
	for i := 0; i < len(arr); i++ {
		if min > int64(arr[i]) {
			min = int64(arr[i])
		}
		if max < int64(arr[i]) {
			max = int64(arr[i])
		}
	}

	//fmt.Println(min, max)

	for i := 0; i < len(arr); i++ {
		if int64(arr[i]) < max {
			resmin = append(resmin, int64(arr[i]))
		}
		if int64(arr[i]) > min {
			resmax = append(resmax, int64(arr[i]))
		}
		if min == max {
			maxel = max * 4
			minel = min * 4
		}
	}

	for i := 0; i < len(resmin); i++ {
		minel += resmin[i]
	}

	for i := 0; i < len(resmax); i++ {
		maxel += resmax[i]
	}

	//fmt.Println(minel, maxel)
}

// func main() {
// 	arr := []int32{140638725, 436257910, 953274816, 734065819, 362748590}
// 	arr2 := []int32{5, 5, 5, 5, 5}
// 	miniMaxSum(arr)
// 	miniMaxSum(arr2)
// }
