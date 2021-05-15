package main

import "fmt"

func main() {
	preferences := map[string][]string{
		`bond_james`:      []string{`Shaken, not stirred`, `Martinis`, `Women`},
		`moneypenny_miss`: []string{`James Bond`, `Literature`, `Computer Science`},
		`no_dr`:           []string{`Being evil`, `Ice cream`, `Sunsets`},
	}

	preferences["bourne_jason"] = []string{`amnesia`, `guns`, `india`}

	for k, v := range preferences {
		fmt.Println("key: ", k, ", value: ", v)
	}
}
