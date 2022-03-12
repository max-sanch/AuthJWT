package repository

type Authentication interface {

}

type User interface {

}

type Repository struct {
	Authentication
	User
}

func NewRepository() *Repository {
	return &Repository{}
}