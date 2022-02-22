package domain

type Choice struct {
	ID       string    `json:"id"`
	Text     string    `json:"text"`
	Question *Question `json:"question"`
	IsAnswer bool      `json:"isAnswer"`
}

type NewChoice struct {
	Text     string `json:"text"`
	IsAnswer bool   `json:"isAnswer"`
}
