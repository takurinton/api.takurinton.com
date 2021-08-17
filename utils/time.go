package utils

import "time"

var layout = "2006-01-02T15:04:05.000Z"

func StringToTime(str string) (t time.Time) {
	t, _ = time.Parse(layout, str)
	return
}

func TimeToString(t time.Time) (str string) {
	str = t.Format(layout)
	return
}
