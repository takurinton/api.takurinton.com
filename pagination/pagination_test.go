package pagination

import (
	"testing"
)

func TestPaginationDjangoFormat(t *testing.T) {
	tests := []struct {
		now      int
		total    int
		next     int
		prev     int
		category string
	}{
		{1, 5, 2, 1, ""},
		{2, 5, 3, 1, ""},
		{1, 5, 2, 1, "インターン"},
	}
	for i, v := range tests {
		category, next, prev := GetParams(v.now, v.total, v.next, v.prev, v.category)
		switch i {
		case 0:
			if next != v.next {
				t.Errorf("got %d, want 2", v.next)
			}
			if prev != nil {
				t.Errorf("got %d, want nil", v.prev)
			}
			return
		case 1:
			if next != v.next {
				t.Errorf("got %d, want 3", v.next)
			}
			if prev != v.prev {
				t.Errorf("got %d, want 1", v.prev)
			}
			return
		case 2:
			if next != v.next && category != v.category {
				t.Errorf("got page=%d, category=%s, want page=2, category=インターン", v.next, v.category)
			}
			if prev != nil {
				t.Errorf("got %d, want nil", v.prev)
			}
			return
		}

	}
}
