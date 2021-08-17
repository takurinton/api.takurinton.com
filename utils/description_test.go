package utils

import (
	"testing"
)

func TestDescription(t *testing.T) {
	tests := [1]string{
		"# どうも\r\n\r\n皆さん、コロナですがどうお過ごしでしょうか？僕はとっても暇です",
	}

	var content string
	for i, test := range tests {
		content = ParseContents(test)
		switch i {
		case 0:
			if content != "どうも皆さん、コロナですがどうお過ごしでしょうか？僕はとっても暇です" {
				t.Errorf("got %s, want どうも皆さん、コロナですがどうお過ごしでしょうか？僕はとっても暇です", content)
			}
			return
		}
	}
}
