package tailwindcss

import (
	"fmt"
	"strings"

	"github.com/gobuffalo/tags/v3"
)

func buildOptions(opts tags.Options, err bool) {
	if opts["class"] == nil {
		opts["class"] = ""
	}

	if opts["tag_only"] != true {
		parts := []string{
			fmt.Sprint(opts["class"]),
			"focus:ring-indigo-500 focus:border-indigo-500 flex-1 block w-full rounded-md sm:text-sm border-gray-300",
		}

		opts["class"] = strings.Join(parts, " ")
	}

	if err {
		opts["class"] = strings.Join([]string{fmt.Sprint(opts["class"]), "is-invalid"}, " ")
	}
}
