package http

import (
	"errors"
	"github.com/aracoool/face/logs"
	"reflect"
	"testing"
)

func TestCreateHandler(t *testing.T) {
	handler := Command{TestRepository{}}
	record, err := handler.Create(logs.Record{
		ID:        "",
		Source:    "",
		Category:  "",
		Level:     100,
		Message:   "",
		Trace:     "",
		Payload:   nil,
		CreatedAt: "",
	})

	if reflect.TypeOf(err) != reflect.TypeOf(HandlerError{}) {
		t.Errorf("returned error should be HandlerError type")
	}

	if record.ID == "" {
		t.Errorf("ID field was not populated")
	}

	if record.CreatedAt == "" {
		t.Errorf("CreatedAt field was not populated")
	}
}

func TestListHandler(t *testing.T) {
	handler := Command{TestRepository{}}
	records, err := handler.List(map[string][]string{})

	if err != nil && reflect.TypeOf(err) != reflect.TypeOf(HandlerError{}) {
		t.Errorf("returned error should be HandlerError type")
	}

	if len(records) != 1 {
		t.Errorf("Wrong number of records")
	}
}

type TestRepository struct{}

func (t TestRepository) List(criteria *logs.SearchCriteria) ([]logs.Record, error) {
	var records = make([]logs.Record, 0)
	records = append(records, logs.Record{
		ID:        "",
		Source:    "",
		Category:  "",
		Level:     0,
		Message:   "",
		Trace:     "",
		Payload:   nil,
		CreatedAt: "",
	})
	return records, nil
}

func (t TestRepository) Resist(record logs.Record) error {
	return errors.New("test")
}
