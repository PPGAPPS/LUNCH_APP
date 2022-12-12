package middlewares

import (
	"lunch/app/models"

	csrf "github.com/gobuffalo/mw-csrf"
	paramlogger "github.com/gobuffalo/mw-paramlogger"
	"github.com/wawandco/ox/pkg/buffalotools"
)

var (
	// DB middleware sets the tx variable
	// to be a transaction and wraps the request on it
	// when the method is POST, PUT, DELETE and PATCH.
	// Otherwise it just sets the db connection in the
	// tx context value.
	DB = buffalotools.DatabaseMiddleware(models.DB(), nil)

	// ParameterLogger logs out parameters that the app received
	// taking care of sensitive data.
	ParameterLogger = paramlogger.ParameterLogger

	// CSRF middleware protects from CSRF attacks.
	CSRF = csrf.New
)
