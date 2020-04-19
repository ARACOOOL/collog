package logs

import "testing"

func TestConstants(t *testing.T) {
	if Debug != 100 {
		t.Errorf("debug constant should have a value 100")
	}

	if Error != 400 {
		t.Errorf("error constant should have a value 400; value is %d", Error)
	}
}
