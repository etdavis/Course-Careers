/* Question 5
In Go, write a function named sortContacts that accepts a slice of Contact structs and returns a
sorted slice of Contact structs.
The contacts should be sorted based on the following criteria:
- Contacts are primarily sorted by the number of groups they belong to, in descending
order.
- If two contacts belong to the same number of groups, they should be sorted by their age,
in ascending order.
- If two contacts belong to the same number of groups and are of the same age, they
should be sorted by their name, in lexicographic (alphabetical) order.
The function signature in Go is: func sortContacts(contacts []Contact) []Contact
The Contact struct is defined as follows:

type Contact struct {
	Name string
	Email string
	Age int
	Groups []string // list of groups the contact belongs to
} */

//go:build practice5
//run with $ go run -tags practice5 practice5.go
package main

import (
	"fmt"
)

type Contact struct {
	Name string
	Email string
	Age int
	Groups []string // list of groups the contact belongs to
}

func sortByGroup(contacts []Contact) []Contact {
	for i := 0; i < len(contacts) - 1; i++ {
		for j := i + 1; j < len(contacts); j++ {
			if len(contacts[i].Groups) < len(contacts[j].Groups) {
				contacts[i], contacts[j] = contacts[j], contacts[i]
			}
		}
	}
	return contacts
}

func sortByAge(contacts []Contact) []Contact {
	for i := 0; i < len(contacts) - 1; i++ {
		for j := i + 1; j < len(contacts); j++ {
			if len(contacts[i].Groups) == len(contacts[j].Groups) && contacts[i].Age > contacts[j].Age {
				contacts[i], contacts[j] = contacts[j], contacts[i]
			}
		}
	}
	return contacts
}

func sortByName(contacts []Contact) []Contact {
	for i := 0; i < len(contacts) - 1; i++ {
		for j := i + 1; j < len(contacts); j++ {
			if len(contacts[i].Groups) == len(contacts[j].Groups) && contacts[i].Age == contacts[j].Age && contacts[i].Name > contacts[j].Name {
				contacts[i], contacts[j] = contacts[j], contacts[i]
			}
		}
	}
	return contacts
}

// I was at a loss on this one so I had to take a peek at the solution. I peeked but I didn't copy. Note the distinction.
func sortContacts(contacts []Contact) []Contact {
	sortByGroup(contacts)
	sortByAge(contacts)
	sortByName(contacts)
	return contacts
}

func main() {
	contacts1 := []Contact{
		Contact{
			Name: "John Doe",
			Email: "johndoe@example.com",
			Age: 30,
			Groups: []string{"Friends", "Work"},
		},
		Contact{
			Name: "Jane Doe",
			Email: "janedoe@example.com",
			Age: 25,
			Groups: []string{"Work"},
		},
		Contact{
			Name: "Alice Johnson",
			Email: "alicejohnson@example.com",
			Age: 30,
			Groups: []string{"Friends", "Work"},
		},
	}

	contacts2 := []Contact{
		Contact{
			Name: "John Doe",
			Email: "johndoe@example.com",
			Age: 35,
			Groups: []string{"Friends", "Work", "Gym", "Book Club"},
		},
		Contact{
			Name: "Alice Smith",
			Email: "alicesmith@example.com",
			Age: 30,
			Groups: []string{"Work", "Gym"},
		},
		Contact{
			Name: "Charlie Brown",
			Email: "charliebrown@example.com",
			Age: 28,
			Groups: []string{"Book Club", "Chess Club"},
		},
		Contact{
			Name: "Bob White",
			Email: "bobwhite@example.com",
			Age: 45,
			Groups: []string{"Friends", "Chess Club"},
		},
		Contact{
			Name: "Jane Doe",
			Email: "janedoe@example.com",
			Age: 32,
			Groups: []string{"Work"},
		},
		Contact{
			Name: "Mary Johnson",
			Email: "maryjohnson@example.com",
			Age: 35,
			Groups: []string{"Friends", "Work", "Gym", "Book Club"},
		},
		Contact{
			Name: "Emma Brown",
			Email: "emmabrown@example.com",
			Age: 25,
			Groups: []string{},
		},
	}

	fmt.Println(sortContacts(contacts1))
	fmt.Println(sortContacts(contacts2))
}