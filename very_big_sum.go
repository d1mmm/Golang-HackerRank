package main
//import "fmt"

func aVeryBigSum(ar []int64) int64 {
	var res int64 = 0
	for i := 0; i < len(ar); i++ {
		res += ar[i]
	}
	return res
}


// func main() {
// 	var a int
// 	var res int64
// 	var ar[10]int64
// 	fmt.Println("Enter size arr")
// 	fmt.Scanf("%v\n", &a)

// 	fmt.Println("Enter number arr")
// 	for i := 0; i < int(a); i++{
// 		fmt.Scanf("%v", &ar[i])
// 	}

// 	fmt.Println("Your arr")
// 	for i := 0; i < int(a); i++{
// 		fmt.Println(ar[i])
// 	}

// 	res = aVeryBigSum(ar[:])
// 	fmt.Println("Res = ", res)
// }