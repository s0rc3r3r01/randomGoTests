package api

const (
	NewGameRoute = "/games"
)

type GameResponse struct {
	ID           string  `json:"id"`
	AttemptsLeft byte    `json:"attempts_left"`
	Guesses      []Guess `json:"guesses"`
	WordLength   byte    `json:"wordlength"`
	Solution     string  `json:"solution"`
	Status       string  `json:"status"`
}

type Guess struct {
	Word     string `json:"word"`
	Feedback string `json:"feedback"`
}
