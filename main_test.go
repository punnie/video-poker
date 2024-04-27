package main

import "testing"

// Test that the deck is initialized correctly
func TestInitializeDeck(t *testing.T) {
  deck := initializeDeck()

  if len(deck) != 52 {
    t.Errorf("Expected 52 cards in the deck, got %d", len(deck))
  }
}

// Test that a card is removed from the deck correctly
func TestRemoveCard(t *testing.T) {
  deck := initializeDeck()
  _, deck = removeCardFromDeck(deck, 0)

  if len(deck) != 51 {
    t.Errorf("Expected 51 cards in the deck, got %d", len(deck))
  }
}

// Test that a Full House is detected correctly
func TestDetectPrizeFullHouse(t *testing.T) {
  hand := []card{
    {rank: "A", rank_i: 12, suite: "S"},
    {rank: "A", rank_i: 12, suite: "H"},
    {rank: "A", rank_i: 12, suite: "C"},
    {rank: "K", rank_i: 11, suite: "D"},
    {rank: "K", rank_i: 11, suite: "S"},
  }

  prize := detectPrize(hand)

  if prize.hand != 6 {
    t.Errorf("Expected a Full House, got %d", prize.hand)
  }
}

// Test that a Flush is detected correctly
func TestDetectPrizeFlush(t *testing.T) {
  hand := []card{
    {rank: "A", rank_i: 12, suite: "S"},
    {rank: "K", rank_i: 11, suite: "S"},
    {rank: "Q", rank_i: 10, suite: "S"},
    {rank: "J", rank_i: 9, suite: "S"},
    {rank: "9", rank_i: 7, suite: "S"},
  }

  prize := detectPrize(hand)

  if prize.hand != 5 {
    t.Errorf("Expected a Flush, got %d", prize.hand)
  }
}

// Test that a Royal Flush is detected correctly
func TestDetectPrizeRoyalFlush(t *testing.T) {
  hand := []card{
    {rank: "A", rank_i: 12, suite: "S"},
    {rank: "K", rank_i: 11, suite: "S"},
    {rank: "Q", rank_i: 10, suite: "S"},
    {rank: "J", rank_i: 9, suite: "S"},
    {rank: "10", rank_i: 8, suite: "S"},
  }

  prize := detectPrize(hand)

  if prize.hand != 9 {
    t.Errorf("Expected a Royal Flush, got %d", prize.hand)
  }
}

// Test that a Straight Flush is detected correctly
func TestDetectPrizeStraightFlush(t *testing.T) {
  hand := []card{
    {rank: "10", rank_i: 8, suite: "S"},
    {rank: "9", rank_i: 7, suite: "S"},
    {rank: "8", rank_i: 6, suite: "S"},
    {rank: "7", rank_i: 5, suite: "S"},
    {rank: "6", rank_i: 4, suite: "S"},
  }

  prize := detectPrize(hand)

  if prize.hand != 8 {
    t.Errorf("Expected a Straight Flush, got %d", prize.hand)
  }
}

// Test that a Four of a Kind is detected correctly
func TestDetectPrizeFourOfAKind(t *testing.T) {
  hand := []card{
    {rank: "A", rank_i: 12, suite: "S"},
    {rank: "A", rank_i: 12, suite: "H"},
    {rank: "A", rank_i: 12, suite: "C"},
    {rank: "A", rank_i: 12, suite: "D"},
    {rank: "K", rank_i: 11, suite: "S"},
  }

  prize := detectPrize(hand)

  if prize.hand != 7 {
    t.Errorf("Expected a Four of a Kind, got %d", prize.hand)
  }
}

// Test that a Straight is detected correctly
func TestDetectPrizeStraight(t *testing.T) {
  hand := []card{
    {rank: "10", rank_i: 8, suite: "S"},
    {rank: "9", rank_i: 7, suite: "H"},
    {rank: "8", rank_i: 6, suite: "C"},
    {rank: "7", rank_i: 5, suite: "D"},
    {rank: "6", rank_i: 4, suite: "S"},
  }

  prize := detectPrize(hand)

  if prize.hand != 4 {
    t.Errorf("Expected a Straight, got %d", prize.hand)
  }
}

// Test that a Three of a Kind is detected correctly
func TestDetectPrizeThreeOfAKind(t *testing.T) {
  hand := []card{
    {rank: "A", rank_i: 12, suite: "S"},
    {rank: "A", rank_i: 12, suite: "H"},
    {rank: "A", rank_i: 12, suite: "C"},
    {rank: "K", rank_i: 11, suite: "D"},
    {rank: "Q", rank_i: 10, suite: "S"},
  }

  prize := detectPrize(hand)

  if prize.hand != 3 {
    t.Errorf("Expected a Three of a Kind, got %d", prize.hand)
  }
}

// Test that a Two Pair is detected correctly
func TestDetectPrizeTwoPair(t *testing.T) {
  hand := []card{
    {rank: "A", rank_i: 12, suite: "S"},
    {rank: "A", rank_i: 12, suite: "H"},
    {rank: "K", rank_i: 11, suite: "C"},
    {rank: "K", rank_i: 11, suite: "D"},
    {rank: "Q", rank_i: 10, suite: "S"},
  }

  prize := detectPrize(hand)

  if prize.hand != 2 {
    t.Errorf("Expected a Two Pair, got %d", prize.hand)
  }
}

