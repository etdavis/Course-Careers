/* Question 4
In Go, write a function named updateInventory that accepts a pointer to a Store struct and a
Transaction struct. The function should update the inventory of the store based on the
transaction details. If the quantity in the Transaction struct is positive, it means items are added
to the inventory. If it's negative, items are removed from the inventory.
If the transaction would cause the quantity of an item in the inventory to become negative, the
function should not update the inventory and should return false. Otherwise, the function should
update the inventory and return true.
If an item does not exist in the store's inventory, the function should treat it as if its quantity is 0.
So, for example, if a transaction tries to remove an item that does not exist in the inventory, the
function should not update the inventory and should return false. Similarly, if a transaction adds
an item that does not exist in the inventory, the function should add it to the inventory with the
correct quantity and should return true.
The function signature in Go is: func updateInventory(s *Store, t Transaction) bool
The Store struct is defined as follows:

type Store struct {
	Name string
	Inventory map[string]int // maps from item names to quantity
}

The Transaction struct is defined as follows:

type Transaction struct {
	ItemName string
	Quantity int // quantity can be positive or negative
} */

//go:build practice4
//run with $ go run -tags practice4 practice4.go
package main

import (
	"fmt"
)

type Store struct {
	Name string
	Inventory map[string]int // maps from item names to quantity
}

type Transaction struct {
	ItemName string
	Quantity int // quantity can be positive or negative
}

func updateInventory(s *Store, t Transaction) bool {
	value, exists := s.Inventory[t.ItemName]
	if (!exists && t.Quantity > 0) || (exists && value + t.Quantity >= 0) {
		s.Inventory[t.ItemName] += t.Quantity
		return true
	} else {
		return false
	}
}

func main() {
	s1 := &Store{
		Name: "Store 1",
		Inventory: map[string]int{
		"Item 1": 10,
		"Item 2": 20,
		},
	}
	t1 := Transaction{
		ItemName: "Item 1",
		Quantity: -5,
	}

	s2:= &Store{
		Name: "Local Store",
		Inventory: map[string]int{
		"Apples": 100,
		"Bananas": 50,
		"Oranges": 30,
		},
	}
	t2 := Transaction{
		ItemName: "Bananas",
		Quantity: -60,
	}


	if updateInventory(s1, t1) { // true
		fmt.Println("Transaction 1 successful")
		fmt.Println(s1) // &Store{ Name: "Store 1", Inventory: map[string]int{ "Item 1": 5, "Item 2": 20, }, }
	} else {
		fmt.Println("Transaction 1 failed")
	}

	if updateInventory(s2, t2) { // false
		fmt.Println("Transaction 2 successful")
		fmt.Println(s2)
	} else {
		fmt.Println("Transaction 2 failed")
	}
}