package main

// import (
// 	"fmt"
// 	"os"
// 	"strconv"
// )

// func timeConversion(s string) string {
// 	var result string
// 	for i := 0; i < 2; i++ {
// 		result += string(s[i])
// 	}

// 	hours, err := strconv.Atoi(result)
// 	if err != nil {
// 		fmt.Println(err)
// 		os.Exit(2)
// 	}

// 	for i := 0; i < len(s); i++ {
// 		if s[i] == 'A' {
// 			if hours >= 12 {
// 				hours = hours + 12 - 24
// 			}
// 		}
// 		if s[i] == 'P' {
// 			if hours < 12 {
// 				hours += 12
// 			}
// 		}
// 	}

// 	//fmt.Println(hours)
// 	if hours < 9 {
// 		result = "0" + strconv.Itoa(hours)
// 	}

// 	for i := 2; i < len(s)-2; i++ {
// 		result += string(s[i])
// 	}
// 	return result
// }
// func main() {
// 	// s := "12:01:00PM"
// 	// fmt.Println(timeConversion(s))
// 	// st := "07:05:45PM"
// 	// fmt.Println(timeConversion(st))
// 	str := "12:40:22AM"
// 	fmt.Println(timeConversion(str))
// }
