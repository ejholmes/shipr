package shipr

import (
	"reflect"
	"testing"

	h "github.com/ejholmes/heroku-go/v3"
	"github.com/remind101/shipr/heroku"
)

func Test_NewBuildResultLines(t *testing.T) {
	tests := []struct {
		lines []struct {
			Line   string `json:"line"`
			Stream string `json:"stream"`
		}
		idx      int
		expected []*herokuLogLine
	}{
		{
			idx: 0,
			lines: []struct {
				Line   string `json:"line"`
				Stream string `json:"stream"`
			}{
				{Line: "Hello\n", Stream: "STDOUT"},
			},
			expected: []*herokuLogLine{
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
			expected: []*herokuLogLine{
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
			expected: []*herokuLogLine{
				{Line: "World\n", Stream: "STDOUT"},
			},
		},
	}

	for i, test := range tests {
		b := &heroku.BuildResult{&h.BuildResult{Lines: test.lines}}
		lines := newLines(b, test.idx)

		if !reflect.DeepEqual(lines, test.expected) {
			t.Errorf("%v: Want %v; Got %v", i, test.expected, lines)
		}
	}
}
