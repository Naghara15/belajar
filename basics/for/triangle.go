package triangle

import ("fmt")

func PrintTriangleLeft(x int) {

	for i := 1; i <= x; i++ {
		var stars string
		for y := 1; y <= i; y++ {

			stars += "*"
		}

		fmt.Println(stars)
	}
}