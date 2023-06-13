package subs

import (
	"fmt"
	"strconv"
	"strings"
)

type Codec struct {
	sep      string
	nilValue string
}

func Constructor297() Codec {
	c := Codec{sep: "x", nilValue: "y"}
	return c
}

// Serializes a tree to a single string.
func (c *Codec) serialize(root *TreeNode) string {
	if root == nil {
		return ""
	}

	var values []string
	f := func(_r *TreeNode) {}
	f = func(_r *TreeNode) {
		if _r == nil {
			values = append(values, c.nilValue)
		} else {
			values = append(values, fmt.Sprintf("%d", _r.Val))
			f(_r.Left)
			f(_r.Right)
		}
	}
	f(root)
	return strings.Join(values, c.sep)
}

// Deserializes your encoded data to tree.
func (c *Codec) deserialize(data string) *TreeNode {
	if data == "" {
		return nil
	}
	values := strings.Split(data, c.sep)
	f := func() *TreeNode { return nil }
	f = func() *TreeNode {
		_v := values[0]
		values = values[1:]
		if _v == c.nilValue {
			return nil
		}
		_vInt, _ := strconv.Atoi(_v)
		node := TreeNode{Val: _vInt}
		node.Left = f()
		node.Right = f()
		return &node
	}
	return f()
}
