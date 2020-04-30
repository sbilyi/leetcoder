package main

import (
	"bufio"
	"fmt"
	"os"
)

/**
|1 2 2 3 5|
|1 1 2 2 5|
|1 2 2 3 5|
|1 2 2 3 5| => 5

*/

func TestCountCountries() int {
	region := [][]int{{1, 2, 2, 3, 5}, {1, 1, 2, 2, 5}, {1, 2, 2, 3, 5}, {1, 2, 2, 3, 5}}
	return countCountries(region)
}

func countCountries(region [][]int) int {
	fmt.Printf("============================== Step %2.d. ==============================\n", 0)
	print(region)
	countriesNumber := 0
	treeNumber := 1
	for i := 0; i < len(region); i++ {
		for j := 0; j < len(region[i]); j++ {
			if region[i][j] > 0 {
				visit(region, i, j, region[i][j])
				fmt.Printf("============================== Step %2.d. ==============================\n", treeNumber)
				print(region)
				bufio.NewReader(os.Stdin).ReadBytes('\n')
				countriesNumber++
			}
		}
	}

	return countriesNumber
}

func visit(region [][]int, i int, j int, countryColor int) {
	if i < 0 || j < 0 || i >= len(region) || j >= len(region[i]) {
		return
	}

	color := &region[i][j]

	if (*color != -2 && *color != -1) && *color == countryColor {
		*color = -2 // grey
		visit(region, i-1, j, countryColor)
		visit(region, i, j+1, countryColor)
		visit(region, i, j-1, countryColor)
		visit(region, i+1, j, countryColor)
		*color = -1 // black
	}

}

func print(region [][]int) {
	for i := 0; i < len(region); i++ {
		fmt.Printf("%d\n", region[i])
	}
}
