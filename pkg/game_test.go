package pkg

import (
  "testing"
)

// Test that a Full House is detected correctly
func TestDetectPrizeFullHouse(t *testing.T) {
  hand := []Card{
    {Rank: 12, Suite: "S"},
    {Rank: 12, Suite: "H"},
    {Rank: 12, Suite: "C"},
    {Rank: 11, Suite: "D"},
    {Rank: 11, Suite: "S"},
  }

  prize := detectPrize(hand)

  if prize.hand != 6 {
    t.Errorf("Expected a Full House, got %d", prize.hand)
  }
}

// Test that a Flush is detected correctly
func TestDetectPrizeFlush(t *testing.T) {
  hand := []Card{
    {Rank: 12, Suite: "S"},
    {Rank: 11, Suite: "S"},
    {Rank: 10, Suite: "S"},
    {Rank: 9,  Suite: "S"},
    {Rank: 7,  Suite: "S"},
  }

  prize := detectPrize(hand)

  if prize.hand != 5 {
    t.Errorf("Expected a Flush, got %d", prize.hand)
  }
}

// Test that a Royal Flush is detected correctly
func TestDetectPrizeRoyalFlush(t *testing.T) {
  hand := []Card{
    {Rank: 12, Suite: "S"},
    {Rank: 11, Suite: "S"},
    {Rank: 10, Suite: "S"},
    {Rank: 9,  Suite: "S"},
    {Rank: 8,  Suite: "S"},
  }

  prize := detectPrize(hand)

  if prize.hand != 9 {
    t.Errorf("Expected a Royal Flush, got %d", prize.hand)
  }
}

// Test that a Straight Flush is detected correctly
func TestDetectPrizeStraightFlush(t *testing.T) {
  hand := []Card{
    {Rank: 8, Suite: "S"},
    {Rank: 7, Suite: "S"},
    {Rank: 6, Suite: "S"},
    {Rank: 5, Suite: "S"},
    {Rank: 4, Suite: "S"},
  }

  prize := detectPrize(hand)

  if prize.hand != 8 {
    t.Errorf("Expected a Straight Flush, got %d", prize.hand)
  }
}

// Test that a Four of a Kind is detected correctly
func TestDetectPrizeFourOfAKind(t *testing.T) {
  hand := []Card{
    {Rank: 12, Suite: "S"},
    {Rank: 12, Suite: "H"},
    {Rank: 12, Suite: "C"},
    {Rank: 12, Suite: "D"},
    {Rank: 11, Suite: "S"},
  }

  prize := detectPrize(hand)

  if prize.hand != 7 {
    t.Errorf("Expected a Four of a Kind, got %d", prize.hand)
  }
}

// Test that a Straight is detected correctly
func TestDetectPrizeStraight(t *testing.T) {
  hand := []Card{
    {Rank: 8, Suite: "S"},
    {Rank: 7, Suite: "H"},
    {Rank: 6, Suite: "C"},
    {Rank: 5, Suite: "D"},
    {Rank: 4, Suite: "S"},
  }

  prize := detectPrize(hand)

  if prize.hand != 4 {
    t.Errorf("Expected a Straight, got %d", prize.hand)
  }
}

// Test that a Three of a Kind is detected correctly
func TestDetectPrizeThreeOfAKind(t *testing.T) {
  hand := []Card{
    {Rank: 12, Suite: "S"},
    {Rank: 12, Suite: "H"},
    {Rank: 12, Suite: "C"},
    {Rank: 11, Suite: "D"},
    {Rank: 10, Suite: "S"},
  }

  prize := detectPrize(hand)

  if prize.hand != 3 {
    t.Errorf("Expected a Three of a Kind, got %d", prize.hand)
  }
}

// Test that a Two Pair is detected correctly
func TestDetectPrizeTwoPair(t *testing.T) {
  hand := []Card{
    {Rank: 12, Suite: "S"},
    {Rank: 12, Suite: "H"},
    {Rank: 11, Suite: "C"},
    {Rank: 11, Suite: "D"},
    {Rank: 10, Suite: "S"},
  }

  prize := detectPrize(hand)

  if prize.hand != 2 {
    t.Errorf("Expected a Two Pair, got %d", prize.hand)
  }
}

func TestDetectPrizeJacksOrHigher(t *testing.T) {
  hand := []Card{
    {Rank: 1, Suite: "S"},
    {Rank: 4, Suite: "H"},
    {Rank: 9, Suite: "C"},
    {Rank: 8, Suite: "D"},
    {Rank: 9, Suite: "S"},
  }

  prize := detectPrize(hand)

  if prize.hand != 1 {
    t.Errorf("Expected Jacks or Higher, got %d", prize.hand)
  }
}

