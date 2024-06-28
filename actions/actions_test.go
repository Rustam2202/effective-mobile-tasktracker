package actions

import (
	"fmt"
	"os"
	"testing"

	"github.com/gobuffalo/suite/v4"
)

type ActionSuite struct {
	*suite.Action
}

func Test_ActionSuite(t *testing.T) {
	env := os.Getenv("GO_ENV")
	fmt.Printf("env: %v\n", env)
	err := os.Setenv("GO_ENV", "test")
	if err != nil {
		t.Fatal(err)
	}
	env = os.Getenv("GO_ENV")

	action := suite.NewAction(App())
	// if err != nil {
	// 	t.Fatal(err)
	// }

	as := &ActionSuite{
		Action: action,
	}
	suite.Run(t, as)
}
