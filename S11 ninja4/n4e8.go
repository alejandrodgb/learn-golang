package main

import "fmt"

func main() {
	preferences := map[string][]string{
		`bond_james`:      []string{`Shaken, not stirred`, `Martinis`, `Women`},
		`moneypenny_miss`: []string{`James Bond`, `Literature`, `Computer Science`},
		`no_dr`:           []string{`Being evil`, `Ice cream`, `Sunsets`},
	}

	for k, v := range preferences {
		fmt.Println("key: ", k, ", value: ", v)
		for i, vv := range v {
			fmt.Println("\tindex: ", i, ", value: ", vv)
		}
	}
}
