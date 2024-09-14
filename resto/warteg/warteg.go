package warteg

import (
	c "belajar/resto/constants"
	"fmt"
)

type Warteg struct {
	Menu  map[string]float64
	Order map[float64]int
}

func (p Warteg) Show() {
	i := 1
	for x, y := range p.Menu {
		fmt.Printf("%d. %s - %.2f\n", i, x, y)
		i++
	}
}

func (p Warteg) ShowOrder() {
	i := 1
	for x, y := range p.Order {
		fmt.Printf("%d. %.2f - %d\n", i, x, y)
		i++
	}
}

func (p Warteg) PlaceOrder() {
	var itemname, more string
	var qty int

out:
	for {
		fmt.Print("Enter the item name: ")
		fmt.Scanln(&itemname)
		price, avail := p.Menu[itemname]
		if avail {
			fmt.Print("Enter the quantity: ")
			fmt.Scanln(&qty)

			if qty <= 0 {
				fmt.Println("Quantity cannot be zero or lower!")
				continue
			}

			p.Order[price] += qty
			fmt.Printf("Order Placed: %d %s(s)\n\n", qty, itemname)
			fmt.Print("Do you want to place another order? (yes/no): ")

			for {
				fmt.Scanln(&more)
				switch more {
				case c.No:
					p.Total()
					return
				case c.Yes:
					continue out
				default:
					fmt.Print("Invalid input. Please enter 'yes' or 'no': ")
				}
			}

		}
		fmt.Println("Item is not available!")
	}

}

func (p Warteg) Total() {
	var total float64
	for x, y := range p.Order {
		total += x * float64(y)
	}

	fmt.Printf("Your total bill is: $%.2f\nThank you for dining with us!\n", total)
}
