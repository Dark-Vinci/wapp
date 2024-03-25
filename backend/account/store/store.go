package store

type Store struct {
}

func New() *string {
	a := "response"
	return &a
}
