package main

import (
	"fmt"
	"sort"
)

func main() {
	// Composite Literal
	x := []int{1, 2, 3, 4, 5} // A SLICE allows for the grouping of VALUES of the same TYPE
	fmt.Println("Original Slice", x)

	//sliceForRange(x)
	//sliceASlice(x, 2, len(x))
	//appendToSlice(x, []int{500, 501, 502})
	//deleteFromSlice(x, 0, 4)
	//makeSlice(10, 12)
	fmt.Println(multiDimensionalSlice([]string{"James", "Bond"}, []string{"John", "Dow"}, []string{"Jaime", "Barrios"}))

}

func sliceForRange(slice []int) {
	fmt.Println("\nSlice For Range")
	for i, v := range slice {
		fmt.Println(i, v)
	}
}

func sliceASlice(slice []int, start int, end int) []int {
	fmt.Println("\nSlice a Slice")
	return slice[1:3]
}

func appendToSlice(slice []int, toAppend []int) {
	fmt.Println("\nAppend to Slice")
	// The formula of append uses a variadic
	// parameter ...x where the elipsis means
	// any number of parameters
	slice = append(slice, 77, 66, 55, 44)
	fmt.Println(slice)

	// When the elipsis on the other side of the variable
	// x... it means that a data structure can be passed.
	slice = append(slice, toAppend...)
	fmt.Println(slice)
}

func deleteFromSlice(slice []int, toDelete ...int) {
	// Pass all the indexes you want to delete
	fmt.Println("\nDelete from Slice")

	fmt.Println("- Initial slice: ", slice)
	fmt.Println("- Indexes to delete", toDelete)

	// Check that all indexes passed are valid
	for _, v := range toDelete {
		if v > len(slice)-1 || v < 0 {
			fmt.Println("Invalid index: ", v)
			return
		}
	}

	sort.Slice(toDelete, func(i, j int) bool {
		return toDelete[i] > toDelete[j]
	})

	for _, v := range toDelete {
		if v == 0 {
			slice = slice[1:]
		} else if v == len(slice) {
			slice = slice[:v]
		} else {
			slice = append(slice[0:v], slice[v+1:]...)
		}
	}
	fmt.Println(slice)
}

func makeSlice(length int, capacity int) {
	fmt.Println("\nMake a Slice")
	// Make a slice with a defined len and capacity
	// to reduce processing time.
	slice := make([]int, length, capacity)
	fmt.Println("- Slice: ", slice)
	fmt.Printf("- Length: %d, Capacity: %d\n", len(slice), cap(slice))

	// Append three more lines to fill the capacity
	slice = append(slice, 444)
	fmt.Println("- Slice appended: ", slice)
	fmt.Printf("- Length: %d, Capacity: %d\n", len(slice), cap(slice))
	slice = append(slice, 555)
	fmt.Println("- Slice appended: ", slice)
	fmt.Printf("- Length: %d, Capacity: %d\n", len(slice), cap(slice))
	slice = append(slice, 666)
	fmt.Println("- Slice appended: ", slice)
	fmt.Printf("- Length: %d, Capacity: %d\n", len(slice), cap(slice))
}

func multiDimensionalSlice(slices ...[]string) [][]string {
	mds := [][]string{}
	for _, v := range slices {
		mds = append(mds, v)
	}
	return mds
}
