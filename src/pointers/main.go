package main

import (
	"fmt"
)


var (
	map_f1_2k17 map[int]string
	grid_yr int
)

func main() {
	grid_yr=2017
	map_f1_2k17=make(map[int]string)
	map_f1_2k17[5]="Seb"
	map_f1_2k17[44]="Lewis"
	fmt.Println(map_f1_2k17[0])
	fmt.Println(&grid_yr)
	pointer1:=&grid_yr
	fmt.Println(*pointer1)
}
