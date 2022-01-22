package link

import "io"

// Link represents a link (<a href="">) in HTML doc
type Link struct {
	Href string
	Text string
}

// Parse takes in HTML doc, returns slice of links parsed from it
func Parse(r io.Reader) ([]Link, error) {
	return nil, nil
}
