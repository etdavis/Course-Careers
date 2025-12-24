package main

import (
	"bufio"
	"fmt"
	"math"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

/*
The structure of this program is mostly courtesy of CourseCareers. I filled in the functions (except getString, enterString, shuffle, and main and getCardValues) withan appropriate amount of help from AI
*/

type Card struct {
	value int
	suit  int // 0 - spades, 1 - hearts, 2 - diamonds, 3 - clubs
}

func (card Card) getString() string {
	var suit string
	var value string

	switch card.suit {
	case 0:
		suit = "♠"
	case 1:
		suit = "♥"
	case 2:
		suit = "♦"
	case 3:
		suit = "♣"
	}

	switch card.value {
	case 11:
		value = "J"
	case 12:
		value = "Q"
	case 13:
		value = "K"
	case 1:
		value = "A"
	default:
		value = fmt.Sprintf("%d", card.value)
	}

	return value + suit
}

type Deck struct {
	cards []Card
}

// deals num cards, returns slice of dealt cards
func (d *Deck) deal(num uint) []Card {
	hand := []Card{}
	for i := uint(0); i < num; i++ {
		hand = append(hand, d.cards[0])
		d.cards = d.cards[1:]
	}
	return hand
}

func (d *Deck) create() {
	for v := 1; v <= 13; v++ {
		for s := 0; s < 4; s++ {
			card := Card{value: v, suit: s}
			d.cards = append(d.cards, card)
		}
	}
	d.shuffle()
}

func (d *Deck) shuffle() {
	rand.Shuffle(len(d.cards), func(i, j int) { d.cards[i], d.cards[j] = d.cards[j], d.cards[i] })
}

type Game struct {
	deck        Deck
	playerCards []Card
	dealerCards []Card
}

func (game *Game) dealStartingCards() {
	game.deck.create()
	game.playerCards = game.deck.deal(2)
	game.dealerCards = game.deck.deal(2)
}

func (game *Game) play(bet float64) float64 {
	game.dealStartingCards()
	playerTotal := getCardValues(game.playerCards)
	dealerTotal := getCardValues(game.dealerCards)

	fmt.Println("Player has been dealt: " + game.playerCards[0].getString() + " " + game.playerCards[1].getString() + " = " + strconv.Itoa(playerTotal) + "\n")
	fmt.Println("Dealer shows: " + game.dealerCards[0].getString() + " ??\n")
	
	if dealerTotal == 21 && playerTotal != 21 {
		fmt.Println("Dealer shows other card is: " + game.dealerCards[1].getString() + "\n")
		fmt.Println("Dealer has blackjack! You lose.\n")
		return -bet
	}
	if playerTotal == 21 && dealerTotal != 21 {
		fmt.Println("Blackjack! You win 1.5x your bet.\n")
		return bet * 1.5
	}
	if playerTotal == 21 && dealerTotal == 21 {
		fmt.Println("Both you and the dealer have blackjack! Push.\n")
		return 0
	}
	
	playerTotal = game.playerTurn(&playerTotal)
	if playerTotal > 21 {
		fmt.Println("You busted! You lose.\n")
		return -bet
	}

	dealerTotal = game.dealerTurn(&dealerTotal)
	if dealerTotal > 21 {
		fmt.Println("Dealer busted! You win.\n")
		return bet
	}
	if playerTotal > dealerTotal {
		fmt.Println("You win!\n")
		return bet
	}
	if playerTotal < dealerTotal {
		fmt.Println("You lose!\n")
		return -bet
	}
	fmt.Println("Push.\n")
	return 0
}

func (game *Game) playerTurn(playerTotal *int) int {
	for true {
		fmt.Printf("Hit or stay? (H/S): ")
		choice := strings.ToLower(enterString())
		switch choice {
		case "s":
			return *playerTotal
		case "h":
			newCard := game.deck.deal(1)[0]
			game.playerCards = append(game.playerCards, newCard)
			*playerTotal = getCardValues(game.playerCards)
			fmt.Println("You are dealt: " + newCard.getString())
			fmt.Println("You now have: ")
			for i := 0; i < len(game.playerCards); i++ {
				fmt.Print(game.playerCards[i].getString() + " ")
			}
			fmt.Println("= " + strconv.Itoa(*playerTotal) + "\n")
			if *playerTotal >= 21 {
				return *playerTotal
			}
		default:
			fmt.Println("Please enter H or S.")
			continue
		}
	}
	return *playerTotal
}

func (game *Game) dealerTurn(dealerTotal *int) int {
	fmt.Println("Dealer reveals: " + game.dealerCards[0].getString() + " " + game.dealerCards[1].getString() + " = " + strconv.Itoa(*dealerTotal) + "\n")
	for *dealerTotal < 17 {
		newCard := game.deck.deal(1)[0]
		game.dealerCards = append(game.dealerCards, newCard)
		*dealerTotal = getCardValues(game.dealerCards)
		fmt.Println("Dealer hits and gets: " + newCard.getString())
		fmt.Println("Dealer now has: ")
		for i := 0; i < len(game.dealerCards); i++ {
			fmt.Print(game.dealerCards[i].getString() + " ")
			}
		fmt.Println("= " + strconv.Itoa(*dealerTotal) + "\n")
		if *dealerTotal >= 21 {
			return *dealerTotal
		}
	}
	fmt.Println("Dealer stays.\n")
	return *dealerTotal
}

//borrowed this function from CourseCareers solution to fix a bug with calculating card totals
func getCardValues(cards []Card) int {
	aces := 0
	value := 0

	for _, card := range cards {
		if card.value == 1 {
			aces++
		} else {
			value += int(math.Min(10, float64(card.value)))
		}
	}

	if aces == 0 {
		return value
	}

	if value < (11 + aces - 1) {
		return value + 11 + aces - 1
	} else {
		return value + aces
	}
}

func enterString() string {
	reader := bufio.NewReader(os.Stdin)
	// ReadString will block until the delimiter is entered
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("An error occurred while reading input. Please try again", err)
		return ""
	}

	// remove the delimiter from the string
	input = strings.TrimSpace(input)
	return input
}

func main() {
	balance := float64(100)

	for balance > 0 {
		fmt.Printf("Your balance is: $%.2f\n", balance)
		fmt.Printf("Enter your bet (q to quit): ")
		bet, err := strconv.ParseFloat(enterString(), 64)
		if err != nil {
			break
		}
		if bet > balance || bet <= 0 {
			fmt.Println("Invalid bet.")
			continue
		}

		game := Game{}
		balance += game.play(bet)
	}

	fmt.Printf("You left with: $%2.f\n", balance)
}
