package main

import (
	"bytes"
	"regexp"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "shimon"
	password = "your-passwd"
	dbname   = "phone_normalize"
)

func main() {

}

// Normalizes phone numbers
// Iterate through string, catching all numerical characters and disposing of extra
func normalize(phone string) string {
	// bytes buffer is more efficient than string concatenation with +
	var buf bytes.Buffer
	for _, ch := range phone {
		if ch >= '0' && ch <= '9' {
			// WriteRune: appends UTF-8 of input to buffer
			buf.WriteRune(ch)
		}
	}

	return buf.String()
}

func normalizeRegex(phone string) string {
	re := regexp.MustCompile("[^0-9]")
	return re.ReplaceAllString(phone, "")
}
