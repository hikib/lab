package main

import (
    "os"
    "fmt"
    "codewars/kata"
    "strconv"
)

func main() {
    input := os.Args[1]

    // num,_ := strconv.ParseUint(input, 10, 64)
    num,_ := strconv.Atoi(input)
    
    // var nums []int
    // for _,v := range input {
    //     num,_ := strconv.Atoi(v)
    //     nums = append(nums, num)
    // }

    fmt.Println(kata.BalancedParens(num))

    // var r []rune
    // for _,s := range input {
    //     r = append(r,s)
    // }
    // fmt.Println(kata.IsValidWalk(r))

    // fmt.Println(kata.RemovNb(num))
}
