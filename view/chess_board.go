package view

import "charm.land/lipgloss/v2"

func RenderChessBoard(board [8][8]Square, boardWidth int) string {
	cellWidth := boardWidth / 11
	cellHeight := max(cellWidth/2, 1)

	var rows []string
	for _, row := range board {
		var cells []string
		for _, sq := range row {
			cells = append(cells, sq.Render(cellWidth, cellHeight))
		}
		rows = append(rows, lipgloss.JoinHorizontal(lipgloss.Top, cells...))
	}
	return lipgloss.JoinVertical(lipgloss.Left, rows...)
}
