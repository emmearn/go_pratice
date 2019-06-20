package main

import "fmt"

func main() {
	mySlice := []string{"Hi", "There", "How", "Are", "You"}

	updateSlice(mySlice)
	fmt.Println(mySlice)
	/*
		Slices are a reference type like
		maps, channels, pointers and functions
	*/
}

func updateSlice(s []string) {
	s[0] = "Bye"
}
