package api

import (
	"randomGoTests/httpgordle/internal/session"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToGameResponse(t *testing.T) {
	id := "1682279480"
	tt := map[string]struct {
		game session.Game
		want GameResponse
	}{
		"nominal": {
			game: session.Game{
				ID:           session.GameID(id),
				AttemptsLeft: 4,
				Guesses: []session.Guess{{
					Word:     "FAUNE",
					Feedback: "⬜️🟡⬜️⬜️⬜️",
				}},
				Status: session.StatusPlaying,
			},
			want: GameResponse{
				ID:           id,
				AttemptsLeft: 4,
				Guesses: []Guess{{
					Word:     "FAUNE",
					Feedback: "⬜️🟡⬜️⬜️⬜️",
				}},
				Solution: "",
				Status:   session.StatusPlaying,
			},
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			got := ToGameResponse(tc.game)
			assert.Equal(t, tc.want, got)
		})
	}
}
