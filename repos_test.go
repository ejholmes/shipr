package shipr

import "testing"

func Test_RepoName(t *testing.T) {
	tests := []struct {
		input       string
		owner, repo string
	}{
		{
			input: "remind101/r101-api",
			owner: "remind101",
			repo:  "r101-api",
		},
	}

	for i, test := range tests {
		repoName := RepoName(test.input)

		if test.owner != repoName.Owner() {
			t.Errorf("%v: Got %v; Want %v", i, test.owner, repoName.Owner())
		}

		if test.repo != repoName.Repo() {
			t.Errorf("%v: Got %v; Want %v", i, test.repo, repoName.Repo())
		}
	}
}
