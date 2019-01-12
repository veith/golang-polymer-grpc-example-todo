package dberrors

import "strings"

// every db delivers error codes in a different form.

//https://www.postgresql.org/docs/10/errcodes-appendix.html
//https://www.sqlite.org/rescode.html
//https://dev.mysql.com/doc/refman/8.0/en/server-error-reference.html

func FindErrorByMessageString(err error, str string) bool {
	return caseInsenstiveContains(err.Error(), str)
}
func caseInsenstiveContains(a, b string) bool {
	return strings.Contains(strings.ToUpper(a), strings.ToUpper(b))
}
