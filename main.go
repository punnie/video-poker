package main

import (
  "fmt"
  "math/rand"
  "time"
)

const (
  clubs string = "♣"
  diamonds string = "♦"
  hearts string = "♥"
  spades string = "♠"
)

var (
  ranks = []string{"2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K", "A"}
  suites = []string{clubs, diamonds, hearts, spades}

  r = rand.New(rand.NewSource(time.Now().UnixNano()))
)

type card struct {
  rank string
  suite string
}

func (c card) toString() string {
  return fmt.Sprintf("%s%s", c.suite, c.rank)
}

func initializeDeck() []card {
  var deck []card

  for _, suite := range suites {
    for _, rank := range ranks {
      deck = append(deck, card{rank: rank, suite: suite})
    }
  }

  return deck
}

func removeCardFromDeck(deck []card, index int) (card, []card) {
  return deck[index], append(deck[:index], deck[index+1:]...)
}

func dealNewHand(deck []card) ([]card, []card) {
  var hand []card
  return dealNewHandRecur(deck, hand)
}

func dealNewHandRecur(deck []card, hand []card) ([]card, []card) {
  card, deck := removeCardFromDeck(deck, r.Intn(len(deck)))
  hand = append(hand, card)

  if len(hand) > 4 { 
    return hand, deck
  } else {
    return dealNewHandRecur(deck, hand)
  }
}

func main() {
  // r := rand.New(rand.NewSource(99)) // For testing purposes

  deck := initializeDeck()

  // for _, card := range deck {
  //   fmt.Printf("%s\n", card.toString())
  // }

  fmt.Printf("Deck has %d cards!\n", len(deck))

  hand, deck := dealNewHand(deck)

  fmt.Printf("Random hand is ( ")

  for i, card := range hand {
    fmt.Printf("%d:%s, ", i, card.toString())
  }

  fmt.Printf(")\n")
  fmt.Printf("Deck has %d cards!\n", len(deck))
}

