// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// AnswersColumns holds the columns for the "answers" table.
	AnswersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "answer_type", Type: field.TypeString},
		{Name: "content", Type: field.TypeString},
		{Name: "crated_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime, Nullable: true},
		{Name: "question_answers", Type: field.TypeString, Nullable: true},
		{Name: "user_answers", Type: field.TypeString, Nullable: true},
	}
	// AnswersTable holds the schema information for the "answers" table.
	AnswersTable = &schema.Table{
		Name:       "answers",
		Columns:    AnswersColumns,
		PrimaryKey: []*schema.Column{AnswersColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "answers_questions_answers",
				Columns:    []*schema.Column{AnswersColumns[5]},
				RefColumns: []*schema.Column{QuestionsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "answers_users_answers",
				Columns:    []*schema.Column{AnswersColumns[6]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// ChoicesColumns holds the columns for the "choices" table.
	ChoicesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "content", Type: field.TypeString},
		{Name: "question_choices", Type: field.TypeString, Nullable: true},
	}
	// ChoicesTable holds the schema information for the "choices" table.
	ChoicesTable = &schema.Table{
		Name:       "choices",
		Columns:    ChoicesColumns,
		PrimaryKey: []*schema.Column{ChoicesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "choices_questions_choices",
				Columns:    []*schema.Column{ChoicesColumns[2]},
				RefColumns: []*schema.Column{QuestionsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// ChoiceAnswersColumns holds the columns for the "choice_answers" table.
	ChoiceAnswersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "crated_at", Type: field.TypeTime},
		{Name: "choice_choiceanswers", Type: field.TypeInt, Nullable: true},
		{Name: "user_choiceanswers", Type: field.TypeString, Nullable: true},
	}
	// ChoiceAnswersTable holds the schema information for the "choice_answers" table.
	ChoiceAnswersTable = &schema.Table{
		Name:       "choice_answers",
		Columns:    ChoiceAnswersColumns,
		PrimaryKey: []*schema.Column{ChoiceAnswersColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "choice_answers_choices_choiceanswers",
				Columns:    []*schema.Column{ChoiceAnswersColumns[2]},
				RefColumns: []*schema.Column{ChoicesColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "choice_answers_users_choiceanswers",
				Columns:    []*schema.Column{ChoiceAnswersColumns[3]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// QuestionsColumns holds the columns for the "questions" table.
	QuestionsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString, Unique: true},
		{Name: "title", Type: field.TypeString},
		{Name: "content", Type: field.TypeString},
		{Name: "text_after_answered", Type: field.TypeString, Default: ""},
		{Name: "enabled", Type: field.TypeBool, Default: true},
		{Name: "answer_type", Type: field.TypeString},
		{Name: "published_at", Type: field.TypeTime, Nullable: true},
		{Name: "finished_at", Type: field.TypeTime, Nullable: true},
		{Name: "crated_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime, Nullable: true},
		{Name: "user_questions", Type: field.TypeString, Nullable: true},
	}
	// QuestionsTable holds the schema information for the "questions" table.
	QuestionsTable = &schema.Table{
		Name:       "questions",
		Columns:    QuestionsColumns,
		PrimaryKey: []*schema.Column{QuestionsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "questions_users_questions",
				Columns:    []*schema.Column{QuestionsColumns[10]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString, Unique: true},
		{Name: "username", Type: field.TypeString},
		{Name: "icon", Type: field.TypeString, Default: ""},
		{Name: "email", Type: field.TypeString, Unique: true},
		{Name: "password", Type: field.TypeString},
		{Name: "description", Type: field.TypeString, Default: ""},
		{Name: "crated_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime, Nullable: true},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:        "users",
		Columns:     UsersColumns,
		PrimaryKey:  []*schema.Column{UsersColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		AnswersTable,
		ChoicesTable,
		ChoiceAnswersTable,
		QuestionsTable,
		UsersTable,
	}
)

func init() {
	AnswersTable.ForeignKeys[0].RefTable = QuestionsTable
	AnswersTable.ForeignKeys[1].RefTable = UsersTable
	ChoicesTable.ForeignKeys[0].RefTable = QuestionsTable
	ChoiceAnswersTable.ForeignKeys[0].RefTable = ChoicesTable
	ChoiceAnswersTable.ForeignKeys[1].RefTable = UsersTable
	QuestionsTable.ForeignKeys[0].RefTable = UsersTable
}
