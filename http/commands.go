package http

import (
	"github.com/aracoool/face/logs"
	"github.com/google/uuid"
	"time"
)

type Command struct {
	Repository logs.Repository
}

//Create persists a logs record
func (h Command) Create(record logs.Record) (logs.Record, error) {
	record.ID = uuid.New().String()
	record.CreatedAt = time.Now().Format(logs.TimeFormat)

	err := h.Repository.Resist(record)
	if err != nil {
		return record, HandlerError{
			Previous: err,
			Message:  "could not handlerLogCreate a logs record",
		}
	}

	return record, nil
}

//List returns a handlerLogsList of logs
func (h Command) List(params map[string][]string) ([]logs.Record, error) {
	searchCriteria := &logs.SearchCriteria{}
	searchCriteria.PopulateWithMap(params)

	records, err := h.Repository.List(searchCriteria)
	if err != nil {
		return nil, HandlerError{
			Previous: err,
			Message:  "could not get a handlerLogsList of logs",
		}
	}

	return records, nil
}

