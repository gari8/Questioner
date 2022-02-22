package domain

type User struct {
	ID            string      `json:"id"`
	Username      string      `json:"username"`
	Icon          string      `json:"icon"`
	Email         string      `json:"email"`
	Description   string      `json:"description"`
	QuestionCount int         `json:"questionCount"`
	Questions     []*Question `json:"questions"`
	CreatedAt     string      `json:"created_at"`
	UpdatedAt     string      `json:"updated_at"`
	PageInfo      *PageInfo   `json:"pageInfo"`
}

type EditUserInput struct {
	Username    *string `json:"username"`
	Icon        *string `json:"icon"`
	Description *string `json:"description"`
}

type LoginInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
