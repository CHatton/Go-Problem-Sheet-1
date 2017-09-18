/*
Problem description

Write a function that calculates the sum of the digits of the factorial of a number. 
n! means n x (n − 1) ... x 3 x 2 x 1. For example, 10! = 10 x 9 x ... x 3 x 2 x 1 = 3628800, 
and the sum of the digits in the number 10! is 3 + 6 + 2 + 8 + 8 + 0 + 0 = 27. 
Find the sum of the digits in the number 100!
*/

package main

import (
    "fmt"
    "math/big"
    "strings"
    "strconv"
)


func sumDigitsInString(resultString string) int {
    allDigits := strings.Split(resultString, "") // get a list of every digit

    totalSum := 0
    for _, digit := range allDigits {
        val, _ := strconv.Atoi(digit) // parse digit to value "9" -> 9
        totalSum += val
    }
    return totalSum
}

func main() {
    // we need to use bit.Int because we can't hold the value with standard types
    bigIntptr := new(big.Int)
    // the actual factorial value
    factorialResult := bigIntptr.MulRange(1,100) // equivalent to 100 x 99 x 98 x ... x 1
    totalSum := sumDigitsInString(factorialResult.String())
    fmt.Println("The sum of the digits of 100! is:", totalSum)
}