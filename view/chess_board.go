package view

import "charm.land/lipgloss/v2"

func RenderChessBoard(board [8][8]Square) string {
	var rows []string
	for _, row := range board {
		var cells []string
		for _, sq := range row {
			cells = append(cells, sq.Render(10, 4))
		}
		rows = append(rows, lipgloss.JoinHorizontal(lipgloss.Top, cells...))
	}
	return lipgloss.JoinVertical(lipgloss.Left, rows...)
}
