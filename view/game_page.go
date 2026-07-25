package view

import (
	"charm.land/lipgloss/v2"
)

func RenderGamePage(width, height int, chatPanel ChatPannel, board [8][8]Square) string {
	chatWidth := width / 3
	boardwidth := width - chatWidth

	chatPanel.InputMessage = "> "

	boardContent := lipgloss.JoinVertical(
			lipgloss.Top,
			RenderChessBoard(board, boardwidth),
		)

	boardView := lipgloss.NewStyle().Width(boardwidth).Render(boardContent)

	chatView := chatPanel.RenderChatPannel(chatWidth, height - 1)

	gamePage := lipgloss.JoinHorizontal(lipgloss.Top, boardView, chatView)

	return gamePage
}
