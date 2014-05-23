package main

import "testing"

func Test_FindRepo(t *testing.T) {
	clean()

	_, err := CreateRepo("remind101/r101-api", false)
	if err != nil {
		t.Error(err)
	}

	found, err := FindRepo("remind101/r101-api")
	if err != nil {
		t.Error(err)
	}

	if found == nil || found.Name != "remind101/r101-api" {
		t.Error("FindRepo expected to be able to find a repo by name.")
	}
}

func Test_FindRepo_NotFound(t *testing.T) {
	clean()

	repo, _ := FindRepo("remind101/foo")
	if repo != nil {
		t.Error("Expected an error.")
	}
}
