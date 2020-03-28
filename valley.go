package main

import "fmt"

func countingValleys(n int, s string) int {
	valley := 0
	height := 0
	goneDown := false
	for _, letter := range s {
		if height < 0 {
			goneDown = true
		}
		if string(letter) == "U" {
			height++
		} else if string(letter) == "D" {
			height--
		}
		if height == 0 && goneDown == true {
			valley++
			goneDown = false
		}
		fmt.Println(height)
	}
	return valley
}

func main() {
	/* s = UDDDUDUU
	   _/\      _
	      \    /
	       \/\/
	*/
	countingValleys(8, "UDDDUDUU")

}
