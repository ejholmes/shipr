package shipr

import "testing"

func Test_FindRepo(t *testing.T) {
	defer cleanup(t)

	_, err := Repos.CreateByName("remind101/r101-api")
	if err != nil {
		t.Error(err)
	}

	found, err := Repos.FindByName("remind101/r101-api")
	if err != nil {
		t.Error(err)
	}

	if found == nil || found.Name != "remind101/r101-api" {
		t.Error("FindRepo expected to be able to find a repo by name.")
	}
}

func Test_FindRepo_NotFound(t *testing.T) {
	defer cleanup(t)

	repo, _ := Repos.FindByName("remind101/foo")
	if repo != nil {
		t.Error("Expected an error.")
	}
}
