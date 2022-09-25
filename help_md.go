package goopt

import (
	"bytes"
	"fmt"
	"strings"
)

func HelpMarkdown(optHelpFormatFunc func(string) string) string {
	buf := new(bytes.Buffer)
	if len(opts) > 1 {
		fmt.Fprintln(buf, "# Flags:\n")
	}
	for _, o := range opts {
		names := []string{}
		for index, r_name := range o.names {
			name := string(r_name)
			if index == 0 && o.allowsArg != nil {
				name += "=" + *o.allowsArg
			}
			append(&names, "`"+name+"`")
		}
		for _, r_name := range o.shortnames {
			name := string(r_name)
			append(&names, "`-"+name+"`")
		}
		fmt.Fprintln(buf, "## "+strings.Join(names, ", "))
		fmt.Fprintf(buf, "%v\n\n", optHelpFormatFunc(o.help))
	}
	return buf.String()
}
