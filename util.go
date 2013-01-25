// Copyright (c) 2012-2013 Jason McVetta.  This is Free Software, released under
// the terms of the GPL v3.  See http://www.gnu.org/copyleft/gpl.html for details.

package neo4j

import (
	"github.com/kr/pretty"
	"runtime"
	"strconv"
	"strings"
)

// Joins URL fragments
func join(fragments ...string) string {
	parts := []string{}
	for _, str := range fragments {
		if str == "" {
			// Only join non-empty fragment strings
			continue
		}
		str = strings.Trim(str, "/")
		parts = append(parts, str)
	}
	return strings.Join(parts, "/")
}

func logPretty(x interface{}) {
	_, file, line, ok := runtime.Caller(1)
	if !ok {
		file = "???"
		line = 0
	}
	lineNo := strconv.Itoa(line)
	s := file + ":" + lineNo + ": %# v\n"
	pretty.Printf(s, x)
}
