package logs

import "fmt"

type RepositoryError struct {
	Previous error
	Message  string
}

func (e RepositoryError) Error() string {
	return fmt.Sprintf("%v: %v", e.Message, e.Previous)
}

//Repository ...
type Repository interface {
	Persist(record Record) error
	List(criteria *SearchCriteria) ([]Record, error)
}
