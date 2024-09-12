package main

import (
	"fmt"
)

type Restaurant interface {
	Show()
	PlaceOrder()
	Total()
}

type Warteg struct {
	Menu map[string]float64
	Order map[float64]int
}

func (p Warteg) Show() {
	i := 1
	for x, y := range p.Menu{
		fmt.Printf("%d. %s - %.2f\n", i, x, y)
		i++
	}
}

func (p Warteg) ShowOrder() {
	i := 1
	for x, y := range p.Order{
		fmt.Printf("%d. %.2f - %d\n", i, x, y)
		i++
	}
}

func (p Warteg) PlaceOrder(price float64, qty int) {
	p.Order[price] += qty
}

func (p Warteg) Total() {
	var total float64
	for x, y := range p.Order {
		total += x * float64(y)
	}

	fmt.Printf("Your total bill is: $%.2f\nThank you for dining with us!\n", total)
}

func main() {
	Pinang := Warteg{
		Menu: map[string]float64{
			"soto": 11.00,
			"telur": 6.50,
			"sop": 12.00,
			"ayam": 8.50,
			"pecel": 9.50,
		},
		Order: map[float64]int{

		},
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
			var itemname string
			var qty int
			var more string

			for {
				fmt.Print("Enter the item name: ")
				fmt.Scanln(&itemname)
				_, avail := Pinang.Menu[itemname]
				if avail {
					fmt.Print("Enter the quantity: ")
					fmt.Scanln(&qty)
					if qty > 0 {
						Pinang.PlaceOrder(Pinang.Menu[itemname], qty)
						fmt.Printf("Order Placed: %d %s(s)\n\n", qty, itemname)

						fmt.Print("Do you want to place another order? (yes/no): ")
						for {
							fmt.Scanln(&more)
							if more == "no" {
								Pinang.Total()
								return
							}
							if more == "yes" {
								break
							} else {
								fmt.Print("Invalid input. Please enter 'yes' or 'no': ")
							}
						}
						
					} else {
						fmt.Println("Quantity cannot be zero or lower!")
					}
				} else {
					fmt.Println("Item is not available!")
				}					
			}			

		case "3":
			fmt.Println("exiting...")
			return

		default:
			fmt.Println("Invalid input!")
			
		}
			
		
	}
}