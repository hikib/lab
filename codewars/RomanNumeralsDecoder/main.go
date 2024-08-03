package main

import "fmt"

// Create a function that takes a Roman numeral as its argument and
// returns its value as a numeric decimal integer. You don't need to
// validate the form of the Roman numeral.

// Modern Roman numerals are written by expressing each decimal digit of
// the number to be encoded separately, starting with the leftmost digit
// and skipping any 0s. So 1990 is rendered "MCMXC" (1000 = M, 900 = CM,
// 90 = XC) and 2008 is rendered "MMVIII" (2000 = MM, 8 = VIII). The
// Roman numeral for 1666, "MDCLXVI", uses each letter in descending
// order.

// "MM"      -> 2000
// "MDCLXVI" -> 1666
// "M"       -> 1000
// "CD"      ->  400
// "XC"      ->   90
// "XL"      ->   40
// "I"       ->    1

//		Thousands	Hundreds	Tens	Units
// 1	M			C			X		I
// 2	MM			CC			XX		II
// 3	MMM			CCC			XXX		III
// 4				CD			XL		IV
// 5				D			L		V
// 6				DC			LX		VI
// 7				DCC			LXX		VII
// 8				DCCC		LXXX	VIII
// 9				CM			XC		IX

var m = map[string]int{
	"I":  1,
	"IV": 4,
	"V":  5,
	"IX": 9,
	"X":  10,
	"L":  50,
	"C":  100,
	"D":  500,
	"M":  1000,
}

func main() {
	fmt.Println(Decode("MDCLXVI"))
	fmt.Println("----")
	fmt.Println(Decode("XI"))
	fmt.Println("----")
	fmt.Println(Decode("IX"))
	fmt.Println("----")
	fmt.Println(Decode("MIV"))
	fmt.Println("----")
	fmt.Println(Decode("VII"))
}

func Decode(roman string) int {
	num := 0
	end := len(roman)
	next := ""
	for i, c := range roman {
		s := string(c)
		if s == "I" && i == end-2 {
			next = string(roman[i+1])
			if next != "I" {
				s += next
				num += m[s]
				break
			}
		}
		num += m[s]
	}
	return num
}
