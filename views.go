package main

import (
  "strings"

  table "github.com/charmbracelet/lipgloss/table"
  lipgloss "github.com/charmbracelet/lipgloss"
  game "github.com/punnie/video-poker/pkg"
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
  view = style.Render(label + strings.Repeat(design, width - lipgloss.Width(label))) + "\n"

  for i := 0; i < height - 1; i++ {
    view += style.Render(strings.Repeat(design, width)) + "\n"
  }

  view += style.Render(strings.Repeat(design, width - lipgloss.Width(reverse_label)) + reverse_label)

  borderStyle := lipgloss.NewStyle().Foreground(color).Background(lipgloss.Color("#FFFFFF")).Padding(0, 1).Margin(1)
  return borderStyle.Render(view)
}

func handCardsView(g game.Game) []string {
  var views []string

  for _, card := range g.HandCards() {
    views = append(views, cardView(card, true))
  }

  return views
}

func payoutTableView() string {
  rows := [][]string{
    { "ROYAL FLUSH",     "250", "500", "750", "1000", "4000" },
    { "STRAIGHT FLUSH",  "50", "100", "150", "200", "250" },
    { "FOUR OF A KIND",  "20", "40", "60", "80", "100" },
    { "FULL HOUSE",      "7", "14", "21", "28", "35" },
    { "FLUSH",           "5", "10", "15", "20", "25" },
    { "STRAIGHT",        "3", "6", "9", "12", "15" },
    { "THREE OF A KIND", "2", "4", "6", "8", "10" },
    { "TWO PAIR",        "1", "2", "3", "4", "5" },
    { "JACKS OR HIGHER", "1", "2", "3", "4", "5" },
  }

  t := table.New().
      Border(lipgloss.NormalBorder()).
      BorderStyle(lipgloss.NewStyle().Foreground(lipgloss.Color("#FFCC00"))).
      Rows(rows...).
      StyleFunc(func(row, col int) lipgloss.Style {
        if col == 0 {
          return lipgloss.NewStyle().Padding(0, 1).Width(35)
        } else {
          return lipgloss.NewStyle().Padding(0, 1).Align(lipgloss.Right).Width(6)
        }
      })

  return t.Render()
}

func prizeView(g game.Game) string {
  var view string

  view = g.Prize()

  return view
}

