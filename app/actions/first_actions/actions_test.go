package first_actions_test

import (
	"lunch/app"
	"testing"

	"github.com/gobuffalo/suite/v4"
)

type ActionSuite struct {
	*suite.Action
}

func Test_ActionSuite(t *testing.T) {
	bapp, err := app.New()
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	as := &ActionSuite{suite.NewAction(bapp)}
	suite.Run(t, as)
}
