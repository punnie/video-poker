package pkg

import (
  "fmt"
  "math/rand"
  "time"
)

const (
  clubs string = "♧"
  diamonds string = "♦"
  hearts string = "♡"
  spades string = "♠"
)

var (
	// r = rand.New(rand.NewSource(99)) // For testing purposes
	r = rand.New(rand.NewSource(time.Now().UnixNano()))
)

type Card struct {
  Rank int
  Suite string
}

func suiteToString(suite string) string {
  return map[string]string{
    "C": clubs,
    "D": diamonds,
    "H": hearts,
    "S": spades,
  }[suite]
}

func rankToString(rank int) string {
  return map[int]string{
    0: "2",
    1: "3",
    2: "4",
    3: "5",
    4: "6",
    5: "7",
    6: "8",
    7: "9",
    8: "10",
    9: "J",
    10: "Q",
    11: "K",
    12: "A",
  }[rank]
}

func (c Card) String() string {
  return fmt.Sprintf("%s%s", rankToString(c.Rank), suiteToString(c.Suite))
}

func (c Card) ReverseString() string {
  return fmt.Sprintf("%s%s", suiteToString(c.Suite), rankToString(c.Rank))
}

type Stack struct {
  cards []Card
}

func (s Stack) Len() int {
  return len(s.cards)
}

func (s Stack) Push(c Card) Stack {
  s.cards = append(s.cards, c)

  return s
}

func (s Stack) PopAndRemove(index int) (Card, Stack) {
  c := s.cards[index]
  s.cards = append(s.cards[:index], s.cards[index+1:]...)

  return c, s
}

func (s Stack) RandomPop() (Card, Stack) {
  index := 0

  if s.Len() > 1 {
    index = r.Intn(s.Len())
  }

  return s.PopAndRemove(index)
}

