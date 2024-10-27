package pkg

import (
	"cmp"
	"slices"
)

var (
	ranks  = []string{"2", "3", "4", "5", "6", "7", "8", "9", "T", "J", "Q", "K", "A"}
	suites = []string{"C", "D", "H", "S"}
)

<<<<<<< HEAD
type Game struct {
  hand Stack
  deck Stack

  bet int

  credits int

  state int
}

func (h Game) DeckLength() int {
  return h.deck.Len()
}

func (h Game) HandLength() int {
  return h.hand.Len()
}

func (h Game) HandCards() []Card {
  return h.hand.cards
}

func (h Game) Prize() string {
  var prizeString string
=======
type Hand struct {
	hand Stack
	deck Stack
	held [5]bool

	state int
}

func (h Hand) DeckLength() int {
	return h.deck.Len()
}

func (h Hand) HandLength() int {
	return h.hand.Len()
}

func (h Hand) HandCards() []Card {
	return h.hand.cards
}

func (h Hand) Prize() string {
	var prizeString string
>>>>>>> 397fb91 (Refactor code formatting and update project description)

	prizeString = detectPrize(h.hand.cards).String()

	return prizeString
}

func (h Hand) IsHeld(index int) bool {
	if index < 0 || index >= 5 {
		return false
	}
	return h.held[index]
}

func (h Hand) ToggleHold(index int) Hand {
	if index < 0 || index >= 5 {
		return h
	}
	h.held[index] = !h.held[index]
	return h
}

func (h Hand) Draw() Hand {
	for i := 0; i < 5; i++ {
		if !h.held[i] {
			card, deck := h.deck.RandomPop()
			h.hand.cards[i] = card
			h.deck = deck
		}
	}
	h.state = 1
	return h
}

func (h Hand) GetPrizeValue(bet int) int {
	prize := detectPrize(h.hand.cards)
	multipliers := map[int][]int{
		1: {1, 2, 3, 4, 5},     // JACKS OR HIGHER
		2: {1, 2, 3, 4, 5},     // TWO PAIR
		3: {2, 4, 6, 8, 10},    // THREE OF A KIND
		4: {3, 6, 9, 12, 15},   // STRAIGHT
		5: {5, 10, 15, 20, 25}, // FLUSH
		6: {7, 14, 21, 28, 35}, // FULL HOUSE
		7: {20, 40, 60, 80, 100}, // FOUR OF A KIND
		8: {50, 100, 150, 200, 250}, // STRAIGHT FLUSH
		9: {250, 500, 750, 1000, 4000}, // ROYAL FLUSH
	}
	
	if multiplier, exists := multipliers[prize.hand]; exists {
		if bet > 0 && bet <= 5 {
			return multiplier[bet-1]
		}
	}
	return 0
}

type prize struct {
	hand int
}

func (p prize) String() string {
	return [...]string{
		"",
		"JACKS OR HIGHER",
		"TWO PAIR",
		"THREE OF A KIND",
		"STRAIGHT",
		"FLUSH",
		"FULL HOUSE",
		"FOUR OF A KIND",
		"STRAIGHT FLUSH",
		"ROYAL FLUSH",
	}[p.hand]
}

func initializeDeck() Stack {
	var deck Stack

	for _, suite := range suites {
		for i := range ranks {
			deck = deck.Push(Card{Rank: i, Suite: suite})
		}
	}

	return deck
}

<<<<<<< HEAD
func InitializeHand() Game {
  var hand Stack
=======
func InitializeHand() Hand {
	var hand Stack
>>>>>>> 397fb91 (Refactor code formatting and update project description)

	deck := initializeDeck()

	for i := 0; i < 5; i++ {
		card := Card{}

		card, deck = deck.RandomPop()
		hand = hand.Push(card)
	}

	// fmt.Printf("Deck has %d cards\n", deck.Len())

<<<<<<< HEAD
  return Game{
    hand: hand,
    deck: deck,
  }
=======
	return Hand{
		state: 0,

		hand: hand,
		deck: deck,
	}
>>>>>>> 397fb91 (Refactor code formatting and update project description)
}

func InitializeGame() {
}

func InitializeGame() {
}

func detectPrize(h []Card) prize {
	ranks := map[int]int{}
	suites := map[string]int{}

	hand := make([]Card, len(h))
	copy(hand, h)

	cardCmp := func(a, b Card) int {
		return cmp.Compare(a.Rank, b.Rank)
	}

	slices.SortFunc(hand, cardCmp)

  // Count ranks and suites in hand
	for _, card := range hand {
		_, rank_is_present := ranks[card.Rank]
		if rank_is_present {
			ranks[card.Rank]++
		} else {
			ranks[card.Rank] = 1
		}

		_, suite_is_present := suites[card.Suite]
		if suite_is_present {
			suites[card.Suite]++
		} else {
			suites[card.Suite] = 1
		}
	}

	is_straight := true

	// Check for straight
	for i := 0; i < len(hand)-1; i++ {
		if hand[i+1].Rank != hand[i].Rank+1 {
			is_straight = false
			break
		}
	}

	// We got ourselves some type of flush
	if len(suites) == 1 {
		if is_straight {
			// Royal straight flush
			if hand[len(hand)-1].Rank == 12 {
				return prize{hand: 9}
			}

			// Straight flush
			return prize{hand: 8}
		}

		// Flush
		return prize{hand: 5}
	}

	// We either got four of a kind or a full house
	if len(ranks) == 2 {
		for _, v := range ranks {
			// Four of a kind
			if v == 4 {
				return prize{hand: 7}
			}
		}

		// Full house
		return prize{hand: 6}
	}

	// Straight
	if is_straight {
		return prize{hand: 4}
	}

	// We likely have a three of a kind or two pair
	if len(ranks) == 3 {
		for _, v := range ranks {
			// Three of a kind
			if v == 3 {
				return prize{hand: 3}
			}
		}

		// Two pair
		return prize{hand: 2}
	}

	// One pair: check if jacks or higher
	if len(ranks) == 4 {
		for r, v := range ranks {
			// Jacks or higher
			if v == 2 && r >= 9 {
				return prize{hand: 1}
			}
		}
	}

	// No prize
	return prize{hand: 0}
}
