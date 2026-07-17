package main

import (
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
	"github.com/charmbracelet/ssh"
	"github.com/logan1o1/go_chess_server/view"
)

type model struct {
	term      string
	profile   string
	width     int
	height    int
	bg        string
	quitStyle lipgloss.Style
	board     [8][8]view.Square
}

func (m model) Init() tea.Cmd {
	return tea.Batch(
		tea.RequestBackgroundColor,
	)
}

func teaHandler(s ssh.Session) (tea.Model, []tea.ProgramOption) {
	pty, _, _ := s.Pty()
	var board [8][8]view.Square

	pieces := *view.ParseFEN("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR")
	for r := range 8 {
		for c := range 8 {
			board[r][c] = view.Square{
				Piece: pieces[r][c],
				Light: (r+c)%2 == 0,
			}
		}

	}

	m := model{
		term:   pty.Term,
		width:  pty.Window.Width,
		height: pty.Window.Height,
		bg:     "light",
		board:  board,
	}

	return m, []tea.ProgramOption{}
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.ColorProfileMsg:
		m.profile = msg.String()
	case tea.BackgroundColorMsg:
		if msg.IsDark() {
			m.bg = "dark"
		}
	case tea.WindowSizeMsg:
		m.height = msg.Height
		m.width = msg.Width
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return m, tea.Quit
		}
	}
	return m, nil
}

func (m model) View() tea.View {
	v := tea.NewView(
		lipgloss.JoinVertical(
			lipgloss.Top,
			view.HowToQuit(),
			view.RenderChessBoard(m.board),
		),
	)
	v.AltScreen = true

	return v
}
