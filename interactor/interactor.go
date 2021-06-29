package interactor

import (
	"faves4/ent"
	"faves4/graph/repository"
)

type (
	interactor struct {
		client *ent.Client
	}
	Interactor interface {
		NewRepository() *Repository
	}
	Repository struct {
		repository.UserRepository
		repository.QuestionRepository
		repository.AnswerRepository
		repository.ChoiceRepository
		repository.ChoiceAnswerRepository
	}
)

func NewInteractor(client *ent.Client) Interactor {
	return &interactor{client}
}

func (i *interactor) NewRepository() *Repository {
	return &Repository{
		repository.NewUserRepository(i.client),
		repository.NewQuestionRepository(i.client),
		repository.NewAnswerRepository(i.client),
		repository.NewChoiceRepository(i.client),
		repository.NewChoiceAnswerRepository(i.client),
	}
}