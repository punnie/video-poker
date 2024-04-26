package main

import (
  "fmt"
  "math/rand"
  "time"
  "bufio"
  "os"
  "io"
  "strings"
)

const (
  clubs string = "♣"
  diamonds string = "♦"
  hearts string = "♥"
  spades string = "♠"
)

var (
  ranks = []string{"2", "3", "4", "5", "6", "7", "8", "9", "T", "J", "Q", "K", "A"}
  suites = []string{clubs, diamonds, hearts, spades}

  // r = rand.New(rand.NewSource(99)) // For testing purposes
  r = rand.New(rand.NewSource(time.Now().UnixNano()))
)

type card struct {
  rank string
  suite string
}

func (c card) toString() string {
  return fmt.Sprintf("%s%s", c.rank, c.suite)
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

func readUserInput() (string, error) {

    reader := bufio.NewReader(os.Stdin)

    line, err := reader.ReadString('\n')

    if err != nil {
      return "", err
    }

    cmd := strings.TrimSpace(line)
    return cmd, nil
}

func printHand(hand []card) {
  fmt.Printf("Hand is ( ")

  for i, card := range hand {
    fmt.Printf("[%d: %s] ", i+1, card.toString())
  }

  fmt.Printf(")\n")
}

func printDeckCardNumber(deck []card) {
  fmt.Printf("Deck has %d cards!\n", len(deck))
}

func main() {
  deck := initializeDeck()

  // for _, card := range deck {
  //   fmt.Printf("%s\n", card.toString())
  // }

  printDeckCardNumber(deck)

  hand, deck := dealNewHand(deck)

  printHand(hand)

  printDeckCardNumber(deck)

  for {
    fmt.Printf("> ")
    cmd, err := readUserInput()

    if err != nil {
      if err == io.EOF {
        fmt.Println("Bye!")
        os.Exit(0)
      }
    }

    switch cmd {
    case "1", "2", "3", "4", "5":
      fmt.Printf("Hold card #%s\n", cmd)
    case "D":
      fmt.Printf("Deal new hand\n")
    default:
      fmt.Printf("Unknown command? Try again.\n")
      continue
    }
  }
}

