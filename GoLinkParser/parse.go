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
	var links []Link

	for _, node := range nodes {
		link = append(links, buildLink(node))
		fmt.Println(node)
	}

	// dfs(doc, "")

	return links, nil
}

func buildLink(n *html.Node) Link {
	var ret Link
	// n.Attr is a slice, need to iterate over every element
	for _, attr := range n.Attr {
		if attr.Key == "href" {
			ret.Href = attr.Val
			break
		}
	}
	// Want the text for return
	ret.Text = text(n)

	return ret
}

// Return string representing all text inside a node, ignoring comments
func text(n *html.Node) string {
	if n.Type == html.TextNode {
		return n.Data
	}
	// If node is a i.e. comment, doctype node
	if n.Type != html.ElementNode {
		return ""
	}

	// For every child, retrieve its text
	var ret string
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		ret += text(c) + " "
	}
	return ret
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
