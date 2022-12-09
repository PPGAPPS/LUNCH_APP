package environment

import (
	"fmt"
	"os"
	"strings"

	"github.com/gobuffalo/envy"
)

// ENV is used to help switch settings based on where the

// application is being run. Default is "development".

const (
	Development      = "development"
	Test             = "test"
	Production       = "production"
	ApplicactionName = "lunch"
	SessionName      = "_lunch_session"
)

func Current() string {
	if env := os.Getenv("GO_ENV"); env != "" {
		return env
	}
	return Development

}

func BaseURL() string {
	baseHost := envy.Get("BASE_URL", fmt.Sprintf("http://localhost:%s", envy.Get("PORT", "3000")))

	if strings.HasPrefix(baseHost, "http://") {
		return baseHost

	}

	return fmt.Sprintf("http://%s", baseHost)
}

func LoginURL() string {
	return fmt.Sprintf("%s/login", BaseURL())
}
