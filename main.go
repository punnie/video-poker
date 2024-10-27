package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
  lipgloss "github.com/charmbracelet/lipgloss"
  game "github.com/punnie/video-poker/pkg"
)

type gameState struct {
  game game.Game

  message string
  credits int
  bet int
}

func initializeModel() *gameState {
  return &gameState {
    message: "Hello world!",
    game: game.InitializeHand(),
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
    case "d":
      g.game = game.InitializeHand()
      return g, nil
    case "e":
      g.message = fmt.Sprintf("Your hand has %d cards!", g.game.HandLength())
      return g, nil
    }
  }

  return g, nil
}

func (g *gameState) View() string {
  style := lipgloss.NewStyle().
      Bold(true).
      Blink(true).
      PaddingTop(2).
      PaddingBottom(2)

  return lipgloss.JoinVertical(
    lipgloss.Center, 
    payoutTableView(), 
    prizeView(g.game),
    lipgloss.JoinHorizontal(lipgloss.Center, handCardsView(g.game)...),
    style.Render(g.message),
  )
}

func main() {
  p := tea.NewProgram(initializeModel())

  if _, err := p.Run(); err != nil {
    fmt.Println("Argh! Error found!")
    os.Exit(1)
  }
}
