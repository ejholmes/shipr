package heroku

import (
	"reflect"
	"testing"

	"github.com/ejholmes/heroku-go/v3"
)

func Test_NewBuildResultLines(t *testing.T) {
	tests := []struct {
		lines []struct {
			Line   string `json:"line"`
			Stream string `json:"stream"`
		}
		idx      int
		expected []*BuildResultLine
	}{
		{
			idx: 0,
			lines: []struct {
				Line   string `json:"line"`
				Stream string `json:"stream"`
			}{
				{Line: "Hello\n", Stream: "STDOUT"},
			},
			expected: []*BuildResultLine{
				{Line: "Hello\n", Stream: "STDOUT"},
			},
		},
		{
			idx: 0,
			lines: []struct {
				Line   string `json:"line"`
				Stream string `json:"stream"`
			}{
				{Line: "Hello\n", Stream: "STDOUT"},
				{Line: "World\n", Stream: "STDOUT"},
			},
			expected: []*BuildResultLine{
				{Line: "Hello\n", Stream: "STDOUT"},
				{Line: "World\n", Stream: "STDOUT"},
			},
		},
		{
			idx: 1,
			lines: []struct {
				Line   string `json:"line"`
				Stream string `json:"stream"`
			}{
				{Line: "Hello\n", Stream: "STDOUT"},
				{Line: "World\n", Stream: "STDOUT"},
			},
			expected: []*BuildResultLine{
				{Line: "World\n", Stream: "STDOUT"},
			},
		},
	}

	for i, test := range tests {
		b := &heroku.BuildResult{Lines: test.lines}
		lines := newBuildResultLines(b, test.idx)

		if !reflect.DeepEqual(lines, test.expected) {
			t.Errorf("%v: Want %v; Got %v", i, test.expected, lines)
		}
	}
}
