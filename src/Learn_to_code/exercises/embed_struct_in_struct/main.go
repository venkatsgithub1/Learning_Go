package main

import (
	"fmt"

	"github.com/davecgh/go-spew/spew"
)

type vehicle struct {
	doors int
	color string
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
	truckVar := truck{
		vehicle: vehicle{
			doors: 2,
			color: "blue",
		},
		fourWheel: true,
	}

	sedanVar := sedan{
		vehicle: vehicle{
			doors: 4,
			color: "white",
		},
		luxury: true,
	}

	spew.Dump(truckVar)
	spew.Dump(sedanVar)

	fmt.Printf("%v\n", truckVar)
	fmt.Println(sedanVar)
	fmt.Println(truckVar.doors)
	fmt.Println(sedanVar.doors)
}
