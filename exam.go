package main

import "fmt"

func main() {

	// name := "Big"
	// age, age2 := 26, 30
	// fmt.Println("My name is ", name)
	// fmt.Println("age => ", age)
	// fmt.Println("age => ", age2)
	// fmt.Printf("data value %v\n", name)
	// fmt.Printf("data type %T\n", age)

	// var username string
	// fmt.Print("กรอกชื่อ")
	// fmt.Scanf("%s", &username)
	// fmt.Println("ชื่อ ===>", username)

	// arrNumber := [3]int{10, 20, 30}
	// fmt.Println("arr == >", arrNumber)

	//defalt {0,0,0}
	// var arrNumber [3]int
	// fmt.Println("arr defalt == >", arrNumber)
	// arrNumber[0] = 100
	// arrNumber[1] = 200
	// arrNumber[2] = 300
	// fmt.Println("arr == >", arrNumber)

	// coin := map[string]string{}
	// coin["ETH"] = "Ether"
	// coin["BTC"] = "Bitcoin"

	// fmt.Println("coin => ", coin)

	// numbers := []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 10}
	// for _, value := range numbers {
	// 	fmt.Println("Value ==> ", value)
	// }

	// language := map[string]string{"TH": "Thai", "EN": "English"}
	// for key, value := range language {
	// 	fmt.Println("Key => ", key, "Value => ", value)
	// }

	result, check := summation(5, 3)
	fmt.Println("ผลรวม = ", result)
	fmt.Println(check)

}

func summation(number1, number2 int) (int, string) {
	total := number1 + number2
	status := ""
	if total%2 == 0 {
		status = "เลขคู่"
	} else {
		status = "เลขคี่"
	}
	return total, status
}
