# Video Poker

## Description

This is a simple video poker game that allows the user to play Jacks or Better video poker.

It is quite rudimentary and does not have any fancy graphics or animations. It is a text-based game that runs in the console.

As of now, the game does not support betting. The player may play as many hands as they like without any risk of losing money.

Hands are also correctly evaluated, hopefully.

## How to Play

1. Run the game by executing `go run .` in the root directory of the project.
2. You will be immediately dealt a hand of five cards, and asked which cards you would like to hold.
3. Type the index number of the cards you would like to hold, confirming with the Enter key.
4. When you are ready, press Enter without typing anything to draw new cards.
5. Your final hand will be displayed, along with the hand's ranking.

## Hand Rankings

The hand rankings are as follows:

1. Royal Flush
2. Straight Flush
3. Four of a Kind
4. Full House
5. Flush
6. Straight
7. Three of a Kind
8. Two Pair
9. Jacks or Better

## Future Improvements

1. Implement betting.
2. Add a graphical interface.
