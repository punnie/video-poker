package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
  game "github.com/punnie/video-poker/pkg"
)

type gameState struct {
  hand game.Hand

  message string
  credits int
  bet int
}

func initializeModel() *gameState {
  return &gameState {
    bet: 1,
    credits: 100,

    message: "Hello world!",

    hand: game.InitializeHand(),
  }
}

func (g *gameState) Init() tea.Cmd {
  return tea.ClearScreen
}

func (g *gameState) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
  switch msg := msg.(type) {

  case tea.KeyMsg:
    switch msg.String() {
    case "q", "ctrl+c", "esc":
      return g, tea.Quit
    case "e":
      g.message = fmt.Sprintf("Your deck has %d cards", g.hand.DeckLength())
      return g, nil
    }
  }

  return g, nil
}

func (g *gameState) View() string {
  return g.message
}

func main() {
  p := tea.NewProgram(initializeModel())

  if _, err := p.Run(); err != nil {
    fmt.Println("Argh! Error found!")
    os.Exit(1)
  }
}
