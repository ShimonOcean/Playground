package main

import (
	"fmt"
	"strings"
)

var exampleHtml = `
<html>
<body>
	<h1>Hi<h1>
	<a href="/other-page">Link to another page</a>
	<a href="/page-two">Link to second page</a>
</body>
</html>
`

func main() {
	// Convert a string into a reader
	r := strings.NewReader(exampleHtml)
	links, err := link.Parse(r)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", links)
}
