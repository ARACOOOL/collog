package http

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/aracoool/face/logs"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCommandCreate(t *testing.T) {
	s := Server{
		command: Command{
			Repository: &EmptyRepository{},
		},
	}
	srv := httptest.NewServer(http.HandlerFunc(s.handlerLogCreate))
	defer srv.Close()

	var logRecord = []byte(`{"source":"postman","category":"test","level":100,"message":"test message","trace":"{}","payload":{"test":"value"}}`)
	resp, err := http.Post(srv.URL, "application/json", bytes.NewBuffer(logRecord))
	if err != nil {
		t.Fatal(err)
	}

	if resp.StatusCode != 201 {
		t.Fatalf("Received non-200 response: %d\n", resp.StatusCode)
	}

	record := &logs.Record{}
	err = json.NewDecoder(resp.Body).Decode(record)
	if err != nil {
		t.Fatal(err)
	}

	if len(record.ID) != 36 {
		t.Errorf("Invalid record ID, the length of the ID should be more then 35 bytes")
	}
}

func TestCommandList(t *testing.T) {
	s := Server{
		command: Command{
			Repository: &EmptyRepository{},
		},
	}
	srv := httptest.NewServer(http.HandlerFunc(s.handlerLogsList))
	defer srv.Close()

	resp, err := http.Get(srv.URL)
	if err != nil {
		t.Fatal(err)
	}

	if resp.StatusCode != 200 {
		t.Fatalf("Received non-200 response: %d\n", resp.StatusCode)
	}

	expected := fmt.Sprintf("[]\n")
	actual, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}

	if expected != string(actual) {
		t.Errorf("Expected the message '%s' and got '%q'\n", expected, actual)
	}
}

type EmptyRepository struct {
	records []logs.Record
}

func (t *EmptyRepository) Persist(record logs.Record) error {
	t.records = append(t.records, record)
	return nil
}

func (t *EmptyRepository) List(criteria *logs.SearchCriteria) ([]logs.Record, error) {
	t.records = make([]logs.Record, 0)
	return t.records, nil
}
