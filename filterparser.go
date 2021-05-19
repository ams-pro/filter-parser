package filterparser

import (
	"bufio"
	"strings"
)

type Node struct {
	Token  string
	Left   *Node
	Right  *Node
	Parent *Node `json:"-"`
}

func ParseFilter(filter string) *Node {
	scanner := bufio.NewScanner(strings.NewReader(filter))
	scanner.Split(bufio.ScanRunes)

	token := ""

	root := &Node{}

	currentNode := root
	vectorMode := false

	for scanner.Scan() {
		t := scanner.Text()

		switch t {
		case "(":
			currentNode.Token = token

			newNode := &Node{
				Parent: currentNode,
			}

			currentNode.Left = newNode

			currentNode = newNode
			token = ""
		case ",":
			if vectorMode {
				token += t
				continue
			}
			if len(currentNode.Token) == 0 {
				currentNode.Token = token
			}

			currentNode = currentNode.Parent

			newNode := &Node{
				Parent: currentNode,
			}

			currentNode.Right = newNode

			currentNode = newNode

			token = ""
		case ")":
			if len(token) > 0 {
				currentNode.Token = token
			}
			currentNode = currentNode.Parent

			token = ""

		case "[":
			vectorMode = true
			token += t
		case "]":
			vectorMode = false
			token += t
		default:
			token += t
		}
	}
	return root
}
