package main

import "fmt"

func main() {
	// Maps are created with syntax "map" "[" KeyType "]" ElementType
	age := map[string]int{
		"James":  44,
		"Sally":  54,
		"George": 35,
	}

	// Accessing map's items
	fmt.Println("\n Accessing map items")
	fmt.Println("- James age: ", age["James"])

	// Checking whether an entry exists
	fmt.Println("\n Checking whether an entry exists")
	if v, ok := age["Ales"]; ok {
		fmt.Println("Entry value: ", v)
	} else {
		fmt.Printf("Entry does not exist. Value: %v, Check: %v\n", v, ok)

	}

	// Adding an element to a map
	fmt.Println("\nAdding an element to a map")
	age["Maria"] = 22
	fmt.Println(age)

	// Ranging over a map
	fmt.Println("\nRanging over the map")
	for k, v := range age {
		fmt.Println(k, v)
	}

	// Deleting from a map
	fmt.Println("\nDeleting from a map")
	fmt.Println("- Original Map: ", age)
	delete(age, "James")
	fmt.Println("- Deleted entry: James")
	fmt.Println("- Final map: ", age)
}
