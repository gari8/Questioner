package domain

type Question struct {
	ID                string    `json:"id"`
	Title             string    `json:"title"`
	Content           string    `json:"content"`
	TextAfterAnswered string    `json:"textAfterAnswered"`
	Enabled           bool      `json:"enabled"`
	Choices           []*Choice `json:"choices"`
	User              *User     `json:"user"`
	CorrectRate       int       `json:"correctRate"`
	CreatedAt         string    `json:"created_at"`
	UpdatedAt         string    `json:"updated_at"`
	PageInfo          *PageInfo `json:"pageInfo"`
}

type EditQuestionInput struct {
	Title             *string      `json:"title"`
	Content           *string      `json:"content"`
	Enabled           *bool        `json:"enabled"`
	TextAfterAnswered *string      `json:"textAfterAnswered"`
	Choices           []*NewChoice `json:"choices"`
}

type NewQuestionInput struct {
	ID                string       `json:"id"`
	Title             string       `json:"title"`
	Content           string       `json:"content"`
	Enabled           bool         `json:"enabled"`
	UserID            string       `json:"userId"`
	Choices           []*NewChoice `json:"choices"`
	TextAfterAnswered *string      `json:"textAfterAnswered"`
}
