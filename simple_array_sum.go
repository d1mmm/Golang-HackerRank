package main
//import "fmt"

func simpleArraySum(ar []int32) int32 {
	var sum int32 = 0
	for i := 0; i < len(ar); i++ {
		sum += ar[i]
	}
	return sum
}

// func main() {
// 	var a, res int32
// 	var ar [1000]int32
// 	fmt.Println("Enter size arr")
// 	fmt.Scanf("%v", &a)
// 	fmt.Println("Enter number arr")
// 	for i := 0; i < int(a); i++ {
// 		fmt.Scanf("%v", &ar[i])
// 	}

// 	for i := 0; i < int(a); i++ {
// 		fmt.Println(ar[i])
// 	}
// 	res = simpleArraySum(ar[:])
// 	fmt.Println("Res = ", res)
// }
