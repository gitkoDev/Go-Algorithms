// 2 Sorting algorithm

/*
Task: given an array of unsorted values, sort them in a descending/ascending order and return the resulting array
*/

package main

import "fmt"

func main() {
	selectionSort([]int{3, 1, 2, -500, 0, 5, 2})
}

func selectionSort(slice []int) {
	smallestIndex := 0

	fmt.Printf("Beginning to sort the array... The array so far is %v\n", slice)
	for i := 0; i < len(slice); i++ {
		// We take slice[i] and compare it with all the other numbers. E.g.: compare 3 and 1, 3 and 2, 3 and -500 and so on. On every iteration we consider slice[i] the smallest number yet, because all the other numbers on the left (i - 1, i - 2...) have already been sorted and we don't need to compare them.

		smallest := slice[i]

		for j := i + 1; j < len(slice); j++ {

			// Inner loop to go through all the other elements after "i" and check if any of them is smaller. (on first iterations it's all the numbers after index 0, then all the numbers after index 1 and so on).

			if slice[j] < smallest {
				// If a "j" number is less than "i" (the number we compare all the others with), we make assign it to "smallest"
				smallest = slice[j]
				smallestIndex = j
			}
		}
		// Create a temporaty variable to store the item that we were comparing everything to. We need to do this to put the smallest number in it's place.
		temp := slice[i]
		// Put the smallest number of this iteration on "i" position. E.g. (on first iteration -500 is the smallest, so slice[i]/slice[0] becomes -500)
		slice[i] = smallest
		// Put the "i" element in place of the smallest number. E.g. we put number 3 in place of -500 (slice[3] = 3)
		slice[smallestIndex] = temp

		/*
			So what happens on first iteration is:

			temporaryVariable <---- 3 (because we need to put -500 in its place)
			slice[0] <---- -500
			slice[3](original position of -500) <---- 3

			-500 and 3 have swapped places
			.......
		*/
		fmt.Printf("Sorting the array... The array so far is %v\n", slice)
	}
	fmt.Printf("The sorted array is %v\n", slice)
}
