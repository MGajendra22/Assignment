package main

import (
	"fmt"
	"time"
)

func convert(amount int, source string, target string) float32 {

	var str string = source + target

	fmt.Println(str)

	mp := map[string]float32{
		"USDINR": 85.44,
		"INRUSD": float32(1 / 85.44),
		"USDEUR": float32(1 / 0.93),
		"EURUSD": 0.93,
		"USDJPY": 156.82,
		"JPYUSD": float32(1 / 156.82),
		"INREUR": 0.010,
		"INRJPY": 1.69,
		"EURINR": float32(1 / 0.010),
		"EURJPY": 166.14,
		"JPYINR": float32(1 / 1.69),
		"JPYEUR": float32(1 / 166.14),
	}

	var res float32 = mp[str] * float32(amount)
	return res

}

func list() {

	var ls = []string{"USD", "INR", "EUR", "JPY"}

	for i := 0; i < len(ls); i++ {
		fmt.Println(ls[i])
	}
}

func validate(source, target string) bool {

	var ls = []string{"USD", "INR", "EUR", "JPY"}

	var flag1 bool = false
	var flag2 bool = false

	for i := 0; i < len(ls); i++ {

		if ls[i] == source {
			flag1 = true
		}

		if ls[i] == target {
			flag2 = true
		}
	}

	return (flag1 && flag2)

}

func main() {

	fmt.Println("------------------------------ Welcome to Currency Converter ------------------------------")
	curtime := time.Now()
	hr := curtime.Hour()

	switch {
	case (hr > 6 && hr < 12):
		fmt.Println("Good Morning!")

	case (hr >= 12 && hr <= 5):
		fmt.Println("Good AfterNoon!")

	case (hr > 5 && hr < 20):
		fmt.Println("Good Evening!")

	default:
		fmt.Println("Good Night!")

	}
	fmt.Println("Disclaimer : Make sure you write in capital letters")

	fmt.Println("Show Available Currencies List : (YES/NO)")
	var temp string
	fmt.Scanf("%s", &temp)

	if temp == "YES" {
		list()
	}

	var amount int
	var source string
	var target string
	fmt.Println("Current Currency : ")
	fmt.Scanf("%s", &source)
	fmt.Println("Enter the amount : ")
	fmt.Scanf("%d", &amount)

	fmt.Println("Target Currency : ")
	fmt.Scanf("%s", &target)

	var res float32 = convert(amount, source, target)
	if !validate(source, target) {
		fmt.Println("Invalid Input , Please try something else with given currencies ☺️")
	} else {
		fmt.Println("Converted Currency : ", res)
	}

	fmt.Println("-------------- Thank You ------------------")
}
