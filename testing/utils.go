package testing

import (
	"encoding/xml"
	"reflect"
	"strings"
)

type XMLNode struct {
	XMLName  xml.Name
	Attrs    []xml.Attr `xml:",any,attr"`
	Content  string     `xml:",chardata"`
	Children []XMLNode  `xml:",any"`
}

func normalizeNode(n *XMLNode) {
	n.Content = strings.TrimSpace(n.Content)
	for i := range n.Children {
		normalizeNode(&n.Children[i])
	}
}

func XMLEqual(a, b []byte) (bool, error) {
	var nodeA, nodeB XMLNode
	if err := xml.Unmarshal(a, &nodeA); err != nil {
		return false, err
	}
	if err := xml.Unmarshal(b, &nodeB); err != nil {
		return false, err
	}
	normalizeNode(&nodeA)
	normalizeNode(&nodeB)
	return reflect.DeepEqual(nodeA, nodeB), nil
}
