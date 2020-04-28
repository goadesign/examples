package persist

import "testing"

func TestActive(t *testing.T) {
	cases := []struct {
		Status Status
		String string
		Active bool
	}{
		{Started, "started", true},
		{Completed, "completed", false},
		{TimedOut, "timed out", false},
		{Failed, "failed", false},
	}
	for _, c := range cases {
		t.Run(c.Status.String(), func(t *testing.T) {
			u := Upload{Status: c.Status}
			if u.Status.String() != c.String {
				t.Errorf("String: got %q, expected %q", c.Status.String(), c.String)
			}
			if u.Active() != c.Active {
				t.Errorf("Active: got %v, expected %v", u.Active(), c.Active)
			}
		})
	}
}
