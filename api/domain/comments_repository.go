package domain

type CommentsRepository interface {
	Save(Comment) error
	Retrieve() ([]Comment, error)
}
