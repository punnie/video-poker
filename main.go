package main

import (
	"fmt"
	"os"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	lipgloss "github.com/charmbracelet/lipgloss"
	table "github.com/charmbracelet/lipgloss/table"
	game "github.com/punnie/video-poker/pkg"
)

type gameState struct {
	hand game.Hand

	message string
	credits int
	bet     int
	gamePhase int // 0: initial deal, 1: hold/draw phase, 2: final result
}

func initializeModel() *gameState {
	return &gameState{
		bet:     1,
		credits: 100,
		gamePhase: -1,

		message: "Welcome to Video Poker! Press 'n' to start a new game",

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
		case "n": // New game
			if g.credits >= g.bet {
				g.credits -= g.bet
				g.hand = game.InitializeHand()
				g.gamePhase = 0
				g.message = "Select cards to hold (1-5) then press SPACE to draw"
			} else {
				g.message = "Not enough credits!"
			}
			return g, nil
		case "1", "2", "3", "4", "5":
			if g.gamePhase == 0 {
				cardIndex := int(msg.String()[0] - '1')
				g.hand = g.hand.ToggleHold(cardIndex)
				g.message = "Select cards to hold (1-5) then press SPACE to draw"
			}
			return g, nil
		case " ": // Space to draw
			if g.gamePhase == 0 {
				g.hand = g.hand.Draw()
				g.gamePhase = 1
				prizeValue := g.hand.GetPrizeValue(g.bet)
				g.credits += prizeValue
				if prizeValue > 0 {
					g.message = fmt.Sprintf("You won %d credits! Press 'n' for new game", prizeValue)
				} else {
					g.message = "No win. Press 'n' for new game"
				}
			}
			return g, nil
		case "+":
			if g.bet < 5 {
				g.bet++
				g.message = fmt.Sprintf("Bet: %d credits", g.bet)
			}
			return g, nil
		case "-":
			if g.bet > 1 {
				g.bet--
				g.message = fmt.Sprintf("Bet: %d credits", g.bet)
			}
			return g, nil
		}
	}

	return g, nil
}

var (
	cardBorder = lipgloss.Border{
		Top:         "─",
		Bottom:      "─",
		Left:        "│",
		Right:       "│",
		TopLeft:     "╭",
		TopRight:    "╮",
		BottomLeft:  "╰",
		BottomRight: "╯",
	}
)

func cardView(c game.Card, visible bool) string {
	var view string
	var design string
	var label string
	var reverse_label string
	var color lipgloss.Color

	const width = 7
	const height = 5

	// Card background design
	if visible {
		design = " "
		label = c.String()
		reverse_label = c.ReverseString()

		if c.Suite == "C" || c.Suite == "S" {
			color = lipgloss.Color("#000000")
		} else {
			color = lipgloss.Color("#FF0000")
		}

	} else {
		design = "╱"
		label = "╱"

		color = lipgloss.Color("#FFFFFF")
	}

	style := lipgloss.NewStyle().Foreground(color).Background(lipgloss.Color("#FFFFFF")).Bold(true)
	view = style.Render(label+strings.Repeat(design, width-lipgloss.Width(label))) + "\n"

	for i := 0; i < height-1; i++ {
		view += style.Render(strings.Repeat(design, width)) + "\n"
	}

	view += style.Render(strings.Repeat(design, width-lipgloss.Width(reverse_label)) + reverse_label)

	borderStyle := lipgloss.NewStyle().Foreground(color).Background(lipgloss.Color("#FFFFFF")).Padding(0, 1).Margin(1)
	return borderStyle.Render(view)
}

func cardViews(h game.Hand, gamePhase int) []string {
	var views []string

	for i, card := range h.HandCards() {
		cardStr := cardView(card, true)
		
		// Add hold indicator and button number
		holdIndicator := ""
		if gamePhase == 0 {
			if h.IsHeld(i) {
				holdIndicator = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("#00FF00")).Render("[HELD]")
			} else {
				holdIndicator = lipgloss.NewStyle().Foreground(lipgloss.Color("#888888")).Render("[    ]")
			}
		}
		
		buttonNumber := lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("#FFCC00")).Render(fmt.Sprintf("  %d  ", i+1))
		
		cardWithButton := lipgloss.JoinVertical(lipgloss.Center, cardStr, holdIndicator, buttonNumber)
		views = append(views, cardWithButton)
	}

	return views
}

func payoutTableView() string {
	rows := [][]string{
		{"ROYAL FLUSH", "250", "500", "750", "1000", "4000"},
		{"STRAIGHT FLUSH", "50", "100", "150", "200", "250"},
		{"FOUR OF A KIND", "20", "40", "60", "80", "100"},
		{"FULL HOUSE", "7", "14", "21", "28", "35"},
		{"FLUSH", "5", "10", "15", "20", "25"},
		{"STRAIGHT", "3", "6", "9", "12", "15"},
		{"THREE OF A KIND", "2", "4", "6", "8", "10"},
		{"TWO PAIR", "1", "2", "3", "4", "5"},
		{"JACKS OR HIGHER", "1", "2", "3", "4", "5"},
	}

	t := table.New().
		Border(lipgloss.NormalBorder()).
		BorderStyle(lipgloss.NewStyle().Foreground(lipgloss.Color("#FFCC00"))).
		Rows(rows...).
		StyleFunc(func(row, col int) lipgloss.Style {
			if col == 0 {
				return lipgloss.NewStyle().Padding(0, 1)
			} else {
				return lipgloss.NewStyle().Padding(0, 1).Align(lipgloss.Right)
			}
		})

	return t.Render()
}

func prizeView(h game.Hand) string {
	var view string

	view = h.Prize()

	return view
}

func (g *gameState) View() string {
	style := lipgloss.NewStyle().
		Bold(true).
		Blink(true).
		PaddingTop(2).
		PaddingBottom(2)

	// Credit and bet information
	creditInfo := lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("#00FF00")).Render(
		fmt.Sprintf("Credits: %d | Bet: %d | Press +/- to change bet, 'n' for new game", g.credits, g.bet))

	return lipgloss.JoinVertical(
		lipgloss.Center,
		payoutTableView(),
		creditInfo,
		prizeView(g.hand),
		lipgloss.JoinHorizontal(lipgloss.Center, cardViews(g.hand, g.gamePhase)...),
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
