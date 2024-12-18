package goopt

import (
	"bytes"
	"fmt"
	"strings"
)

func HelpMarkdown(optHelpFormatFunc func(string) string) string {
	buf := new(bytes.Buffer)
	if len(opts) > 1 {
		fmt.Fprintf(buf, "# Flags:\n\n")
	}
	for _, o := range opts {
		names := []string{}
		for index, name := range o.names {
			if index == 0 && o.allowsArg != nil {
				name += "=" + *o.allowsArg
			}
			names = append(names, "`"+name+"`")
		}
		for _, r_name := range o.shortnames {
			name := string(r_name)
			names = append(names, "`-"+name+"`")
		}
		fmt.Fprintln(buf, "## "+strings.Join(names, ", "))
		fmt.Fprintf(buf, "%v\n\n", optHelpFormatFunc(o.help))
	}
	return buf.String()
}
