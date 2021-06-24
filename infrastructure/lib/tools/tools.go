package tools

import (
	"context"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"os"
	"faves4/ent"
	"faves4/graph/model"
	"strings"
	"time"
)

const TimeLayout = "2006-01-02 15:04:05"

func CreateId() (*string, error) {
	u, err := uuid.NewRandom()
	if err != nil { return nil, err }
	id := u.String()
	return &id, nil
}

func MuskWord(word string) *string {
	musk := strings.Repeat("*", len(word))
	return &musk
}

func DigestWord(word string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(word),12)
	return string(hash), err
}

func CompareHashAndPlane(hash string, plane string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(plane))
	if err != nil {
		return true
	}
	return false
}

func GetEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return defaultValue
	}
	return value
}

func ConvertStringToTime(timeString string) (*time.Time, error) {
	t, err := time.Parse(TimeLayout, timeString)
	if err != nil { return nil, err }
	return &t, nil
}

func CastUser(u *ent.User) *model.User {
	ca := u.CratedAt.String()
	ua := u.UpdatedAt.String()
	return &model.User{
		ID: u.ID,
		Username: u.Username,
		Email: &u.Email,
		Description: &u.Description,
		Password: MuskWord(u.Password),
		Icon: &u.Icon,
		CreatedAt: &ca,
		UpdatedAt: &ua,
	}
}

func CastQuestion(ctx context.Context, q *ent.Question) (*model.Question, error) {
	ca := q.CratedAt.String()
	ua := q.UpdatedAt.String()
	pa := q.PublishedAt.String()
	fa := q.FinishedAt.String()
	cnt, err := q.QueryAnswers().Count(ctx)
	if err != nil { return nil, err }
	ut, err := q.QueryOwner().Only(ctx)
	tmp := &model.Question{
		ID: q.ID,
		User: CastUser(ut),
		Title: q.Title,
		AnswerCount: cnt,
		Content: &q.Content,
		TermStart: &pa,
		TermEnd: &fa,
		TextAfterAnswered: &q.TextAfterAnswered,
		AnswerType: q.AnswerType,
		Answered: false,
		CreatedAt: &ca,
		UpdatedAt: &ua,
	}
	return tmp, nil
}

func CastAnswer(ctx context.Context, a *ent.Answer) (*model.Answer, error) {
	ca := a.CratedAt.String()
	ua := a.UpdatedAt.String()
	ut, err := a.QueryOwner().Only(ctx)
	if err != nil { return nil, err }
	qt, err := a.QueryParent().Only(ctx)
	if err != nil { return nil, err }
	qu, err := CastQuestion(ctx, qt)
	if err != nil { return nil, err }
	return &model.Answer{
		ID: a.ID,
		User: CastUser(ut),
		Question: qu,
		Content: &a.Content,
		CreatedAt: &ca,
		UpdatedAt: &ua,
	}, nil
}

func CastChoice(ctx context.Context, c *ent.Choice) (*model.Choice, error) {
	qt, err := c.QueryParent().Only(ctx)
	if err != nil { return nil, err }
	q, err := CastQuestion(ctx, qt)
	if err != nil { return nil, err }
	cnt, err := c.QueryChoiceanswers().Count(ctx)
	if err != nil { return nil, err }
	return &model.Choice{
		ID: c.ID,
		Content: c.Content,
		Value: cnt,
		Question: q,
	}, nil
}

