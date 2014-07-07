package main

import "testing"

func Test_HerokuDeploy_AppName(t *testing.T) {
	tests := []struct {
		deploy   *HerokuDeploy
		expected string
	}{
		{
			&HerokuDeploy{&FakeDeployment{environment: "production", repoName: "remind101/r101-api"}},
			"r101-api",
		},
		{
			&HerokuDeploy{&FakeDeployment{environment: "staging", repoName: "remind101/r101-api"}},
			"r101-api-staging",
		},
	}

	for i, test := range tests {
		if test.deploy.AppName() != test.expected {
			t.Errorf("%v: Want %v; Got %v", i, test.expected, test.deploy.AppName())
		}
	}
}
