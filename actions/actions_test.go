package actions

import (
	"os"
	"testing"

	"github.com/gobuffalo/suite/v4"
)

type ActionSuite struct {
	*suite.Action
}

func Test_ActionSuite(t *testing.T) {
	err := os.Setenv("GO_ENV", "test")
	if err != nil {
		t.Fatal(err)
	}

	action, err := suite.NewActionWithFixtures(App(), os.DirFS("../fixtures"))
	if err != nil {
		t.Fatal(err)
	}
	as := &ActionSuite{
		Action: action,
	}
	suite.Run(t, as)
}
