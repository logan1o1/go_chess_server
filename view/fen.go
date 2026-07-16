package view

import "strings"

func parseFen(fen string) *[8][8]string {
	var board [8][8]string
	ranks := strings.Split(fen, "/")
	for r, rank := range ranks {
		c := 0
		for _, ch := range rank {
			if ch >= '1' && ch <= '8' {
				c += int(ch - '0')
				continue
			}
			board[r][c] = fenToUnicode[ch]
			c++
		}
	}
	return &board
}

var fenToUnicode = map[rune]string{
	'r': "♜", 'n': "♞", 'b': "♝", 'q': "♛", 'k': "♚", 'p': "♟",
	'R': "♖", 'N': "♘", 'B': "♗", 'Q': "♕", 'K': "♔", 'P': "♙",
}
