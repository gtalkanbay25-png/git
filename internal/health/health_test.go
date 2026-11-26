package health

import "testing"

func TestScoreMessage(t *testing.T) {
	cases := []struct {
		msg  string // сообщение коммита
		want int    // сколько баллов ждём
	}{
		{"add: parser for stack input in parser.go", 3},
		{"wip", 0},
		{"update README", 2},
		{"add: README with project description", 3},
		{"fixed stuff", 0},
		{"", 0},
	}
	for _, c := range cases {
		got := ScoreMessage(c.msg)
		if got != c.want {
			t.Errorf("ScoreMessage(%q) = %d, want %d", c.msg, got, c.want)
		}
	}
}

func TestLabel(t *testing.T) {
	cases := []struct {
		avg  float64
		want string
	}{
		{2.6, "tidy"},
		{2.5, "tidy"},
		{2.0, "acceptable"},
		{1.5, "acceptable"},
		{1.4, "messy"},
	}
	for _, c := range cases {
		got := Label(c.avg)
		if got != c.want {
			t.Errorf("Label(%v) = %q, want %q", c.avg, got, c.want)
		}
	}
}
