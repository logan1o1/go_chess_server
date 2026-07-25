package view

import (
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
)

func RenderGamePage(width, height int, chatPanel ChatPannel, board [8][8]Square) tea.View {
	chatWidth := width / 3
	boardwidth := width - chatWidth

	chatPanel.InputMessage = "> "

	boardContent := lipgloss.JoinVertical(
			lipgloss.Top,
			HowToQuit(),
			RenderChessBoard(board, boardwidth),
		)

	boardView := lipgloss.NewStyle().Width(boardwidth).Render(boardContent)

	chatView := chatPanel.RenderChatPannel(chatWidth, height - 1)

	v := tea.NewView(
		lipgloss.JoinHorizontal(lipgloss.Top, boardView, chatView),
	)

	v.AltScreen = true

	return v
}
