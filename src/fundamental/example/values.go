package main

import "fmt"
import "math"

func main() {
	const inflationRate = 2.5
	var investmentAmount float64
	expectedReturnRate := 5.5
	var years = 10.0
	fmt.Scan(&investmentAmount)
	futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)
	fmt.Println(futureValue)
	fmt.Println(futureRealValue)

}
