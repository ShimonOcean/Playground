package link

import (
	"fmt"
	"io"

	"golang.org/x/net/html"
)

// Link represents a link (<a href="">) in HTML doc
type Link struct {
	Href string
	Text string
}

// Parse takes in HTML doc, returns slice of links parsed from it
func Parse(r io.Reader) ([]Link, error) {
	doc, err := html.Parse(r)
	if err != nil {
		return nil, err
	}
	nodes := linkNodes(doc)
	for _, node := range nodes {
		fmt.Println(node)
	}

	// dfs(doc, "")

	return nil, nil
}

// DFS, find all a tag nodes for their links
func linkNodes(n *html.Node) []*html.Node {
	if n.Type == html.ElementNode && n.Data == "a" {
		return []*html.Node{n}
	}
	var ret []*html.Node

	for c := n.FirstChild; c != nil; c = c.NextSibling {

		ret = append(ret, linkNodes(c)...)
	}
	return ret
}

// Depth First Search through all nodes of html doc
// func dfs(n *html.Node, padding string) {

// 	msg := n.Data
// 	// If html element, add correct brackets
// 	if n.Type == html.ElementNode {
// 		msg = "<" + msg + ">"
// 	}
// 	fmt.Println(padding, msg)

// 	// c = n's first child, if c is not nil continue to next sibling
// 	for c := n.FirstChild; c != nil; c = c.NextSibling {
// 		dfs(c, padding+"  ")
// 	}
// }
