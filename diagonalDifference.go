package main

// import (
// 	"fmt"
// )

// func main() {
// 	arr := [][]int32{{11, 2, 4}, {4, 5, 6}, {10, 8, -12}}
// 	fmt.Println("Your arr")

// 	for i := 0; i < len(arr); i++ {
// 		for j := 0; j < len(arr); j++ {
// 			fmt.Print("\t", arr[i][j])
// 		}
// 		fmt.Println()
// 	}
// 	var maind, secd int32

// 	for i := 0; i < len(arr); i++ {
//         maind += arr[i][i]
// 	}
// 	fmt.Println("Sum of main diagonal = ", maind)

// 	for j := 0; j < len(arr); j++ {
// 		secd += arr[j][len(arr)-1-j]
// 	}
// 	fmt.Println("Sum of add diagonal", secd)

// 	res := maind - secd

// 	if res < 0{
// 		res *= -1
// 	}

// 	fmt.Println("Your diagonal differnce", res)
// }
