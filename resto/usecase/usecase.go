package usecase

import (
	w "belajar/resto/warteg"
	"fmt"
)

type Restaurant interface {
	Show()
	PlaceOrder()
	Total()
}

func Runcli() {
	Pinang := w.Warteg{
		Menu: map[string]float64{
			"soto":  11.00,
			"telur": 6.50,
			"sop":   12.00,
			"ayam":  8.50,
			"pecel": 9.50,
		},
		Order: map[float64]int{},
	}

	var options string
	fmt.Println("Welcome to the pinang restaurant!\n1. View Menu\n2. Place Order\n3. Exit")

	for {
		fmt.Print("Please select an option: ")
		fmt.Scanln(&options)
		switch options {
		case "1":
			Pinang.Show()
		case "2":
			Pinang.PlaceOrder()
			return
		case "3":
			fmt.Println("exiting...")
			return
		default:
			fmt.Println("Invalid input!")
		}
	}
}
