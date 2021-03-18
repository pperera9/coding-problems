package main

import "fmt"

func main() {
	findNumbersAddingUp(10, []int{9, 15, 3, 7})
	findNumbersAddingUp(15, []int{9, 15, 3, 6})
	findNumbersAddingUp(12, []int{9, 7, 5, 6})
	findNumbersAddingUp(99, []int{9, 7, 5, 6})

}

func findNumbersAddingUp(numberToFind int, numbers []int) {
	lastIndex := len(numbers) - 1

	for i, number := range numbers  {

		if i == lastIndex {
			fmt.Println("numbers not found in array:", numbers, "adding up to", numberToFind)
			return
		}


		for x := (i + 1); x < len(numbers); x++ {
				if numbers[x] + number == numberToFind {

					fmt.Println("found numbers at index", i, "and", x, "in array:", numbers, "adding up to", numberToFind)
					return
				}
		}
	}

}
