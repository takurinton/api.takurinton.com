package utils

import (
	"fmt"
	"testing"
	"time"
)

func TestTimeToString(t *testing.T) {
	tests := [2]time.Time{
		time.Now(),
		time.Date(2020, 6, 1, 17, 44, 13, 0, time.Local),
	}

	for i, test := range tests {
		str := TimeToString(test)
		testType := fmt.Sprintf("%T", str)
		switch i {
		case 0:
			if testType != "string" {
				t.Errorf("got %s, want string", testType)
			}
			return
		case 1:
			if testType != "string" {
				t.Errorf("got %s, want string", testType)
			}
			return
		}
	}
}

func TestStringToTimeTest(t *testing.T) {
	tests := [3]string{
		"2006-01-02T15:04:05.000Z",
		"20020-12-28T15:04:05.000Z",
		"9999-01-01T15:04:05.000Z",
	}

	for i, test := range tests {
		_time := StringToTime(test)
		testType := fmt.Sprintf("%T", _time)
		switch i {
		case 0:
			if testType != "time.Time" {
				t.Errorf("got %s, want time.Time", testType)
			}
			return
		case 1:
			if testType != "time.Time" {
				t.Errorf("got %s, want time.Time", testType)
			}
			return
		case 2:
			if testType != "time.Time" {
				t.Errorf("got %s, want time.Time", testType)
			}
			return
		}
	}
}
