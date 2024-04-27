package main

import (
  "fmt"
  "math/rand"
  "time"
  "bufio"
  "os"
  "io"
  "strings"
  "strconv"
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
  ret := make([]card, 0)
  ret = append(ret, deck[:index]...)
  return deck[index], append(ret, deck[index+1:]...)
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
    if err == io.EOF {
      fmt.Println("Bye!")
      os.Exit(0)
    }
  }

  cmd := strings.TrimSpace(line)

  return cmd, nil
}

func printHand(hand []card, hold []bool) {
  fmt.Printf("Hand is ( ")

  for i, card := range hand {
    var hold_s string

    if hold[i] {
      hold_s = "H"
    } else {
      hold_s = " "
    }

    fmt.Printf("[%d: %s [%s]] ", i+1, card.toString(), hold_s)
  }

  fmt.Printf(")\n")
}

func printDeckCardNumber(deck []card) {
  fmt.Printf("Deck has %d cards!\n", len(deck))
}

type prize struct {
  hand int
}

func (p prize) toString() string {
  return [...]string{"NONE", "JACKS", "TWO PAIR", "THREE OF A KIND", "STRAIGHT", "FLUSH", "FULL HOUSE", "FOUR OF A KIND", "STRAIGHT FLUSH", "ROYAL FLUSH"}[p.hand]
}

func detectPrize(hand []card) prize {
  ranks := map[string]int{}
  suites := map[string]int{}

  for _, card := range hand {
    _, rank_is_present := ranks[card.rank]
    if  rank_is_present {
      ranks[card.rank]++ 
    } else {
      ranks[card.rank] = 1
    }

    _, suite_is_present := suites[card.suite]
    if  suite_is_present {
      suites[card.suite]++ 
    } else {
      suites[card.suite] = 1
    }
  }

  fmt.Println("Ranks: ", ranks)
  fmt.Println("Ranks length: ", len(ranks))
  fmt.Println("Suites: ", suites)
  fmt.Println("Suites length: ", len(suites))

  // We got ourselves some type of flush
  if len(suites) == 1 {
    return prize{hand: 5}
  }

  // We either got four of a kind or a full house
  if len(ranks) == 2 {
    for _, v := range ranks {
      if v == 4 {
        return prize{hand: 7}
      }
    }

    return prize{hand: 6}
  }

  // We likely have a three of a kind or two pair
  if len(ranks) == 3 {
    for _, v := range ranks {
      if v == 3 {
        return prize{hand: 3}
      }
    }

    return prize{hand: 2}
  }

  return prize{hand: 0}
}

func main() {
  for {
    var hand []card
    hold := []bool{false, false, false, false, false}

    deck := initializeDeck()
    hand, deck = dealNewHand(deck)

    deal_loop := true

    for deal_loop {
      printHand(hand, hold)
      fmt.Printf("[1..5]: Hold card  [RETURN]: Deal new hand\n")

      fmt.Printf("> ")
      cmd, _ := readUserInput()

      switch cmd {
      case "1", "2", "3", "4", "5":
        cmd_i, err := strconv.Atoi(cmd)

        if err != nil {
          panic(err)
        }

        hold[cmd_i - 1] = !hold[cmd_i - 1]

      case "":

        deal_loop = false

        for i := 0; i < 5; i++ {
          if hold[i] {
            continue
          } else {
            hand[i], deck = removeCardFromDeck(deck, r.Intn(len(deck)))
            hold[i] = false
          } 
        }

        p := detectPrize(hand)

        printHand(hand, hold)
        fmt.Printf("Prize: %s\n", p.toString())

      default:
        fmt.Printf("?\n")
        continue
      }
    }

    fmt.Printf("\n")
    _, _ = readUserInput()
  }
}

