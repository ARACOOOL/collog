package logs

import (
	"database/sql"
	"encoding/json"
	"github.com/aracoool/face/errors"
	_ "github.com/go-sql-driver/mysql"
)

var tableName = "logs"

type MysqlRepository struct {
	Credentials string
	db          *sql.DB
	connected   bool
}

//connect tries to connect to DB
func (m *MysqlRepository) connect() error {
	if !m.connected {
		db, err := sql.Open("mysql", m.Credentials)
		if err != nil {
			return err
		}

		m.db = db
		m.connected = true
	}

	return nil
}

func (m *MysqlRepository) List(criteria *SearchCriteria) ([]Record, error) {
	var op errors.Op = "MysqlRepository.List"

	err := m.connect()
	if err != nil {
		return []Record{}, errors.E(op, err)
	}

	stmt, err := m.db.Prepare(criteria.BuildQuery(tableName))
	if err != nil {
		return []Record{}, errors.E(op, err)
	}
	defer stmt.Close()

	var records = make([]Record, 0)

	rows, _ := stmt.Query(criteria.GetConditionValues()...)
	for rows.Next() {
		r := Record{}
		var payload string
		_ = rows.Scan(&r.ID, &r.Source, &r.Category, &r.Level, &r.Message, &r.Trace, &payload, &r.CreatedAt)
		err = json.Unmarshal([]byte(payload), &r.Payload)
		if err != nil {
			return []Record{}, errors.E(op, err)
		}

		records = append(records, r)
	}

	return records, nil
}

func (m *MysqlRepository) Persist(record Record) error {
	var op errors.Op = "MysqlRepository.Persist"

	err := m.connect()
	if err != nil {
		return errors.E(op, err)
	}

	stmt, err := m.db.Prepare("INSERT INTO " + tableName + "(id, source, category, level, message, trace, payload, created_at) VALUES (?,?,?,?,?,?,?,?)")
	if err != nil {
		return errors.E(op, err)
	}
	defer stmt.Close()

	payload, err := json.Marshal(record.Payload)
	if err != nil {
		return errors.E(op, err)
	}

	_, err = stmt.Exec(record.ID, record.Source, record.Category, record.Level, record.Message, record.Trace, payload, record.CreatedAt)
	if err != nil {
		return errors.E(op, err)
	}

	return nil
}
