package http

import (
	"github.com/aracoool/face/errors"
	"github.com/aracoool/face/logs"
	"github.com/google/uuid"
	"time"
)

type Command struct {
	Repository logs.Repository
}

//Create persists a logs record
func (h Command) Create(record logs.Record) (logs.Record, error) {
	var op errors.Op = "Command.Create"

	record.ID = uuid.New().String()
	record.CreatedAt = time.Now().Format(logs.TimeFormat)

	err := h.Repository.Persist(record)
	if err != nil {
		return record, errors.E(op, err)
	}

	return record, nil
}

//List returns a handlerLogsList of logs
func (h Command) List(params map[string][]string) ([]logs.Record, error) {
	var op errors.Op = "Command.List"

	searchCriteria := &logs.SearchCriteria{}
	searchCriteria.PopulateWithMap(params)

	records, err := h.Repository.List(searchCriteria)
	if err != nil {
		return nil, errors.E(op, err)
	}

	return records, nil
}

