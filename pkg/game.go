package pkg

import (
	"cmp"
	"slices"
)

var (
	ranks  = []string{"2", "3", "4", "5", "6", "7", "8", "9", "T", "J", "Q", "K", "A"}
	suites = []string{"C", "D", "H", "S"}
)

type Hand struct {
  hand Stack
  deck Stack

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

type prize struct {
	hand int
}

func (p prize) toString() string {
	return [...]string{
    "NONE", 
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

func InitializeHand() Hand {
  var hand Stack

  deck := initializeDeck()

  for i := 0; i < 5; i++ {
    card := Card{}

    card, deck = deck.RandomPop()
    hand = hand.Push(card)
  }

  // fmt.Printf("Deck has %d cards\n", deck.Len())

  return Hand{
    state: 0,

    hand: hand,
    deck: deck,
  }
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

	// fmt.Println("Sorted hand: ", hand)

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

	// fmt.Println("Ranks: ", ranks)
	// fmt.Println("Ranks length: ", len(ranks))
	// fmt.Println("Suites: ", suites)
	// fmt.Println("Suites length: ", len(suites))

	is_straight := true

	// Check for straight
	for i := 0; i < len(hand)-1; i++ {
		if hand[i + 1].Rank != hand[i].Rank + 1 {
			is_straight = false
			break
		}
	}

	// We got ourselves some type of flush
	if len(suites) == 1 {
		if is_straight {
			// Royal straight flush
			if hand[len(hand) - 1].Rank == 12 {
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
