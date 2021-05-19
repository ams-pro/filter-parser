package filterparser

import (
	"bufio"
	"errors"
	"strings"
)

type Node struct {
	Token  string
	Left   *Node
	Right  *Node
	Parent *Node `json:"-"`
}

type TapFunction func(node *Node)

func (n *Node) Inorder(t TapFunction) []string {
	arr := []string{}
	if n != nil {
		arr = append(arr, n.Left.Inorder(t)...)

		t(n)
		arr = append(arr, n.Token)

		arr = append(arr, n.Right.Inorder(t)...)
	}
	return arr
}

func (n *Node) ReverseOrder() []string {
	arr := []string{}
	if n != nil {
		arr = append(arr, n.Right.ReverseOrder()...)

		arr = append(arr, n.Token)

		arr = append(arr, n.Left.ReverseOrder()...)
	}
	return arr
}

func ParseFilter(filter string) (*Node, error) {
	scanner := bufio.NewScanner(strings.NewReader(filter))
	scanner.Split(bufio.ScanRunes)

	token := ""

	root := &Node{}

	currentNode := root
	sliceMode := false

	for scanner.Scan() {
		t := scanner.Text()

		switch t {
		case "(":
			if sliceMode {
				return nil, errors.New("Parse Error: Token not allowed")
			}
			currentNode.Token = token

			newNode := &Node{
				Parent: currentNode,
			}

			currentNode.Left = newNode

			currentNode = newNode
			token = ""
		case ",":
			if sliceMode {
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
			if sliceMode {
				return nil, errors.New("Parse Error: Token not allowed")
			}
			if len(token) > 0 {
				currentNode.Token = token
			}
			currentNode = currentNode.Parent

			token = ""

		case "[":
			sliceMode = true
			token += t
		case "]":
			sliceMode = false
			token += t
		default:
			token += t
		}
	}
	return root, nil
}
