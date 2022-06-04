package main

import "fmt"

func main() {
	var n [3]int
	fmt.Scan(&n[0], &n[1], &n[2])
	check := 0

	for i := 0; i < n[0]; i++ {
		for h := 0; h < n[1]; h++ {

			for j := 0; j < n[0]; j++ {
				for w := 0; w < n[2]; w++ {

					if check%2 == 0 {
						fmt.Print(".")
					} else {
						fmt.Print("#")
					}

				}

				check += 1
			}

			if n[0]%2 == 1 {
				check += 1
			}

			fmt.Print("\n")

		}

		check += 1
	}

}
