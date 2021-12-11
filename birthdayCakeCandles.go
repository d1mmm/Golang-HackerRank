package main

//import "fmt"

func birthdayCakeCandles(candles []int32) int32 {
	var res int32
	max := candles[0]
	for i := 0; i < len(candles); i++ {
		if max < candles[i]{
			max = candles[i]
		}
	}
	for i := 0; i < len(candles); i++ {
		if max == candles[i]{
			res += 1
		}
	}
	return res
}

// func main() {
// 	//candles := []int32{302310, 123123,56734,435345}
// 	candles := []int32{18, 90, 90, 13, 90, 75, 90, 8, 90, 43}
// 	fmt.Println(birthdayCakeCandles(candles))
// }
