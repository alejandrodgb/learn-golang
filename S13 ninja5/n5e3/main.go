package main

import "fmt"

type vehicle struct {
	doors  int
	colour string
}

type truck struct {
	vehicle
	fourWheel bool
}

type sedan struct {
	vehicle
	luxury bool
}

func main() {
	hmmvw := truck{
		vehicle: vehicle{
			doors:  4,
			colour: "army green",
		},
		fourWheel: true,
	}

	passat := sedan{
		vehicle: vehicle{
			doors:  5,
			colour: "black",
		},
		luxury: true,
	}

	fmt.Println(hmmvw, hmmvw.fourWheel)
	fmt.Println(passat, passat.doors)

}
