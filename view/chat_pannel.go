package view

import "charm.land/lipgloss/v2"

type ChatPannel struct {
	Transcript string
	InputMessage string
}

func (c ChatPannel) RenderChatPannel(pannelWidth, pannelHeight int) string {	
	inputHeight := 3
	outputHeight := pannelHeight - inputHeight
   	
	outputStyle := lipgloss.NewStyle().
			Width(pannelWidth).Height(outputHeight).
			Border(lipgloss.NormalBorder()).
		 	Align(lipgloss.Right, lipgloss.Top).
			Render(c.Transcript)
	
	inputStyle := lipgloss.NewStyle().
			Width(pannelWidth).Height(inputHeight).
			Border(lipgloss.NormalBorder()).
		 	Align(lipgloss.Right, lipgloss.Bottom).
			Render(c.InputMessage)

	return lipgloss.JoinVertical(lipgloss.Right, outputStyle, inputStyle)
}
