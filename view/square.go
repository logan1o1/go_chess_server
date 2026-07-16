package view

import "charm.land/lipgloss/v2"

type Square struct {
	Piece            string
	Light            bool
	CorrectHighlight bool
	WrongHighlight   bool
}

func (s Square) Render(cellWidth, cellHeight int) string {
	bg := lipgloss.Black

	if s.Light {
		bg = lipgloss.White
	}

	if s.CorrectHighlight {
		bg = lipgloss.Green
	}

	if s.WrongHighlight {
		bg = lipgloss.Red
	}

	cell := lipgloss.NewStyle().
		Width(cellWidth).Height(cellHeight).
		Background(bg).
		Align(lipgloss.Center, lipgloss.Center).
		Render(s.Piece)

	return cell
}
