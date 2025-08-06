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

	message        string
	credits        int
	bet            int
	gamePhase      int // 0: initial deal, 1: hold/draw phase, 2: final result, 3: high score entry
	highScores     game.HighScores
	isHighScore    bool
	initials       []rune
	initialsCursor int
}

func initializeModel() *gameState {
	highScores, err := game.LoadHighScores()
	if err != nil {
		// Handle error, maybe log it or set a default
		highScores = game.HighScores{}
	}

	return &gameState{
		bet:     1,
		credits: 100,
		gamePhase: -1,

		message: "Welcome to Video Poker! Press SPACE to start a new game",

		hand:           game.InitializeHand(),
		highScores:     highScores,
		initials:       []rune{'A', 'A', 'A'},
		initialsCursor: 0,
	}
}

func (g *gameState) Init() tea.Cmd {
	return tea.ClearScreen
}

func (g *gameState) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.KeyMsg:
		if g.gamePhase == 3 {
			switch msg.String() {
			case "q", "ctrl+c", "esc":
				return g, tea.Quit
			case " ":
				g.highScores = g.highScores.Add(game.HighScore{
					Initials: string(g.initials),
					Score:    g.credits,
				})
				g.highScores.Save()
				return g, tea.Quit
			case "left":
				if g.initialsCursor > 0 {
					g.initialsCursor--
				}
				return g, nil
			case "right":
				if g.initialsCursor < 2 {
					g.initialsCursor++
				}
				return g, nil
			case "up":
				g.initials[g.initialsCursor]++
				if g.initials[g.initialsCursor] > 'Z' {
					g.initials[g.initialsCursor] = 'A'
				}
				return g, nil
			case "down":
				g.initials[g.initialsCursor]--
				if g.initials[g.initialsCursor] < 'A' {
					g.initials[g.initialsCursor] = 'Z'
				}
				return g, nil
			}
		}

		switch msg.String() {
		case "h":
			g.gamePhase = 4
			g.message = "High Scores. Press SPACE to return to the game"
			return g, nil
		case "q", "ctrl+c", "esc":
			if g.gamePhase == 4 {
				g.gamePhase = -1
				g.message = "Welcome to Video Poker! Press SPACE to start a new game"
				return g, nil
			}
			if g.highScores.IsHighScore(g.credits) {
				g.gamePhase = 3
				g.message = "New high score! Enter your initials"
				return g, nil
			}
			return g, tea.Quit
		case "1", "2", "3", "4", "5":
			if g.gamePhase == 0 {
				cardIndex := int(msg.String()[0] - '1')
				g.hand = g.hand.ToggleHold(cardIndex)
				g.message = "Select cards to hold (1-5) then press SPACE to draw"
			}
			return g, nil
		case " ": // Space - multi-purpose based on game state
			if g.gamePhase == -1 || g.gamePhase == 1 {
				// Start a new game if we're at the initial state or after a hand is complete
				if g.credits >= g.bet {
					g.credits -= g.bet
					g.hand = game.InitializeHand()
					g.gamePhase = 0
					g.message = "Select cards to hold (1-5) then press SPACE to draw"
				} else {
					g.message = "Not enough credits!"
				}
			} else if g.gamePhase == 0 {
				// Draw cards if we're in the hold/draw phase
				g.hand = g.hand.Draw()
				g.gamePhase = 1
				prizeValue := g.hand.GetPrizeValue(g.bet)
				g.credits += prizeValue
				if prizeValue > 0 {
					g.message = fmt.Sprintf("You won %d credits! Press SPACE for new game", prizeValue)
				} else {
					g.message = "No win. Press SPACE for new game"
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

	// Improved color palette
	creditColor    = lipgloss.Color("#22C55E")  // Softer green
	betColor       = lipgloss.Color("#3B82F6")  // Blue
	holdColor      = lipgloss.Color("#10B981")  // Emerald
	buttonColor    = lipgloss.Color("#6366F1")  // Indigo
	borderColor    = lipgloss.Color("#64748B")  // Slate
	redSuitColor   = lipgloss.Color("#DC2626")  // Softer red
	blackSuitColor = lipgloss.Color("#1F2937")  // Dark gray
	messageColor   = lipgloss.Color("#374151")  // Gray
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
			color = blackSuitColor
		} else {
			color = redSuitColor
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
				holdIndicator = lipgloss.NewStyle().Bold(true).Foreground(holdColor).Render("[HELD]")
			} else {
				holdIndicator = lipgloss.NewStyle().Foreground(lipgloss.Color("#9CA3AF")).Render("[    ]")
			}
		}
		
		buttonNumber := lipgloss.NewStyle().Bold(true).Foreground(buttonColor).Render(fmt.Sprintf("  %d  ", i+1))
		
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
		BorderStyle(lipgloss.NewStyle().Foreground(borderColor)).
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
	prize := h.Prize()
	if prize == "" {
		return ""
	}
	
	// Style the prize with a nice color
	prizeStyle := lipgloss.NewStyle().
		Bold(true).
		Foreground(holdColor).
		Padding(0, 1)
	
	return prizeStyle.Render(prize)
}

func highScoreEntryView(g *gameState) string {
	var s strings.Builder
	s.WriteString("New High Score!\n\n")
	s.WriteString("Enter your initials:\n\n")

	for i, r := range g.initials {
		style := lipgloss.NewStyle().Padding(0, 1)
		if i == g.initialsCursor {
			style = style.Bold(true).Foreground(lipgloss.Color("205"))
		}
		s.WriteString(style.Render(string(r)))
	}

	s.WriteString("\n\nUse arrow keys to change initials, space to confirm.")
	return s.String()
}

func highScoreTableView(g *gameState) string {
	rows := [][]string{
		{"Initials", "Score"},
	}
	for _, hs := range g.highScores {
		rows = append(rows, []string{hs.Initials, fmt.Sprintf("%d", hs.Score)})
	}

	t := table.New().
		Border(lipgloss.NormalBorder()).
		BorderStyle(lipgloss.NewStyle().Foreground(borderColor)).
		Rows(rows...).
		StyleFunc(func(row, col int) lipgloss.Style {
			if row == 0 {
				return lipgloss.NewStyle().Padding(0, 1).Bold(true)
			}
			return lipgloss.NewStyle().Padding(0, 1)
		})

	return t.Render()
}

func (g *gameState) View() string {
	if g.gamePhase == 3 {
		return highScoreEntryView(g)
	}
	if g.gamePhase == 4 {
		return lipgloss.JoinVertical(
			lipgloss.Center,
			highScoreTableView(g),
			lipgloss.NewStyle().Padding(1).Render(g.message),
		)
	}
	// Reduced padding for better spacing
	messageStyle := lipgloss.NewStyle().
		Bold(true).
		Foreground(messageColor).
		PaddingTop(1).
		PaddingBottom(1)

	// Credit and bet information with improved styling
	creditInfo := lipgloss.NewStyle().Bold(true).Render(
		lipgloss.NewStyle().Foreground(creditColor).Render(fmt.Sprintf("Credits: %d", g.credits)) +
		" | " +
		lipgloss.NewStyle().Foreground(betColor).Render(fmt.Sprintf("Bet: %d", g.bet)) +
		" | Press +/- to change bet, SPACE for new game")

	return lipgloss.JoinVertical(
		lipgloss.Center,
		payoutTableView(),
		creditInfo,
		prizeView(g.hand),
		lipgloss.JoinHorizontal(lipgloss.Center, cardViews(g.hand, g.gamePhase)...),
		messageStyle.Render(g.message),
	)
}

func main() {
	p := tea.NewProgram(initializeModel())

	if _, err := p.Run(); err != nil {
		fmt.Println("Argh! Error found!")
		os.Exit(1)
	}
}
